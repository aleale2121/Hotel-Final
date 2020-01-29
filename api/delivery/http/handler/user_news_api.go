package handler

import (
	"fmt"
	"github.com/yuidegm/Hotel-Rental-Managemnet-System/api/news"
	"html/template"
	"net/http"
)

// newsHandler handles menu related requests
type newsHandler struct {
	tmpl        *template.Template
	newsService news.NewsService
}

// NewEventsHandler initializes and returns new newsHandler
func NewMenuHandler(T *template.Template, CS news.NewsService) *newsHandler {
	return &newsHandler{tmpl: T, newsService: CS}
}

// Index handles request on route /
func (mh *newsHandler) Index(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	news, err := mh.newsService.News()
	if err != nil {
		panic(err)
	}

	mh.tmpl.ExecuteTemplate(w, "index.layout", news)
}

func (mh *newsHandler) News_page(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/News" {
		http.NotFound(w, r)
		fmt.Printf("ayyyy---1")
		return

	}

	news, err := mh.newsService.News()
	if err != nil {
		panic(err)
	}
    fmt.Printf("here to excute")
	mh.tmpl.ExecuteTemplate(w, "news_layout", news)
	fmt.Printf("it must excute perfectly")
}
