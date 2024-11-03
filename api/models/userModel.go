package models

type User struct {
	ID           int       `json:"id"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	Email        string    `json:"email"`
	SessionKey   string    `json:"session"`
	PostsLike    []int     `json:"posts_like"`
	PostsDislike []int     `json:"posts_dislike"`
	Posts        []Poste   `json:"posts"`
	Comments     []Comment `json:"comments"`
}
