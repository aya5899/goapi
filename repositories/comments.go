package repositories

import (
	"database/sql"
	"fmt"

	"github.com/aya5899/goapi/models"
)

func InsertComment(db *sql.DB, comment models.Comment) (models.Comment, error) {
	const sqlStr = `
		insert into comments (article_id, message, created_at)
		values (?, ?, now())
		`
	var newComment models.Comment
	newComment.ArticleID, newComment.Message = comment.ArticleID, comment.Message
	// コメントの挿入
	result, err := db.Exec(sqlStr, newComment.ArticleID, newComment.Message)
	if err != nil {
		fmt.Println(err)
		return models.Comment{}, err
	}
	id, _ := result.LastInsertId()
	newComment.CommentID = int(id)

	return newComment, nil
}

func SelectCommentList(db *sql.DB, articleID int) ([]models.Comment, error) {
	const sqlStr = `
		select *
		from comments
		where article_id = ?;
		`

	rows, err := db.Query(sqlStr, articleID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()
	commentArray := make([]models.Comment, 0)
	for rows.Next() {
		var comment models.Comment
		var createdTime sql.NullTime
		err = rows.Scan(&comment.ArticleID, &comment.CommentID, &comment.Message, &createdTime)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		if createdTime.Valid {
			comment.CreatedAt = createdTime.Time
		}
		commentArray = append(commentArray, comment)
	}
	return commentArray, nil
}
