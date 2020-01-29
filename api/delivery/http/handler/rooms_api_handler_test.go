

package handler


//func TestGetRooms(t *testing.T) {
//	EventsRepo := news_repository.NewMockNewsRepo(nil)
//	EventsServ := news_services.NewNewsService(EventsRepo)
//	handler := NewAdminNewsHandlerApi(EventsServ)
//	router := httprouter.New()
//	router.GET("/hotel/newss", handler.GetNews)
//
//	req, _ := http.NewRequest("GET", "/hotel/newss", nil)
//	rr := httptest.NewRecorder()
//
//	router.ServeHTTP(rr, req)
//	if status := rr.Code; status != http.StatusOK {
//		t.Errorf("Wrong status")
//	}
//	if rr.Code != 200 {
//		t.Errorf("Response code is %v", rr.Code)
//	}
//	var post=entity.News{
//		Id:          1,
//		Header:      "Mock newss 01",
//		Description: "two newss",
//		Image:       "tutu.png",
//	}
//
//	json.Unmarshal(rr.Body.Bytes(), &post)
//	if post.Id != 1 {
//		t.Errorf("Cannot retrieve JSON News")
//	}
//
//
//}
//func TestGetNewsById(t *testing.T) {
//	EventsRepo := news_repository.NewMockNewsRepo(nil)
//	EventsServ := news_services.NewNewsService(EventsRepo)
//	handler := NewAdminNewsHandlerApi(EventsServ)
//	router := httprouter.New()
//	router.GET("/hotel/newss/:id", handler.GetNewsById)
//
//	req, _ := http.NewRequest("GET", "/hotel/newss/1", nil)
//	rr := httptest.NewRecorder()
//
//	router.ServeHTTP(rr, req)
//	if status := rr.Code; status != http.StatusOK {
//		t.Errorf("Wrong status")
//	}
//	if rr.Code != 200 {
//		t.Errorf("Response code is %v", rr.Code)
//	}
//	//var post=entity.News{
//	//	Id:          1,
//	//	Header:      "Mock newss 01",
//	//	Description: "two newss",
//	//	Image:       "tutu.png",
//	//}
//	//
//	//json.Unmarshal(rr.Body.Bytes(), &post)
//	//if post.Id != 1 {
//	//	t.Errorf("Cannot retrieve JSON News")
//	//}
//}
//
//func TestPostNews(t *testing.T) {
//
//	var jsonStr = []byte(`{"id": 74,
//        "header": "my newss",
//        "description": "this is newss",
//        "image": "hh.jpg"}`)
//	EventsRepo := news_repository.NewMockNewsRepo(nil)
//	EventsServ := news_services.NewNewsService(EventsRepo)
//	handler := NewAdminNewsHandlerApi(EventsServ)
//
//	req, err := http.NewRequest("POST", "/hotel/newss", bytes.NewBuffer(jsonStr))
//	if err != nil {
//		t.Fatal(err)
//	}
//	req.Header.Set("Content-Type", "application/json")
//	rr := httptest.NewRecorder()
//	router := httprouter.New()
//	router.POST("/hotel/newss/", handler.PostNews)
//	router.ServeHTTP(rr, req)
//	if status := rr.Code; status != http.StatusPermanentRedirect {
//		t.Errorf("handler returned wrong status code: got %v want %v",
//			status, http.StatusPermanentRedirect)
//	}
//
//}
//func TestDeleteNews(t *testing.T) {
//	EventsRepo := news_repository.NewMockNewsRepo(nil)
//	EventsServ := news_services.NewNewsService(EventsRepo)
//	handler := NewAdminNewsHandlerApi(EventsServ)
//	router := httprouter.New()
//	router.DELETE("/hotel/newss/:id", handler.DeleteNews)
//
//	req, _ := http.NewRequest("DELETE", "/hotel/newss/1", nil)
//	rr := httptest.NewRecorder()
//
//	router.ServeHTTP(rr, req)
//	if status := rr.Code; status != http.StatusOK {
//		t.Errorf("Wrong status code it is %v",rr.Code)
//	}
//
//
//}