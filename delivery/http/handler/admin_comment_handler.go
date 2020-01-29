package handler

import (
	"github.com/aleale2121/Hotel-Final/comment"
)

// AdminCommentHandler handles comment related http requests
type AdminCommentHandler struct {
	commentService comment.CommentServices
}

// NewAdminCommentHandler returns new AdminCommentHandler object
func NewAdminCommentHandler(cmntService comment.CommentServices) *AdminCommentHandler {
	return &AdminCommentHandler{commentService: cmntService}
}

