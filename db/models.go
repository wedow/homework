package db

import (
	"fmt"
	"gopkg.in/guregu/null.v3"
	"time"
)

type Post struct {
	Id        int64     `json:"id"`
	ParentId  null.Int  `json:"parent_id"`
	UserName  string    `json:"username"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func GetPosts() ([]Post, error) {
	rows, err := db.Query("SELECT id, parent_id, username, content, created_at FROM posts")
	if err != nil {
		return nil, err
	}

	var posts []Post

	defer rows.Close()
	for rows.Next() {
		var post Post
		err = rows.Scan(&post.Id, &post.ParentId, &post.UserName, &post.Content, &post.CreatedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (p *Post) Save() error {
	query := "INSERT INTO posts (parent_id, content, username) VALUES ($1, $2, $3) RETURNING id, created_at"
	rows := db.QueryRow(query, p.ParentId, p.Content, p.UserName)
	return rows.Scan(&p.Id, &p.CreatedAt)
}

func (p *Post) Delete() error {
	if p.Id == 0 {
		return fmt.Errorf("Cannot delete item with no ID")
	}

	_, err := db.Exec("DELETE FROM posts WHERE id = $1", p.Id)
	return err
}
