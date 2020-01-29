package handler

import (
	"github.com/yuidegm/Hotel-Rental-Managemnet-System/comment"
)

// AdminCommentHandler handles comment related http requests
type AdminCommentHandler struct {
	commentService comment.CommentServices
}

// NewAdminCommentHandler returns new AdminCommentHandler object
func NewAdminCommentHandler(cmntService comment.CommentServices) *AdminCommentHandler {
	return &AdminCommentHandler{commentService: cmntService}
}

