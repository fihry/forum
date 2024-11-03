package models

type Comment struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
	PosteID int    `json:"poste_id"`
	Likes   int    `json:"likes"`
	Dislikes int    `json:"dislikes"`
}

type Poste struct {
	ID       int       `json:"id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Author   string    `json:"author"`
	Category string    `json:"category"`
	Comment  []Comment `json:"comment"`
	Reaction string	`json:"reaction"`
}


