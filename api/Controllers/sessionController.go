package Controllers

import (
	"forum/api/Models"

	"github.com/gofrs/uuid"
)

func (db *Database) NewSession(user Models.User) (Models.User, error) {
	UUID, _ := uuid.NewV4()
	SessionKey := UUID.String()
	_, err := db.DB.Exec("INSERT OR IGNORE INTO session (id, Key) VALUES (?, ?)",
		user.ID,
		SessionKey)
	user.Session = SessionKey
	if err != nil {
		return user, err
	}
	return user, nil
}

func (db *Database) GetSession(Key string) (Models.User, error) {
	user := Models.User{}
	err := db.DB.QueryRow("SELECT * FROM session WHERE key = ?", Key).Scan(
		&user.ID,
		&user.Session)
	if err != nil {
		return user, err
	}
	return db.GetUser(user.ID)
}