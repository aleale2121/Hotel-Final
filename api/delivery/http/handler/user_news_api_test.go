package handler

//import (
//	"encoding/json"
//	"github.com/julienschmidt/httprouter"
//	"github.com/yuidegm/Hotel-Rental-Managemnet-System/api/entity"
//	"github.com/yuidegm/Hotel-Rental-Managemnet-System/api/newss/news_repository"
//	"github.com/yuidegm/Hotel-Rental-Managemnet-System/api/newss/news_services"
//	"net/http"
//	"net/http/httptest"
//	"testing"
//)
//
//func TestUserNewsPage(t *testing.T) {
//
//
//	EventsRepo := news_repository.NewMockNewsRepo(nil)
//	NewsServ := news_services.NewNewsService(EventsRepo)
//	handler := NewMenuHandler(NewsServ)
//
//	router := httprouter.New()
//	router.GET("/newss", handler.News_page)
//
//	req, _ := http.NewRequest("GET", "/newss", nil)
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
//		Header:      "Mock newss 01",
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
