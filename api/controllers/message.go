package controllers

import (
	"fmt"

	"forum/api/models"
)

func GetCommentsForPost(postId, userId int) ([]models.Comment, error) {
	// get all the comments of the post
	comments, err := GetAllComments(postId, userId)
	if err != nil {
		return nil, err
	}
    // check if the user session is valid 
    // if !IsUserValid(userId) {
    //     return comments, nil
    // }
	// get all the likes and dislikes for the past comment
	for i, comment := range comments {
		liked, disliked, err := GetCommentReaction(comment.ID, userId)
		if err != nil {
			return nil, err
		}
		comments[i].Liked = &liked
		comments[i].Disliked = &disliked
	}

	return comments, nil
}

func GetCommentReaction(commentId, userId int) (bool, bool, error) {
	stmt, err := Database.Prepare("SELECT liked, disliked FROM engagement WHERE commentId =? AND userId =?")
    if err!= nil {
        fmt.Println("error preparing", err)
        return false, false, err
    }
    defer stmt.Close()
    rows, err := stmt.Query(commentId, userId)
    if err!= nil {
        fmt.Println("error querying", err)
        return false, false, err
    }
    defer rows.Close()

    var liked, disliked bool
    err = rows.Scan(&liked, &disliked)
    if err!= nil {
        fmt.Println("error scanning", err)
        return false, false, err
    }
    return liked, disliked, nil
}

func GetAllComments(postId, userId int) ([]models.Comment, error) {
	stmt, err := Database.Prepare("SELECT * FROM comments WHERE postId =?")
    if err!= nil {
        fmt.Println("error preparing", err)
        return nil, err
    }
    defer stmt.Close()
    rows, err := stmt.Query(postId)
    if err!= nil {
        fmt.Println("error querying", err)
        return nil, err
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
        if err!= nil {
            fmt.Println("error scanning", err)
            return nil, err
        }
        comments = append(comments, comment)
	}
	return comments, nil
}

