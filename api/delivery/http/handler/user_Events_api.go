
package handler

import (
	"fmt"
	"github.com/aleale2121/Hotel-Final/api/events"
	"html/template"
	"net/http"
)

// newsHandler handles menu related requests
type eventHandler struct {
	tmpl          *template.Template
	eventsService events.EventsService
}

// NewEventsHandler initializes and returns new newsHandler
func NewEventsHandler(T *template.Template, CS events.EventsService) *eventHandler {
	return &eventHandler{tmpl: T, eventsService: CS}
}

func (mh *eventHandler) Event_page(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/Events" {
		http.NotFound(w, r)
		fmt.Printf("ayyyy---1")
		return
	}
	events, err := mh.eventsService.Events()
	if err != nil {
		panic(err)
	}
	fmt.Printf("here to excute")
	mh.tmpl.ExecuteTemplate(w, "event.layout", events)
	fmt.Printf("it must excute perfectly")
}
