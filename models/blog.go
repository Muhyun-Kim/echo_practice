package models

import "time"

type Blog struct {
	ID        uint          `gorm:"primaryKey"`
	Title     string        `gorm:"size:255;not null"`
	Content   string        `gorm:"size:255;not null"`
	AuthorID  uint          `gorm:"not null"`
	Author    User          `gorm:"foreignKey:AuthorID"`
	CreatedAt time.Time     `gorm:"autoCreateTime"`
	UpdatedAt time.Time     `gorm:"autoUpdateTime"`
	Comments  []BlogComment `gorm:"foreignKey:BlogID"`
}

type BlogComment struct {
	ID        uint      `gorm:"primaryKey"`
	Content   string    `gorm:"size:255;not null"`
	BlogID    uint      `gorm:"not null"`
	Blog      Blog      `gorm:"foreignKey:BlogID"`
	AuthorID  uint      `gorm:"not null"`
	Author    User      `gorm:"foreignKey:AuthorID"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type BlogDTO struct {
	ID        uint         `json:"id"`
	Title     string       `json:"title"`
	Content   string       `json:"content"`
	AuthorID  uint         `json:"author_id"`
	Author    AuthorDTO    `json:"author"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	Comments  []CommentDTO `json:"comments,omitempty"`
}

type CommentDTO struct {
	ID        uint      `json:"id"`
	Content   string    `json:"content"`
	BlogID    uint      `json:"blog_id"`
	AuthorID  uint      `json:"author_id"`
	Author    AuthorDTO `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
