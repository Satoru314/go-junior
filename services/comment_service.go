package services

import (
	"myapi/apperrors"
	"myapi/models"
	"myapi/repositories"
)

func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {
	resComment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to recode data")
		return models.Comment{}, err
	}
	return resComment, nil
}
