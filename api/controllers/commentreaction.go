package controllers

import "fmt"

// add 1 to the likes count of the comment
func LikeComment(commentId int) (int64, error) {
	query := "UPDATE comments SET likesCount = likesCount + 1 WHERE id = ?"
	stmt, err := Database.Prepare(query)
	if err != nil {
		fmt.Println("Error preparing", err)
		return 0, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(commentId)
	if err != nil {
		return 0, fmt.Errorf("failed to execute statement: %w", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("failed to get rows affected: %w", err)
	}
	return rows, nil
}

// remove 1 from the likes count of the comment
func RemoveLikeComment(commentId int) (int64, error) {
	query := "UPDATE comments SET likesCount = likesCount - 1 WHERE id = ?"
	stmt, err := Database.Prepare(query)
	if err != nil {
		return 0, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(commentId)
	if err != nil {
		return 0, fmt.Errorf("failed to execute statement: %w", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("failed to get rows affected: %w", err)
	}
	return rows, nil
}

// add 1 to the dislikes count of the comment
func DislikeComment(commentId int) (int64, error) {
	query := "UPDATE comments SET dislikesCount = dislikesCount + 1 WHERE id = ?"
	stmt, err := Database.Prepare(query)
	if err != nil {
		return 0, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(commentId)
	if err != nil {
		return 0, fmt.Errorf("failed to execute statement: %w", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("failed to get rows affected: %w", err)
	}
	return rows, nil
}

// remove 1 from the dislikes count of the comment
func RemoveDislikeComment(commentId int) (int64, error) {
	stmt, err := Database.Prepare("UPDATE comments SET dislikesCount = dislikesCount - 1 WHERE id = ?")
	if err != nil {
		return 0, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(commentId)
	if err != nil {
		return 0, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("failed to get rows affected: %w", err)
	}
	return rows, nil
}

// get the likes and dislikes count of the comment
func GetLikesAndDislikeComment(commentId int) (int, int, error) {
	stmt, err := Database.Prepare("SELECT likesCount, dislikesCount FROM comments WHERE id = ?")
	if err != nil {
		return 0, 0, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()
	rows := stmt.QueryRow(commentId)
	var likes, dislikes int
	err = rows.Scan(&likes, &dislikes)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to scan rows: %w", err)
	}
	return likes, dislikes, nil
}

// update the liked status of the comment
func UpdateLikedComment(commentId, userId int, status bool) error {
	stmt, err := Database.Prepare("UPDATE engagement SET liked = ? disliked =? WHERE commentId = ? AND userId = ?")
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(status, false, commentId, userId)
	if err != nil {
		return fmt.Errorf("failed to execute statement: %w", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rows == 0 {
		stmt, err := Database.Prepare("INSERT INTO engagement (commentId, userId, liked, disliked) VALUES (?, ?, ?, ?)")
		if err != nil {
			return fmt.Errorf("failed to prepare statement: %w", err)
		}
		defer stmt.Close()
		_, err = stmt.Exec(commentId, userId, status, false)
		if err != nil {
			return fmt.Errorf("failed to execute statement: %w", err)
		}
	}
	return nil
}

// update the disliked status of the comment
func UpdateDislikedComment(commentId, userId int, status bool) error {
	stmt, err := Database.Prepare("UPDATE engagement SET liked = ? disliked =? WHERE commentId = ? AND userId = ?")
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(false, status, commentId, userId)
	if err != nil {
		return fmt.Errorf("failed to execute statement: %w", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rows == 0 {
		stmt, err := Database.Prepare("INSERT INTO engagement (commentId, userId, liked, disliked) VALUES (?, ?, ?, ?)")
		if err != nil {
			return fmt.Errorf("failed to prepare statement: %w", err)
		}
		defer stmt.Close()
		_, err = stmt.Exec(commentId, userId, false, status)
		if err != nil {
			return fmt.Errorf("failed to execute statement: %w", err)
		}
	}
	return nil
}
