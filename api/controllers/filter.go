package controllers

import (
	"fmt"

	"forum/models"
)

func FilterPostsByCategory(category string) ([]models.Poste, error) {
	// select the posts with the given category
	query := "SELECT * FROM posts WHERE category =?"
	stmt, err := Database.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(category)
	if err != nil {
		return nil, fmt.Errorf("failed to execute statement: %w", err)
	}
	defer rows.Close()

	var posts []models.Poste
	for rows.Next() {
		var post models.Poste
		err := rows.Scan(&post.ID, &post.CreatedAt, &post.Title, &post.Content, &post.Author, &post.Category, &post.LikesCount, &post.DislikeCount, &post.Liked, &post.Disliked)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func FilterByReaction(userId int) ([]models.Poste, error) {
	// select the posts where the user has liked or disliked
	var posts []models.Poste
	// select just the posts that the user has liked or disliked
	query := "SELECT p.id, p.title, p.createdAt, p.content, p.author, p.category, p.likesCount, p.dislikesCount, e.like AS liked, e.dislike AS disliked FROM posts p JOIN engagements e ON p.id = e.postId WHERE e.userId =?"
	stm, err := Database.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stm.Close()
	rows, err := stm.Query(userId)
	if err != nil {
		return nil, fmt.Errorf("failed to execute statement: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var post models.Poste
		err := rows.Scan(&post.ID, &post.Title, &post.CreatedAt, &post.Content, &post.Author, &post.Category,
			&post.LikesCount, &post.DislikeCount, &post.Liked, &post.Disliked)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func FilterByAuthor(userId int) ([]models.Poste, error) {
	// select the posts where the user is the author
	var posts []models.Poste
	query := "SELECT * FROM posts WHERE author =?"
	stm, err := Database.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stm.Close()
	rows, err := stm.Query(userId)
	if err != nil {
		return nil, fmt.Errorf("failed to execute statement: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var post models.Poste
		err := rows.Scan(&post.ID, &post.CreatedAt, &post.Title, &post.Content, &post.Author, &post.Category, &post.LikesCount, &post.DislikeCount, &post.Liked, &post.Disliked)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
