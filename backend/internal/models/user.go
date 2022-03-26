package models

import (
	"time"
)

//User model
type User struct {
	ID        int       `json:"id" gorm:"primary_key"`
	Username  string    `json:"username" gorm:";unique;not null"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at" gorm:"NOT NULL DEFAULT NOW()"`
}
