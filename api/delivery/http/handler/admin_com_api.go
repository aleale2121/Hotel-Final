package handler

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/yuidegm/Hotel-Rental-Managemnet-System/api/comment"
	"net/http"
	"strconv"
)

// AdminComHandlerApi handles comment related http requests
type AdminComHandlerApi struct {
	commentService comment.CommentServices
}

// NewAdminCommentHandler returns new AdminComHandlerApi object
func NewAdminComHandlerApi(cmntService comment.CommentServices) *AdminComHandlerApi {
	return &AdminComHandlerApi{commentService: cmntService}
}

// GetEvents handles GET /v1/admin/comments request
func (ach *AdminComHandlerApi) GetCom(w http.ResponseWriter,
	r *http.Request,_ httprouter.Params) {

	comments, errs := ach.commentService.Comments()


	if len(errs) > 0 {


		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(comments, "", "\t\t")


	if err != nil {


		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}


// DeleteEvents handles DELETE /v1/admin/comments/:id request
func (ach *AdminComHandlerApi) DeleteCom(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, errs := ach.commentService.DeleteCom((id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return
}

