package handler

import (
	"bytes"
	"github.com/aleale2121/Hotel-Final/entity"
	"github.com/aleale2121/Hotel-Final/events/events_repository"
	"github.com/aleale2121/Hotel-Final/events/events_services"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestAdminEvents(t *testing.T) {

	tmpl := template.Must(template.ParseGlob("C:/Users/getch/go/src/github.com/yuidegm/Hotel-Rental-Managemnet-System/ui/templates/admin_template/*"))

	EventsRepo := events_repository.NewMockEventsRepo(nil)
	EventsServ := events_services.NewGormNewsServiceImpl(EventsRepo)
	csrfSignKeyFake :=[]byte("alefew")
	adminEventsHandler := NewAdminEventsHandler(tmpl, EventsServ,csrfSignKeyFake)

	mux := http.NewServeMux()
	mux.HandleFunc("/admin/events", adminEventsHandler.AdminEvents)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	url := ts.URL

	resp, err := tc.Get(url + "/admin/events")
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

	if !bytes.Contains(body, []byte("Mock events 01")) {
		t.Errorf("want body to contain %q", body)
	}
}

func TestAdminEventsNew(t *testing.T) {

	tmpl := template.Must(template.ParseGlob("C:/Users/getch/go/src/github.com/yuidegm/Hotel-Rental-Managemnet-System/ui/templates/admin_template/*"))

	EventsRepo := events_repository.NewMockEventsRepo(nil)
	EventsServ := events_services.NewGormNewsServiceImpl(EventsRepo)

	adminEventsHandler := NewAdminEventsHandler(tmpl, EventsServ,csrfSignKeyFake)

	mux := http.NewServeMux()
	mux.HandleFunc("/admin/events/new", adminEventsHandler.AdminEvents)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	sURL := ts.URL

	form := url.Values{}
	form.Add("name", entity.EventsMock.Header)
	form.Add("Description", entity.EventsMock.Description)
	form.Add("Image", entity.EventsMock.Image)

	resp, err := tc.PostForm(sURL+"/admin/events/new", form)
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

	if !bytes.Contains(body, []byte("Mock events 01")) {
		t.Errorf("want body to contain %q", body)
	}

}

func TestAdminEventsUpdate(t *testing.T) {

	tmpl := template.Must(template.ParseGlob("C:/Users/getch/go/src/github.com/yuidegm/Hotel-Rental-Managemnet-System/ui/templates/admin_template/*"))

	EventsRepo := events_repository.NewMockEventsRepo(nil)
			EventsServ := events_services.NewGormNewsServiceImpl(EventsRepo)

	adminEventsHandler := NewAdminEventsHandler(tmpl, EventsServ,csrfSignKeyFake)

	mux := http.NewServeMux()
	mux.HandleFunc("/admin/events/update", adminEventsHandler.AdminEvents)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	sURL := ts.URL

	form := url.Values{}

	form.Add("Id", string(entity.EventsMock.Id))
	form.Add("Name", entity.EventsMock.Header)
	form.Add("kescription", entity.EventsMock.Description)
	form.Add("Image", entity.EventsMock.Image)

	resp, err := tc.PostForm(sURL+"/admin/events/update?id=1", form)
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

	if !bytes.Contains(body, []byte("Mock events 01")) {
		t.Errorf("want body to contain %q", body)
	}

}

func TestAdminEventsDelete(t *testing.T) {

	tmpl := template.Must(template.ParseGlob("C:/Users/getch/go/src/github.com/yuidegm/Hotel-Rental-Managemnet-System/ui/templates/admin_template/*"))

	EventsRepo := events_repository.NewMockEventsRepo(nil)
	EventsServ := events_services.NewGormNewsServiceImpl(EventsRepo)

	adminEventsHandler := NewAdminEventsHandler(tmpl, EventsServ,csrfSignKeyFake)

	mux := http.NewServeMux()
	mux.HandleFunc("/admin/events/delete", adminEventsHandler.AdminEvents)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	sURL := ts.URL

	form := url.Values{}

	form.Add("Id", string(entity.EventsMock.Id))
	form.Add("Name", entity.EventsMock.Header)
	form.Add("kescription", entity.EventsMock.Description)
	form.Add("Image", entity.EventsMock.Image)

	resp, err := tc.PostForm(sURL+"/admin/events/delete?id=1", form)
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

	if !bytes.Contains(body, []byte("Mock events 01")) {
		t.Errorf("want body to contain %q", body)
	}

}

