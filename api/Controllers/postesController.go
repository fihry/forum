package Controllers

import (
	"forum/api/Models"
)

func (D *Database) GetAllPosts() ([]Models.Poste, error) {
	stmt, err := D.DB.Prepare("SELECT * FROM posts")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
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
	stmt, err := D.DB.Prepare("SELECT * FROM posts WHERE id = ?")
	if err != nil {
		return poste, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(P.ID).Scan(
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

func (D *Database) CreatePoste(P Models.Poste) (int64, error) {
	stmt, err := D.DB.Prepare("INSERT INTO posts (title, content, author, category) VALUES (?, ?, ?, ?)")
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

// func (D *Database) UpdatePoste(P Models.Poste) (int64, error) {
// 	stmt, err := D.DB.Prepare("UPDATE posts SET title = ?, content = ?, author = ?, category = ? WHERE id = ?")
// 	if err != nil {
// 		return 0, err
// 	}
// 	defer stmt.Close()
// 	result, err := stmt.Exec(P.Title, P.Content, P.Author, P.Category, P.ID)
// 	if err != nil {
// 		return 0, err
// 	}
// 	rows, err := result.RowsAffected()
// 	if err != nil {
// 		return 0, err
// 	}
// 	return rows, nil
// }

// func (D *Database) DeletePoste(P Models.Poste) (int64, error) {
// 	stmt, err := D.DB.Prepare("DELETE FROM posts WHERE id = ?")
// 	if err != nil {
// 		return 0, err
// 	}
// 	defer stmt.Close()
// 	result, err := stmt.Exec(P.ID)
// 	if err != nil {
// 		return 0, err
// 	}
// 	rows, err := result.RowsAffected()
// 	if err != nil {
// 		return 0, err
// 	}
// 	return rows, nil
// }

func (D *Database) GetPostsByCategory(category string) ([]Models.Poste, error) {
	stmt, err := D.DB.Prepare("SELECT * FROM posts WHERE category = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(category)
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

func (D *Database) GetPostsByAuthor(author string) ([]Models.Poste, error) {
	stmt, err := D.DB.Prepare("SELECT * FROM posts WHERE author = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(author)
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

func (D *Database) AddComment(C Models.Comment) (int64, error) {
	stmt, err := D.DB.Prepare("INSERT INTO comments (content, author, post_id) VALUES (?, ?, ?)")
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

func (D *Database) GetCommentsByPostId(id int) ([]Models.Comment, error) {
	stmt, err := D.DB.Prepare("SELECT * FROM comments WHERE post_id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	comments := []Models.Comment{}
	for rows.Next() {
		comment := Models.Comment{}
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

// like and dislike the post
func (D *Database) LikePost(P Models.Poste) (int64, error) {
	stmt, err := D.DB.Prepare("UPDATE posts SET likes = likes + 1 WHERE id = ?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(P.ID)
	if err != nil {
		return 0, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rows, nil
}

// remove like from the post
func (d *Database) RemoveLike(P Models.Poste) (int64, error) {
	stmt, err := d.DB.Prepare("UPDATE posts SET likes = likes - 1 WHERE id = ?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(P.ID)
	if err != nil {
		return 0, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rows, nil
}

func (D *Database) DislikePost(P Models.Poste) (int64, error) {
	stmt, err := D.DB.Prepare("UPDATE posts SET dislikes = dislikes + 1 WHERE id = ?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(P.ID)
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
func (d *Database) RemoveDislike(P Models.Poste) (int64, error) {
	stmt, err := d.DB.Prepare("UPDATE posts SET dislikes = dislikes - 1 WHERE id = ?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(P.ID)
	if err != nil {
		return 0, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rows, nil
}

func (D *Database) GetLIkesAndDislike(P Models.Poste) (int, int, error) {
	stmt, err := D.DB.Prepare("SELECT likes, dislikes FROM posts WHERE id = ?")
	if err != nil {
		return 0, 0, err
	}
	defer stmt.Close()
	var likes, dislikes int
	err = stmt.QueryRow(P.ID).Scan(&likes, &dislikes)
	if err != nil {
		return 0, 0, err
	}
	return likes, dislikes, nil
}
