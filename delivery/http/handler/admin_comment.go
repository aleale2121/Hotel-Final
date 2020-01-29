
package handler

import (
	"github.com/aleale2121/Hotel-Final/comment"
	"html/template"

	"net/http"

	"strconv"
)

// AdminNewsHandler handles category handler admin requests
type AdminComHandler struct {
	tmpl        *template.Template
	newsService comment.CommentServices
}

// NewAdminNewsHandler initializes and returns new AdminCateogryHandler
func NewAdminComHandler(T *template.Template, NS comment.CommentServices) *AdminComHandler {
	return &AdminComHandler{tmpl: T, newsService: NS}
}

// AdminNews handle requests on route /admin/newss
func (ach *AdminComHandler) AdminCom(w http.ResponseWriter, r *http.Request) {
	news, err := ach.newsService.Comment()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	ach.tmpl.ExecuteTemplate(w, "admin.com.layout", news)
}

// AdminNewsDelete handle requests on route /admin/categories/delete
func (ach *AdminComHandler) AdminComDelete(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")

		id, err := strconv.Atoi(idRaw)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		_,err2 := ach.newsService.DeleteCom(id)

		if len(err2)>0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

	}

	http.Redirect(w, r, "/admin/com", http.StatusSeeOther)
}
