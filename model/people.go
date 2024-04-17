package model

type CreatePeople struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Relation    string `json:"relation" binding:"required"`
}
