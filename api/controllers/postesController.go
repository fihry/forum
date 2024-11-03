package controllers

import (
	"log"

	"forum/api/models"
)

func GetAllPosts() ([]models.Poste, error) {
	stmt, err := Database.Prepare("SELECT * FROM posts")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	posts := []models.Poste{}
	for rows.Next() {
		poste := models.Poste{}
		err := rows.Scan(
			&poste.ID,
			&poste.Title,
			&poste.Content,
			&poste.Author,
			&poste.Category,
			&poste.LikesCount,
			&poste.DislikeCount,
			&poste.Disliked,
			&poste.Liked)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		posts = append(posts, poste)
	}
	return posts, nil
}

func GetPost(id int) (models.Poste, error) {
	poste := models.Poste{}
	stmt, err := Database.Prepare("SELECT * FROM posts WHERE id = ?")
	if err != nil {
		return poste, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(
		&poste.ID,
		&poste.Title,
		&poste.Content,
		&poste.Author,
		&poste.Category,
		&poste.LikesCount,
		&poste.DislikeCount,
		&poste.Liked,
		&poste.Disliked)
	if err != nil {
		return poste, err
	}
	return poste, nil
}

func CreatePoste(P models.Poste) (int64, error) {
	stmt, err := Database.Prepare("INSERT INTO posts (title, content, author, category) VALUES (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(P.Title, P.Content, P.Author, P.Category)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func GetPostsByCategory(category string) ([]models.Poste, error) {
	stmt, err := Database.Prepare("SELECT * FROM posts WHERE category = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	posts := []models.Poste{}
	for rows.Next() {
		poste := models.Poste{}
		err := rows.Scan(
			&poste.ID,
			&poste.Title,
			&poste.Content,
			&poste.Author,
			&poste.Category,
			&poste.LikesCount,
			&poste.DislikeCount)
		if err != nil {
			return nil, err
		}
		posts = append(posts, poste)
	}
	return posts, nil
}

func GetPostsByAuthor(author string) ([]models.Poste, error) {
	stmt, err := Database.Prepare("SELECT * FROM posts WHERE author = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(author)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	posts := []models.Poste{}
	for rows.Next() {
		poste := models.Poste{}
		err := rows.Scan(
			&poste.ID,
			&poste.Title,
			&poste.Content,
			&poste.Author,
			&poste.Category,
			&poste.LikesCount,
			&poste.DislikeCount)
		if err != nil {
			return nil, err
		}
		posts = append(posts, poste)
	}
	return posts, nil
}

func AddComment(C models.Comment) (int64, error) {
	stmt, err := Database.Prepare("INSERT INTO comments (content, author, post_id) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(C.Content, C.Author, C.ID)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func GetCommentsByPostId(id int) ([]models.Comment, error) {
	stmt, err := Database.Prepare("SELECT * FROM comments WHERE post_id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	comments := []models.Comment{}
	for rows.Next() {
		comment := models.Comment{}
		err := rows.Scan(
			&comment.ID,
			&comment.Content,
			&comment.Author,
			&comment.ID)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}
