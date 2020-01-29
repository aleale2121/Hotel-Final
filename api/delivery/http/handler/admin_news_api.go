package handler

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/yuidegm/Hotel-Rental-Managemnet-System/api/entity"
	"github.com/yuidegm/Hotel-Rental-Managemnet-System/api/news"
	"net/http"
	"strconv"
)

// AdminEventsHandlerApi handles comment related http requests
type AdminNewsHandlerApi struct {
	commentService news.NewsService
}

// NewAdminCommentHandler returns new AdminEventsHandlerApi object
func NewAdminNewsHandlerApi(cmntService news.NewsService) *AdminNewsHandlerApi {
	return &AdminNewsHandlerApi{commentService: cmntService}
}

// GetEvents handles GET /v1/admin/comments request
func (ach *AdminNewsHandlerApi) GetNews(w http.ResponseWriter,
	r *http.Request, _ httprouter.Params) {

	comments, errs := ach.commentService.News()
	fmt.Println(comments,"hand")

	if len(errs) > 0 {
		fmt.Println(comments,"hand")

		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(comments, "", "\t\t")
	fmt.Println(comments,"hand")

	if err != nil {
		fmt.Println(comments,"hand")

		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}

// GetEventById handles GET /v1/admin/comments/:id request
func (ach *AdminNewsHandlerApi) GetNewsById(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {
	fmt.Println("handler")
	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		fmt.Println("handler1")

		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	comment, errs := ach.commentService.NewsById(uint(id))
	fmt.Println("handler2",comment)
	if len(errs) > 0 {
		fmt.Println("handler2",comment)

		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(comment, "", "\t\t")

	if err != nil {
		fmt.Println("handler3")

		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	fmt.Println("handler4",comment,output)

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

// PostEvents handles POST /v1/admin/comments request
func (ach *AdminNewsHandlerApi) PostNews(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	comment := &entity.News{}

	err := json.Unmarshal(body, comment)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	comment, errs := ach.commentService.StoreNews(comment)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	p := fmt.Sprintf("/hotel/newss/%d", comment.Id)
	w.Header().Set("Location", p)
	w.WriteHeader(http.StatusCreated)
	return
}

// PutEvents handles PUT /v1/admin/comments/:id request
func (ach *AdminNewsHandlerApi) PutNews(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	comment, errs := ach.commentService.NewsById(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	l := r.ContentLength

	body := make([]byte, l)

	r.Body.Read(body)

	json.Unmarshal(body, &comment)

	comment, errs = ach.commentService.UpdateNews(comment)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(comment, "", "\t\t")

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
func (ach *AdminNewsHandlerApi) DeleteNews(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, errs := ach.commentService.DeleteNews(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return
}
