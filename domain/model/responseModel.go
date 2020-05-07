package model

type ResponseNoData struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ResponseSearchData struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []User `json:"data"`
}

type ResponseInsertData struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    User   `json:"data"`
}
