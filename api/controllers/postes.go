package controllers

import (
	"database/sql"
	"fmt"
	"log"

	"forum/models"
)

func CreatePoste(P models.Poste) (int64, error) {
	query := "INSERT INTO posts (title, content, author, category) VALUES (?, ?, ?, ?)"
	stmt, err := Database.Prepare(query)
	if err != nil {
		return 0, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(P.Title, P.Content, P.Author, P.Category)
	if err != nil {
		return 0, fmt.Errorf("failed to execute statement: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert id: %w", err)
	}
	return id, nil
}

func GetAllPosts() ([]models.Poste, error) {
	query := "SELECT id,title,createdAt,content,author,category,likesCount,dislikesCount FROM posts"
	stmt, err := Database.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, fmt.Errorf("failed to execute statement: %w", err)
	}
	defer rows.Close()
	posts := []models.Poste{}
	for rows.Next() {
		poste := models.Poste{}
		err := rows.Scan(
			&poste.ID,
			&poste.Title,
			&poste.CreatedAt,
			&poste.Content,
			&poste.Author,
			&poste.Category,
			&poste.LikesCount,
			&poste.DislikeCount)
		if err != nil {
			return nil, fmt.Errorf("failed to scan rows: %w", err)
		}
		poste.Comments, err = GetPostComments(poste.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to get comments: %w", err)
		}
		posts = append(posts, poste)
	}
	return posts, nil
}

// this function will return all posts related to a user to enable the user to see if he liked or disliked the post
func GetAllPostsWithEngagement(userId int) ([]models.Poste, error) {
	query := `
		SELECT
			p.id, p.title,p.createdAt, p.content, p.author, p.category,
			p.likesCount, p.dislikesCount,
			e.like AS liked, e.dislike AS disliked
			FROM posts p
			LEFT JOIN engagement e ON p.id = e.postId AND e.userId = ?
			`
	stmt, err := Database.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(userId)
	if err != nil {
		return nil, fmt.Errorf("failed to execute statement: %w", err)
	}
	defer rows.Close()
	posts := []models.Poste{}
	for rows.Next() {
		poste := models.Poste{}
		var liked, disliked sql.NullBool
		err := rows.Scan(
			&poste.ID,
			&poste.Title,
			&poste.CreatedAt,
			&poste.Content,
			&poste.Author,
			&poste.Category,
			&poste.LikesCount,
			&poste.DislikeCount,
			&liked,
			&disliked)
		if err != nil {
			log.Println(err)
			return nil, fmt.Errorf("failed to scan rows: %w", err)
		}
		if liked.Valid {
			poste.Liked = &liked.Bool
		} else {
			poste.Liked = new(bool)
		}
		if disliked.Valid {
			poste.Disliked = &disliked.Bool
		} else {
			poste.Disliked = new(bool)
		}
		posts = append(posts, poste)
	}
	return posts, nil
}

func GetPoste(id int) (models.Poste, error) {
	poste := models.Poste{}
	query := "SELECT * FROM posts WHERE id = ?"
	stmt, err := Database.Prepare(query)
	if err != nil {
		return poste, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(
		&poste.ID,
		&poste.Title,
		&poste.CreatedAt,
		&poste.Content,
		&poste.Author,
		&poste.Category,
		&poste.LikesCount,
		&poste.DislikeCount,
		&poste.Liked,
		&poste.Disliked)
	if err != nil {
		return poste, fmt.Errorf("failed to scan rows: %w", err)
	}
	return poste, nil
}

func GetPostsByCategory(category string) ([]models.Poste, error) {
	query := "SELECT * FROM posts WHERE category = ?"
	stmt, err := Database.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
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
			&poste.CreatedAt,
			&poste.Content,
			&poste.Author,
			&poste.Category,
			&poste.LikesCount,
			&poste.DislikeCount)
		if err != nil {
			return nil, fmt.Errorf("failed to scan rows: %w", err)
		}
		posts = append(posts, poste)
	}
	return posts, nil
}

func GetPostsByAuthor(author string) ([]models.Poste, error) {
	stmt, err := Database.Prepare("SELECT * FROM posts WHERE author = ?")
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(author)
	if err != nil {
		return nil, fmt.Errorf("failed to execute statement: %w", err)
	}
	defer rows.Close()
	posts := []models.Poste{}
	for rows.Next() {
		poste := models.Poste{}
		err := rows.Scan(
			&poste.ID,
			&poste.Title,
			&poste.CreatedAt,
			&poste.Content,
			&poste.Author,
			&poste.Category,
			&poste.LikesCount,
			&poste.DislikeCount)
		if err != nil {
			return nil, fmt.Errorf("failed to scan rows: %w", err)
		}
		posts = append(posts, poste)
	}
	return posts, nil
}

func GetPostComments(postId int) ([]models.Comment, error) {
	stmt, err := Database.Prepare("SELECT * FROM comments WHERE postId = ?")
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(postId)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("failed to execute statement: %w", err)
	}
	defer rows.Close()
	comments := []models.Comment{}
	for rows.Next() {
		comment := models.Comment{}
		err := rows.Scan(
			&comment.ID,
			&comment.PosteID,
			&comment.CreatedAt,
			&comment.Author,
			&comment.Content,
			&comment.LikesCount,
			&comment.DislikeCount,
			&comment.Liked,
			&comment.Disliked,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan rows: %w", err)
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

func CreateComment(C models.Comment) (int64, error) {
	query := "INSERT INTO comments (content, author, post_id, createdAt) VALUES (?, ?, ?, ?)"
	stmt, err := Database.Prepare(query)
	if err != nil {
		return 0, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(C.Content, C.Author, C.ID, C.CreatedAt)
	if err != nil {
		return 0, fmt.Errorf("failed to execute statement: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert id: %w", err)
	}
	return id, nil
}

func GetCommentsByPostId(id int) ([]models.Comment, error) {
	query := "SELECT * FROM comments WHERE post_id = ?"
	stmt, err := Database.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, fmt.Errorf("failed to execute statement: %w", err)
	}
	defer rows.Close()
	comments := []models.Comment{}
	for rows.Next() {
		comment := models.Comment{}
		err := rows.Scan(
			&comment.ID,
			&comment.Content,
			&comment.CreatedAt,
			&comment.Author,
			&comment.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to scan rows: %w", err)
		}
		comments = append(comments, comment)
	}
	return comments, nil
}
