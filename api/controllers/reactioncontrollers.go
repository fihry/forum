package controllers

import (
	"fmt"

	"forum/api/models"
)

// like and dislike the post
func LikePost(postId int) (int64, error) {
	stmt, err := Database.Prepare("UPDATE posts SET likesCount = likesCount + 1 WHERE id = ?")
	if err != nil {
		fmt.Println("Error preparing", err)
		return 0, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(postId)
	if err != nil {
		fmt.Println("error executing", err)
		return 0, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		fmt.Println("error in rows affected", err)
		return 0, err
	}
	return rows, nil
}

// remove like from the post
func RemoveLike(postId int) (int64, error) {
	stmt, err := Database.Prepare("UPDATE posts SET likesCount = likesCount - 1 WHERE id = ?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(postId)
	if err != nil {
		return 0, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rows, nil
}

func DislikePost(postId int) (int64, error) {
	stmt, err := Database.Prepare("UPDATE posts SET lislikesCount = dislikesCount + 1 WHERE id = ?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(postId)
	if err != nil {
		return 0, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rows, nil
}

// remove dislike from the post
func RemoveDislike(postId int) (int64, error) {
	stmt, err := Database.Prepare("UPDATE posts SET dislikesCount = dislikesCount - 1 WHERE id = ?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(postId)
	if err != nil {
		return 0, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rows, nil
}

func GetPostByEngagement(userId int, Post models.Poste) models.Poste {
	query := `SELECT like, dislike FROM engagement WHERE userId = ? AND postId = ?`
	Database.QueryRow(query, userId, Post.ID).Scan(&Post.Liked, &Post.Disliked)
	return Post
}

func GetLIkesAndDislike(postId int) (int, int, error) {
	stmt, err := Database.Prepare("SELECT likes, dislikes FROM posts WHERE id = ?")
	if err != nil {
		return 0, 0, err
	}
	defer stmt.Close()
	var likes, dislikes int
	err = stmt.QueryRow(postId).Scan(&likes, &dislikes)
	if err != nil {
		return 0, 0, err
	}
	return likes, dislikes, nil
}

func AddLikeToEngament(postId, userId int) error {
	stmt, err := Database.Prepare("INSERT INTO engagement (postId, userId, like, dislike) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(postId, userId, true, false)
	if err != nil {
		return err
	}
	return nil
}

func UpdateLikeToEngament(postId, userId int) error {
	stmt, err := Database.Prepare("UPDATE engagement SET like =? dislike =? WHERE postId =? AND userId =?")
	if err != nil {
		return AddDislikeToEngament(postId, userId)
	}
	defer stmt.Close()
	_, err = stmt.Exec(true, false, postId, userId)
	if err != nil {
		return err
	}
	return nil
}

func AddDislikeToEngament(postId, userId int) error {
	stmt, err := Database.Prepare("INSERT INTO engagement (postId, userId, like, dislike) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(postId, userId,false, true)
	if err != nil {
		return err
	}
	return nil
}

func UpdateDislikeToEngagement(postId, userId int) error {
	stmt, err := Database.Prepare("UPDATE engagement SET dislike =? WHERE postId =? AND userId =?")
	if err != nil {
		return AddDislikeToEngament(postId, userId)
	}
	defer stmt.Close()
	_, err = stmt.Exec(true, postId, userId)
	if err != nil {
		return err
	}
	return nil
}

func RemLikeFromEngagement(postId, userId int) error {
	stmt, err := Database.Prepare("UPDATE engagement SET like =? WHERE postId = ? AND userId = ?")
	if err != nil {
		fmt.Println("error in prepare statement", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(false, postId, userId)
	if err != nil {
		fmt.Println("error in execute statement", err)
		return err
	}
	return nil
}

func RemDislikeFromEngagement(postId, userId int) error {
	stmt, err := Database.Prepare("UPDATE engagement SET dislike =? WHERE postId = ? AND userId = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(true, postId, userId)
	if err != nil {
		return err
	}
	return nil
}
