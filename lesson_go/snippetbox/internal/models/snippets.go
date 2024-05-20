package models

import (
	"database/sql"
	"time"
)

type Snippet struct {
	ID int
	Tile string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

func(m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires)
    VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`
	return 0, nil
}

func(m *SnippetModel) Get(id int) (Snippet,error) {
	return Snippet{}, nil 
}

func(m *SnippetModel) Latest() ([]Snippet, error) {
	return nil, nil
}