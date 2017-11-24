package db

import (
	"context"
	"fmt"
	"github.com/mlbright/forecast/v2"
	"googlemaps.github.io/maps"
	"gopkg.in/guregu/null.v3"
	"strconv"
	"time"
)

type Post struct {
	Id        int64     `json:"id"`
	ParentId  null.Int  `json:"parent_id"`
	UserName  string    `json:"username"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`

	City null.String `json:"city"`
	Lat  null.Float  `json:"latitude"`
	Long null.Float  `json:"longitude"`
	Temp null.Float  `json:"temperature"`
}

func GetPosts() ([]Post, error) {
	rows, err := db.Query("SELECT id, parent_id, username, content, created_at, city, lat, long, temp FROM posts")
	if err != nil {
		return nil, err
	}

	var posts []Post

	defer rows.Close()
	for rows.Next() {
		var p Post
		err = rows.Scan(&p.Id, &p.ParentId, &p.UserName, &p.Content, &p.CreatedAt, &p.City, &p.Lat, &p.Long, &p.Temp)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}

	return posts, nil
}

func (p *Post) Save() error {
	if p.Lat.IsZero() || p.Long.IsZero() {
		if err := p.UpdateCoordinates(); err != nil {
			return err
		}
	}
	if p.Temp.IsZero() {
		if err := p.UpdateWeather(); err != nil {
			return err
		}
	}

	query := "INSERT INTO posts (parent_id, content, username, city, lat, long, temp) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, created_at"
	rows := db.QueryRow(query, p.ParentId, p.Content, p.UserName, p.City, p.Lat, p.Long, p.Temp)
	return rows.Scan(&p.Id, &p.CreatedAt)
}

func (p *Post) Delete() error {
	if p.Id == 0 {
		return fmt.Errorf("Cannot delete item with no ID")
	}

	_, err := db.Exec("DELETE FROM posts WHERE id = $1", p.Id)
	return err
}

func (p *Post) UpdateCoordinates() error {
	c, err := maps.NewClient(maps.WithAPIKey("AIzaSyC26xl41gd31W4o_WbV5PSVOlDDmJn5FXw"))
	if err != nil {
		return err
	}

	r := &maps.GeocodingRequest{Address: p.City.String}
	result, err := c.Geocode(context.Background(), r)
	if len(result) == 0 {
		return fmt.Errorf("No latlng found for provided address")
	}

	loc := result[0].Geometry.Location
	p.Lat = null.FloatFrom(loc.Lat)
	p.Long = null.FloatFrom(loc.Lng)

	return nil
}

func (p *Post) UpdateWeather() error {
	key := "a6b8d6d9f6b39876591441b23fa87026"

	lat := strconv.FormatFloat(p.Lat.Float64, 'f', -1, 64)
	long := strconv.FormatFloat(p.Long.Float64, 'f', -1, 64)

	f, err := forecast.Get(key, lat, long, "now", forecast.SI, forecast.English)
	if err != nil {
		return err
	}

	p.Temp = null.FloatFrom(f.Currently.Temperature)

	return nil
}
