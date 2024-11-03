package models

type Comment struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
	PosteID int    `json:"poste_id"`
}
type Poste struct {
	ID           int       `json:"id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	Author       string    `json:"author"`
	Category     string    `json:"category"`
	Comment      []Comment `json:"comments"`
	Liked        bool      `json:"liked"`
	Disliked     bool      `json:"disliked"`
	LikesCount   int       `json:"likes_count"`
	DislikeCount int       `json:"dislike_count"`
}
