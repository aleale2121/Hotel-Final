package handler

//import (
//	"encoding/json"
//	"github.com/julienschmidt/httprouter"
//	"github.com/yuidegm/Hotel-Rental-Managemnet-System/api/entity"
//	"github.com/yuidegm/Hotel-Rental-Managemnet-System/api/events/events_repository"
//	"github.com/yuidegm/Hotel-Rental-Managemnet-System/api/events/events_services"
//	"net/http"
//	"net/http/httptest"
//	"testing"
//)
//
//func TestUserEventsPage(t *testing.T) {
//
//
//	EventsRepo := events_repository.NewMockEventsRepo(nil)
//	EventsServ := events_services.NewEventService(EventsRepo)
//	handler := NewuserEventHandler(EventsServ)
//
//	router := httprouter.New()
//	router.GET("/events", handler.Event_page)
//
//	req, _ := http.NewRequest("GET", "/events", nil)
//	rr := httptest.NewRecorder()
//
//	router.ServeHTTP(rr, req)
//	if status := rr.Code; status != http.StatusOK {
//		t.Errorf("Wrong status")
//	}
//	if rr.Code != 200 {
//		t.Errorf("Response code is %v", rr.Code)
//	}
//	var post=entity.Events{
//		Id:          1,
//		Header:      "Mock events 01",
//		Description: "two",
//		Image:       "tutu.png",
//	}
//
//	json.Unmarshal(rr.Body.Bytes(), &post)
//	if post.Id != 1 {
//		t.Errorf("Cannot retrieve JSON News")
//	}
//
//}
