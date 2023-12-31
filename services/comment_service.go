package services

import (
	"github.com/aya5899/goapi/apperrors"
	"github.com/aya5899/goapi/models"
	"github.com/aya5899/goapi/repositories"
)

func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {
	// databaseへのコメントへの挿入
	newComment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "failed to record data")
		return models.Comment{}, err
	}

	return newComment, nil
}
