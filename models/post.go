package models

import (
	"gorm.io/gorm"
)

// Post represents the structure of a blog post or an item.
type Post struct {
	gorm.Model
	Title string `json:"title"`
	Body  string `json:"body"`
}
