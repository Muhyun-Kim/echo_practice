package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Username  string    `gorm:"size:255;not null"`
	Email     string    `gorm:"size:255;not null;unique"`
	Password  string    `gorm:"size:255;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

type UserDTO struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type AuthorDTO struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
