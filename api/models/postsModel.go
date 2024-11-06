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
type Poste struct {
	ID           int       `json:"id"`
	CreatedAt    string    `json:"created_at"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	Author       string    `json:"author"`
	Category     string    `json:"category"`
	Comments     []Comment `json:"comments"`
	Liked        *bool     `json:"liked"`
	Disliked     *bool     `json:"disliked"`
	LikesCount   int       `json:"likes_count"`
	DislikeCount int       `json:"dislike_count"`
}

type Engagement struct {
	UserID        int    `json:"user_id"`
	PosteID       int    `json:"post_id"`
	CommentID     int    `json:"comment_id"`
	LikeAction    string `json:"like_action"`
	DislikeAction string `json:"dislike_action"`
}

type Filter struct {
	Type string `json:"type"`
	Category string `json:"categories"`
}