
package handler

import (
	"github.com/aleale2121/Hotel-Final/entity"
	"github.com/yuidegm/Hotel-Rental-Managemnet-System/events"
	"html/template"
	"net/http"
)

// newsHandler handles menu related requests
type eventHandler struct {
	tmpl          *template.Template
	eventsService events.EventsService
	userHandler *UserHandler
}
type eventsData struct {
	IsLogged string
	Events []entity.Events
}

// NewEventsHandler initializes and returns new newsHandler
func NewEventsHandler(T *template.Template, CS events.EventsService,userHandler *UserHandler) *eventHandler {
	return &eventHandler{tmpl: T, eventsService: CS,userHandler:userHandler}
}

func (mh *eventHandler) Event_page(w http.ResponseWriter, r *http.Request) {
	isLogged :="false"
	if mh.userHandler.loggedIn(r) {
		isLogged="true"
	}
	events, err := mh.eventsService.Events()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	eventData:=eventsData{
		IsLogged: isLogged,
		Events:   events,
	}
	_ = mh.tmpl.ExecuteTemplate(w, "event.layout", eventData)
}
