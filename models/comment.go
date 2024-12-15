package models

type Comment struct {
	ID           int    `json:"id"`
	CreatedAt    string `json:"created_at"`
	Content      string `json:"content"`
	Author       string `json:"author"`
	PosteID      int    `json:"poste_id"`
	Liked        *bool  `json:"liked"`
	Disliked     *bool  `json:"disliked"`
	LikesCount   int    `json:"likes_count"`
	DislikeCount int    `json:"dislike_count"`
}
