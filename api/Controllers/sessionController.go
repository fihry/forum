package Controllers

import (
    "forum/api/Models"
    "github.com/gofrs/uuid"
)

func (db *Database) NewSession(user Models.User) (Models.User, error) {
    UUID, _ := uuid.NewV4()
    user.SessionKey = UUID.String()

    stmt, err := db.DB.Prepare("INSERT OR IGNORE INTO sessions (id, Key) VALUES (?, ?)")
    if err != nil {
        return user, err
    }
    defer stmt.Close()

    _, err = stmt.Exec(user.ID, user.SessionKey)
    if err != nil {
        return user, err
    }
    return user, nil
}

func (db *Database) GetSession(Key string) (Models.User, error) {
    user := Models.User{}

    stmt, err := db.DB.Prepare("SELECT * FROM session WHERE key = ?")
    if err != nil {
        return user, err
    }
    defer stmt.Close()

    err = stmt.QueryRow(Key).Scan(&user.ID, &user.SessionKey)
    if err != nil {
        return user, err
    }
    return db.GetUserById(user.ID)
}

func (db *Database) DeleteSession(Key string) error {
    stmt, err := db.DB.Prepare("DELETE FROM session WHERE key = ?")
    if err != nil {
        return err
    }
    defer stmt.Close()

    _, err = stmt.Exec(Key)
    if err != nil {
        return err
    }
    return nil
}