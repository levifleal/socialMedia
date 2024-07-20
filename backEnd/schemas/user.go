package schemas

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id           string `json:"id"`
	Name         string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"password"`
}

type UserRespose struct {
	Id        string    `json:"id"`
	Name      string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
}
