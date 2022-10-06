package models

import "time"

// Model declaration
type User struct {
	ID       uint      `json:"id" gorm:"primary_key"`
	Username string    `json:"username" gorm:"unique"`
	LastGame time.Time `json:"last_game"`
}

// Input declaration
type CreateUserInput struct {
	Username string    `json:"username" binding:"required"`
	LastGame time.Time `json:"last_game"`
}
