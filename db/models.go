package db

import (
	"fmt"
	"time"
)

type Post struct {
	Id        int64
	Content   string
	CreatedAt time.Time `db:"created_at"`
}

func (t *Post) Save() error {
	rows := db.QueryRow("INSERT INTO posts (content) VALUES ($1) RETURNING id, created_at", t.Content)
	return rows.Scan(&t.Id, &t.CreatedAt)
}

func (t *Post) Delete() error {
	if t.Id == 0 {
		return fmt.Errorf("Cannot delete item with no ID")
	}

	_, err := db.Exec("DELETE FROM posts WHERE id = $1", t.Id)
	return err
}
