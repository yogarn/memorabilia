package model

type CreateDiary struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}

type UpdateDiary struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}
