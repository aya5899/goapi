package services

import (
	"github.com/aya5899/goapi/models"
	"github.com/aya5899/goapi/repositories"
)

func PostCommentService(comment models.Comment) (models.Comment, error) {
	// databaseへの接続
	db, err := connectDB()
	if err != nil {
		return models.Comment{}, err
	}
	defer db.Close()
	// databaseへのコメントへの挿入
	newComment, err := repositories.InsertComment(db, comment)
	if err != nil {
		return models.Comment{}, err
	}

	return newComment, nil
}
