package models

import (
	"time"
)

type User struct {
	ID uint `json:"id"`
	Username string `gorm:"uniqueIndex" json:"username" binding:"required"`
}

type Post struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	AuthorID uint `json:"author_id"`
	Author *User `gorm:"constraint:OnDelete:SET NULL;" json:"author,omitempty"`
	Topics []*Topic `gorm:"many2many:post_topics" json:"topics"`
}

// Structure of request body for creating a new post
type CreatePostRequest struct {
	Title string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required,min=10"`
	AuthorID uint `json:"author_id" binding:"required"`
}

type Comment struct {
	ID uint `json:"id"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	PostID uint `json:"postId"`
	AuthorID uint `json:"authorId"`
	Author User `gorm:"constraint:OnDelete:SET NULL;" json:"author"`
}

// Structure of request body for creating a new comment
type CreateCommentRequest struct {
	Content string `json:"content" binding:"required"`
	PostID uint `json:"postId"`
	AuthorID uint `json:"authorId" binding:"required"`
}

type Topic struct {
	ID uint `json:"id"`
	Name string `gorm:"uniqueIndex" json:"name" binding:"required"`
	Posts []*Post `gorm:"many2many:post_topics;"`
}