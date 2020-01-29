package handler

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/aleale2121/Hotel-Final/api/comment/com_Repository"
	"github.com/aleale2121/Hotel-Final/api/comment/com_Service"
)

func TestGetCom(t *testing.T) {
	comRepo := com_Repository.NewMockComRepo(nil)
	comServ := com_Service.NewGormComServiceImpl(comRepo)
	handler := NewAdminComHandlerApi(comServ)
	router := httprouter.New()
	router.GET("/hotel/com", handler.GetCom)

	req, _ := http.NewRequest("GET", "/hotel/com", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status")
	}
	if rr.Code != http.StatusOK {
		t.Errorf("Response code is %v", rr.Code)
	}
	if rr.Code != http.StatusOK {
		fmt.Printf("Response code is %d", rr.Code)
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

func TestDeleteCom(t *testing.T) {
	comRepo := com_Repository.NewMockComRepo(nil)
	comServ := com_Service.NewGormComServiceImpl(comRepo)
	handler := NewAdminComHandlerApi(comServ)

	router := httprouter.New()
	router.DELETE("/hotel/com/:id", handler.DeleteCom)

	req, _ := http.NewRequest("DELETE", "/hotel/com/1", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status code it is %v",rr.Code)
	}


}
