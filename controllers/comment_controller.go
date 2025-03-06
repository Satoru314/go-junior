package controllers

import (
	"encoding/json"
	"io"
	"myapi/apperrors"
	"myapi/controllers/services"
	"myapi/models"
	"net/http"
)

// CommentControllerの定義
type CommentController struct {
	service services.CommentServicer
}

// 外部でのCommentControllerの作成
func NewCommentController(s services.CommentServicer) *CommentController {
	return &CommentController{service: s}
}

// 外部からのCommentControllerの利用
func (c *CommentController) CommentHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Comment...\n")
	var reqComment models.Comment
	err := json.NewDecoder(req.Body).Decode(&reqComment)
	if err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		apperrors.ErrorHandler(w, req, err)
		return
	}
	resComment, err := c.service.PostCommentService(reqComment)
	if err != nil {
		err = apperrors.ServiceFuncFailed.Wrap(err, "fail to post comment service")
		apperrors.ErrorHandler(w, req, err)
		return
	}
	json.NewEncoder(w).Encode(resComment)
}
