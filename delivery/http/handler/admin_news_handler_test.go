package handler

import (
	"bytes"
	"github.com/yuidegm/Hotel-Rental-Managemnet-System/entity"
	"github.com/yuidegm/Hotel-Rental-Managemnet-System/news/news_repository"
	"github.com/yuidegm/Hotel-Rental-Managemnet-System/news/news_services"
	"html/template"
	"io/ioutil"

	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

var csrfSignKeyFake =[]byte("fakecsrf")
func TestAdminNews(t *testing.T) {

	tmpl := template.Must(template.ParseGlob("C:/Users/getch/go/src/github.com/yuidegm/Hotel-Rental-Managemnet-System/ui/templates/admin_template/*"))

	EventsRepo := news_repository.NewMockNewsRepo(nil)
	EventsServ := news_services.NewGormNewsServiceImpl(EventsRepo)

	adminEventsHandler := NewAdminNewsHandler(tmpl, EventsServ,csrfSignKeyFake)

	mux := http.NewServeMux()
	mux.HandleFunc("/admin/newss", adminEventsHandler.AdminNews)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	url := ts.URL
	resp, err := tc.Get(url + "/admin/newss")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()
	//
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Contains(body, []byte("Mock News 01")) {

		t.Errorf("want body to contain %q", body)
	}
}

func TestAdminNewsNew(t *testing.T) {

	tmpl := template.Must(template.ParseGlob("C:/Users/getch/go/src/github.com/yuidegm/Hotel-Rental-Managemnet-System/ui/templates/admin_template/*"))

	EventsRepo := news_repository.NewMockNewsRepo(nil)
	EventsServ := news_services.NewGormNewsServiceImpl(EventsRepo)

	adminEventsHandler := NewAdminNewsHandler(tmpl, EventsServ,csrfSignKeyFake)

	mux := http.NewServeMux()
	mux.HandleFunc("/admin/newss/new", adminEventsHandler.AdminNews)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	sURL := ts.URL

	form := url.Values{}
	form.Add("name", entity.NewsMock.Header)
	form.Add("Description", entity.NewsMock.Description)
	form.Add("Image", entity.NewsMock.Image)

	resp, err := tc.PostForm(sURL+"/admin/newss/new", form)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Contains(body, []byte("Mock News 01")) {
		t.Errorf("want body to contain %q", body)
	}

}

func TestAdminNewsUpdate(t *testing.T) {

	tmpl := template.Must(template.ParseGlob("C:/Users/getch/go/src/github.com/yuidegm/Hotel-Rental-Managemnet-System/ui/templates/admin_template/*"))

	EventsRepo := news_repository.NewMockNewsRepo(nil)
	EventsServ := news_services.NewGormNewsServiceImpl(EventsRepo)

	adminEventsHandler := NewAdminNewsHandler(tmpl, EventsServ,csrfSignKeyFake)

	mux := http.NewServeMux()
	mux.HandleFunc("/admin/newss/update", adminEventsHandler.AdminNews)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	sURL := ts.URL

	form := url.Values{}

	form.Add("Id", string(entity.NewsMock.Id))
	form.Add("Name", entity.NewsMock.Header)
	form.Add("kescription", entity.NewsMock.Description)
	form.Add("Image", entity.NewsMock.Image)

	resp, err := tc.PostForm(sURL+"/admin/newss/update?id=1", form)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Contains(body, []byte("Mock News 01")) {
		t.Errorf("want body to contain %q", body)
	}

}

func TestAdminNewsDelete(t *testing.T) {

	tmpl := template.Must(template.ParseGlob("C:/Users/getch/go/src/github.com/yuidegm/Hotel-Rental-Managemnet-System/ui/templates/admin_template/*"))

	EventsRepo := news_repository.NewMockNewsRepo(nil)
	EventsServ := news_services.NewGormNewsServiceImpl(EventsRepo)

	adminEventsHandler := NewAdminNewsHandler(tmpl, EventsServ,csrfSignKeyFake)

	mux := http.NewServeMux()
	mux.HandleFunc("/admin/newss/delete", adminEventsHandler.AdminNews)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	sURL := ts.URL

	form := url.Values{}

	form.Add("Id", string(entity.NewsMock.Id))
	form.Add("Name", entity.NewsMock.Header)
	form.Add("kescription", entity.NewsMock.Description)
	form.Add("Image", entity.NewsMock.Image)

	resp, err := tc.PostForm(sURL+"/admin/newss/delete?id=1", form)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Contains(body, []byte("Mock News 01")) {
		t.Errorf("want body to contain %q", body)
	}

}
