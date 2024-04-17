package entity

import (
	"time"

	"github.com/google/uuid"
)

type People struct {
	ID          uuid.UUID `json:"id"`
	UserID      uuid.UUID `json:"userId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Relation    string    `json:"relation"`
	Picture     string    `json:"picture"`
	CreatedAt   time.Time `json:"createdAt"`
}
