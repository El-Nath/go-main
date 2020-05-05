package model

import "time"

type Data struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	Content   string    `json:"content"`
}

type DataRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
