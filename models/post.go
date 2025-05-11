package models

// Post represents the structure of a blog post or an item.
type Post struct {
	ID    string `json:"id,omitempty"`
	Title string `json:"title"`
	Body  string `json:"body"`
}
