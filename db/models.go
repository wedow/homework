package db

import (
	"fmt"
	"time"
)

type Post struct {
	Id        int64     `json:"id"`
	UserName  string    `json:"username"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

func GetPosts() ([]Post, error) {
	rows, err := db.Query("SELECT id, username, content, created_at FROM posts")
	if err != nil {
		return nil, err
	}

	var posts []Post

	defer rows.Close()
	for rows.Next() {
		var post Post
		rows.Scan(&post.Id, &post.UserName, &post.Content, &post.CreatedAt)
		posts = append(posts, post)
	}

	return posts, nil
}

func (p *Post) Save() error {
	rows := db.QueryRow("INSERT INTO posts (content, username) VALUES ($1, $2) RETURNING id, created_at", p.Content, p.UserName)
	return rows.Scan(&p.Id, &p.CreatedAt)
}

func (p *Post) Delete() error {
	if p.Id == 0 {
		return fmt.Errorf("Cannot delete item with no ID")
	}

	_, err := db.Exec("DELETE FROM posts WHERE id = $1", p.Id)
	return err
}
