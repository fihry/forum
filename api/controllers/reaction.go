package controllers

import (
	"fmt"
)

// like and dislike the post
func LikePost(postId int) (int64, error) {
	query := "UPDATE posts SET likesCount = likesCount + 1 WHERE id = ?"
	stmt, err := Database.Prepare(query)
	if err != nil {
		return 0, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(postId)
	if err != nil {
		return 0, fmt.Errorf("failed to execute statement: %w", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("failed to get rows affected: %w", err)
	}
	return rows, nil
}

// remove like from the post
func RemoveLike(postId int) (int64, error) {
	query := "UPDATE posts SET likesCount = likesCount - 1 WHERE id = ?"
	stmt, err := Database.Prepare(query)
	if err != nil {
		return 0, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(postId)
	if err != nil {
		return 0, fmt.Errorf("failed to execute statement: %w", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("failed to get rows affected: %w", err)
	}
	return rows, nil
}

func DislikePost(postId int) (int64, error) {
	query := "UPDATE posts SET dislikesCount = dislikesCount + 1 WHERE id = ?"
	stmt, err := Database.Prepare(query)
	if err != nil {
		return 0, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(postId)
	if err != nil {
		return 0, fmt.Errorf("failed to execute statement: %w", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("failed to get rows affected: %w", err)
	}
	return rows, nil
}

// remove dislike from the post
func RemoveDislike(postId int) (int64, error) {
	query := "UPDATE posts SET dislikesCount = dislikesCount - 1 WHERE id = ?"
	stmt, err := Database.Prepare(query)
	if err != nil {
		return 0, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(postId)
	if err != nil {
		return 0, fmt.Errorf("failed to execute statement: %w", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("failed to get rows affected: %w", err)
	}
	return rows, nil
}

func GetLIkesAndDislike(postId int) (int, int, error) {
	query := "SELECT likes, dislikes FROM posts WHERE id = ?"
	stmt, err := Database.Prepare(query)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()
	var likes, dislikes int
	err = stmt.QueryRow(postId).Scan(&likes, &dislikes)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to scan rows: %w", err)
	}
	return likes, dislikes, nil
}

func AddLikeToEngagement(postId, userId int) error {
	query := "INSERT INTO engagement (postId, userId, like, dislike) VALUES (?, ?, ?, ?)"
	stmt, err := Database.Prepare(query)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(postId, userId, true, false)
	if err != nil {
		return fmt.Errorf("failed to execute statement: %w", err)
	}
	return nil
}

func UpdateLikeToEngagement(postId, userId int) error {
	query := "UPDATE engagement SET like =? WHERE postId = ? AND userId = ?"
	stmt, err := Database.Prepare(query)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(true, postId, userId)
	if err != nil {
		return fmt.Errorf("failed to execute statement: %w", err)
	}
	// Check how many rows were affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	// If no rows were updated, call AddLikeToEngagement
	if rowsAffected == 0 {
		return AddLikeToEngagement(postId, userId)
	}
	return nil
}

func AddDislikeToEngagement(postId, userId int) error {
	query := "INSERT INTO engagement (postId, userId, like, dislike) VALUES (?, ?, ?, ?)"
	stmt, err := Database.Prepare(query)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(postId, userId, false, true)
	if err != nil {
		return fmt.Errorf("failed to execute statement: %w", err)
	}
	return nil
}

func UpdateDislikeToEngagement(postId, userId int) error {
	query := "UPDATE engagement SET dislike =? WHERE postId = ? AND userId = ?"
	stmt, err := Database.Prepare(query)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(true, postId, userId)
	if err != nil {
		return fmt.Errorf("failed to execute statement: %w", err)
	}

	// Check how many rows were affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	// If no rows were updated, call AddDislikeToEngagement
	if rowsAffected == 0 {
		return AddDislikeToEngagement(postId, userId)
	}
	return nil
}

func RemLikeFromEngagement(postId, userId int) error {
	query := "UPDATE engagement SET like =? WHERE postId = ? AND userId = ?"
	stmt, err := Database.Prepare(query)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(false, postId, userId)
	if err != nil {
		return fmt.Errorf("failed to execute statement: %w", err)
	}
	return nil
}

func RemDislikeFromEngagement(postId, userId int) error {
	query := "UPDATE engagement SET dislike =? WHERE postId = ? AND userId = ?"
	stmt, err := Database.Prepare(query)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(false, postId, userId)
	if err != nil {
		return fmt.Errorf("failed to execute statement: %w", err)
	}
	return nil
}
