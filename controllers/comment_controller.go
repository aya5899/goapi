package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/aya5899/goapi/apperrors"
	"github.com/aya5899/goapi/controllers/services"
	"github.com/aya5899/goapi/models"
)

type CommentController struct {
	service services.CommentServicer
}

func NewCommentController(s services.CommentServicer) *CommentController {
	return &CommentController{service: s}
}

// POST /comment
func (c *CommentController) PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		http.Error(w, "failed to encode json\n", http.StatusInternalServerError)
		return
	}
	comment, err := c.service.PostCommentService(reqComment)
	if err != nil {
		http.Error(w, "internal exec failure\n", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(comment)
}
