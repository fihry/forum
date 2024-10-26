package Models

type User struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	SessionKey string `json:"session"`
}
