package handler

import (
	"bytes"
	"github.com/yuidegm/Hotel-Rental-Managemnet-System/comment/com_Repository"
	"github.com/yuidegm/Hotel-Rental-Managemnet-System/comment/com_Service"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAdminFeedback(t *testing.T) {

	tmpl := template.Must(template.ParseGlob("C:/Users/getch/go/src/github.com/yuidegm/Hotel-Rental-Managemnet-System/ui/templates/admin_template/*"))

	EventsRepo := com_Repository.NewMockComRepo(nil)
	EventsServ := com_Service.NewGormComServiceImpl(EventsRepo)

	adminEventsHandler := NewAdminComHandler(tmpl, EventsServ)

	mux := http.NewServeMux()
	mux.HandleFunc("/admin/com", adminEventsHandler.AdminCom)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	url := ts.URL

	resp, err := tc.Get(url + "/admin/com")
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

	if !bytes.Contains(body, []byte("gech")) {
		t.Errorf("want body to contain %q", body)
	}

}
