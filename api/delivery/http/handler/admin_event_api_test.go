package handler


import (
"bytes"
//"encoding/json"
//"github.com/api/entity"
"github.com/julienschmidt/httprouter"
"net/http"
"net/http/httptest"
"testing"
//"github.com/stretchr/testify/assert"
"github.com/aleale2121/Hotel-Final/api/events/events_repository"
"github.com/aleale2121/Hotel-Final/api/events/events_services"
)

func TestGetEvents(t *testing.T) {

	EventsRepo := events_repository.NewMockEventsRepo(nil)
	EventsServ := events_services.NewEventService(EventsRepo)
	handler := NewAdminEventsHandlerApi(EventsServ)
	router := httprouter.New()
	router.GET("/hotel/events", handler.GetEvents)

	req, _ := http.NewRequest("GET", "/hotel/events", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status")
	}
	if rr.Code != http.StatusOK {
		t.Errorf("Response code is %v", rr.Code)
	}
	//var post=entity.News{
	//	Id:          1,
	//	Header:      "Mock newss 01",
	//	Description: "two newss",
	//	Image:       "tutu.png",
	//}
	//
	//json.Unmarshal(rr.Body.Bytes(), &post)
	//if post.Id != 1 {
	//	t.Errorf("Cannot retrieve JSON News")
	//}


}
func TestGetEventsById(t *testing.T) {
	EventsRepo := events_repository.NewMockEventsRepo(nil)
	EventsServ := events_services.NewEventService(EventsRepo)
	handler := NewAdminEventsHandlerApi(EventsServ)
	router := httprouter.New()
	router.GET("/hotel/events/:id", handler.GetEventById)

	req, _ := http.NewRequest("GET", "/hotel/events/1", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status")
	}
	if rr.Code != 200 {
		t.Errorf("Response code is %v", rr.Code)
	}
	//var post=entity.News{
	//	Id:          1,
	//	Header:      "Mock newss 01",
	//	Description: "two newss",
	//	Image:       "tutu.png",
	//}
	//
	//json.Unmarshal(rr.Body.Bytes(), &post)
	//if post.Id != 1 {
	//	t.Errorf("Cannot retrieve JSON News")
	//}
}

func TestPostEvents(t *testing.T) {

	var jsonStr = []byte(`{"id": 74,
        "header": "my newss",
        "description": "this is newss",
        "image": "hh.jpg"}`)
	EventsRepo := events_repository.NewMockEventsRepo(nil)
	EventsServ := events_services.NewEventService(EventsRepo)
	handler := NewAdminEventsHandlerApi(EventsServ)

	req, err := http.NewRequest("POST", "/hotel/events", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.POST("/hotel/events/", handler.PostEvents)
	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusPermanentRedirect {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusPermanentRedirect)
	}

}
//func TestEditEvents(t *testing.T) {
//
//	var jsonStr = []byte(`{"id": 74,
//        "header": "my newss",
//        "description": "this is newss",
//        "image": "hh.jpg"}`)
//	EventsRepo := events_repository.NewMockEventsRepo(nil)
//	EventsServ := events_services.NewEventService(EventsRepo)
//	handler := NewAdminEventsHandlerApi(EventsServ)
//	req, err := http.NewRequest("PUT", "/hotel/events:id", bytes.NewBuffer(jsonStr))
//	if err != nil {
//		t.Fatal(err)
//	}
//	req.Header.Set("Content-Type", "application/json")
//	rr := httptest.NewRecorder()
//	router := httprouter.New()
//	router.PUT("/hotel/events/4", handler.PutEvents)
//	router.ServeHTTP(rr, req)
//	if status := rr.Code; status != http.StatusOK {
//		t.Errorf("handler returned wrong status code: got %v want %v",
//			status, http.StatusOK)
//	}
//	//expected := `{"id":4,"first_name":"xyz change","last_name":"pqr","email_address":"xyz@pqr.com","phone_number":"1234567890"}`
//	//if rr.Body.String() != expected {
//	//	t.Errorf("handler returned unexpected body: got %v want %v",
//	//		rr.Body.String(), expected)
//	//}
//}
func TestDeleteEvents(t *testing.T) {
	EventsRepo := events_repository.NewMockEventsRepo(nil)
	EventsServ := events_services.NewEventService(EventsRepo)
	handler := NewAdminEventsHandlerApi(EventsServ)
	router := httprouter.New()
	router.DELETE("/hotel/events/:id", handler.DeleteEvents)

	req, _ := http.NewRequest("DELETE", "/hotel/events/1", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status code it is %v",rr.Code)
	}


}