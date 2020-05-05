package model

import "time"

type ResponseNoData struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ResponseSearchData struct {
	Status  int            `json:"status"`
	Message string         `json:"message"`
	Data    SearchResponse `json:"data"`
}

type DataResponse struct {
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	Content   string    `json:"content"`
}

type SearchResponse struct {
	Time      string         `json:"time"`
	Hits      string         `json:"hits,omitempty"`
	Documents []DataResponse `json:"data"`
}
