package handler

import (
	"github.com/yuidegm/Hotel-Rental-Managemnet-System/entity"
	"github.com/yuidegm/Hotel-Rental-Managemnet-System/news"
	"html/template"
	"net/http"
)

// newsHandler handles menu related requests
type newsHandler struct {
	tmpl        *template.Template
	newsService news.NewsService
	userHandler *UserHandler
}
type newsData struct {
	IsLogged string
    News []entity.News
}
// NewEventsHandler initializes and returns new newsHandler
func NewMenuHandler(T *template.Template, CS news.NewsService,userHandler *UserHandler) *newsHandler {
	return &newsHandler{tmpl: T, newsService: CS,userHandler:userHandler}
}


func (mh *newsHandler) News_page(w http.ResponseWriter, r *http.Request) {
	isLogged :="false"
	if mh.userHandler.loggedIn(r) {
		isLogged="true"
	}
	news, err := mh.newsService.News()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
		}
		newData:=newsData{
			IsLogged: isLogged,
			News:     news,
		}
	 mh.tmpl.ExecuteTemplate(w, "news_layout", newData)
}
