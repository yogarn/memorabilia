package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	Username       string    `json:"username"`
	Password       string    `json:"password"`
	Email          string    `json:"email"`
	RoleID         int       `json:"role"`
	ProfilePicture string    `json:"profilePicture"`
	CreatedAt      time.Time `json:"createdAt"`
}
