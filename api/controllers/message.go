package controllers

import (
	"fmt"

	"forum/models"
)

func GetCommentsForPost(postId, userId int) ([]models.Comment, error) {
	// get all the comments of the post
	comments, err := GetAllComments(postId, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to get comments: %w", err)
	}
	// check if the user session is valid
	// if !IsUserValid(userId) {
	//     return comments, nil
	// }
	// get all the likes and dislikes for the past comment
	for i, comment := range comments {
		liked, disliked, err := GetCommentReaction(comment.ID, userId)
		if err != nil {
			return nil, fmt.Errorf("failed to get comment reaction: %w", err)
		}
		comments[i].Liked = &liked
		comments[i].Disliked = &disliked
	}
	return comments, nil
}

func GetCommentReaction(commentId, userId int) (bool, bool, error) {
	query := "SELECT liked, disliked FROM engagement WHERE commentId =? AND userId =?"
	stmt, err := Database.Prepare(query)
	if err != nil {
		return false, false, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(commentId, userId)
	if err != nil {
		return false, false, fmt.Errorf("failed to execute statement: %w", err)
	}
	defer rows.Close()

	var liked, disliked bool
	err = rows.Scan(&liked, &disliked)
	if err != nil {
		return false, false, fmt.Errorf("failed to scan rows: %w", err)
	}
	return liked, disliked, nil
}

func GetAllComments(postId, userId int) ([]models.Comment, error) {
	query := "SELECT * FROM comments WHERE postId =?"
	stmt, err := Database.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(postId)
	if err != nil {
		return nil, fmt.Errorf("failed to execute statement: %w", err)
	}
	defer rows.Close()

	var comments []models.Comment
	for rows.Next() {
		comment := models.Comment{}
		err := rows.Scan(
			&comment.ID,
			&comment.PosteID,
			&comment.CreatedAt,
			&comment.Content,
			&comment.Author,
			&comment.LikesCount,
			&comment.DislikeCount,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan rows: %w", err)
		}
		comments = append(comments, comment)
	}
	return comments, nil
}
