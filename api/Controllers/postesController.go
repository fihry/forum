package controllers

import (
	"forum/api/Models"
)

func (D *Database) GetAllPosts() ([]Models.Poste, error) {
	rows, err := D.DB.Query("SELECT * FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	posts := []Models.Poste{}
	for rows.Next() {
		poste := Models.Poste{}
		err := rows.Scan(
			&poste.ID,
			&poste.Title,
			&poste.Content,
			&poste.Author,
			&poste.Category)
		if err != nil {
			return nil, err
		}
		posts = append(posts, poste)
	}
	return posts, nil
}

func (D *Database) GetPoste(P Models.Poste) (Models.Poste, error) {
	poste := Models.Poste{}
	err := D.DB.QueryRow("SELECT * FROM posts WHERE id = ?", P.ID).Scan(
		&poste.ID,
		&poste.Title,
		&poste.Content,
		&poste.Author,
		&poste.Category)
	if err != nil {
		return poste, err
	}
	return poste, nil
}

func (D *Database) PostPoste(P Models.Poste) (int64, error) {
	result, err := D.DB.Exec("INSERT INTO posts (title, content, author, category) VALUES (?, ?, ?, ?)",
		P.Title, P.Content, P.Author, P.Category)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (D *Database) UpdatePoste(P Models.Poste) (int64, error) {
	result, err := D.DB.Exec("UPDATE posts SET title = ?, content = ?, author = ?, category = ? WHERE id = ?",
		P.Title, P.Content, P.Author, P.Category, P.ID)
	if err != nil {
		return 0, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rows, nil
}

func (D *Database) DeletePoste(P Models.Poste) (int64, error) {
	result, err := D.DB.Exec("DELETE FROM posts WHERE id = ?", P.ID)
	if err != nil {
		return 0, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rows, nil
}
