package handler

import (
	"fmt"
	"github.com/aleale2121/Hotel-Final/entity"
	"github.com/aleale2121/Hotel-Final/events"
	"github.com/aleale2121/Hotel-Final/form"
	"github.com/aleale2121/Hotel-Final/rtoken"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
)

// AdminNewsHandler handles category handler admin requests
type AdminEventsHandler struct {
	tmpl          *template.Template
	eventsService events.EventsService
	csrfSignKey   []byte
}

// NewAdminNewsHandler initializes and returns new AdminCateogryHandler
func NewAdminEventsHandler(T *template.Template, NS events.EventsService,csrfSignKey    []byte) *AdminEventsHandler {
	return &AdminEventsHandler{tmpl: T, eventsService: NS,csrfSignKey:csrfSignKey}
}

// AdminNews handle requests on route /admin/newss
func (ach *AdminEventsHandler) AdminEvents(w http.ResponseWriter, r *http.Request) {
	news, err := ach.eventsService.Events()

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	ach.tmpl.ExecuteTemplate(w, "admin.events.layout", news)
}

// AdminNewsNew hanlde requests on route /admin/new
func (ach *AdminEventsHandler) AdminEventsNew(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(ach.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if r.Method == http.MethodPost {
		v := url.Values{}
		v.Add("name", r.FormValue("name"))
		v.Add("description", r.FormValue("description"))
		eventsForm := form.Input{Values: v, VErrors: form.ValidationErrors{}}
		eventsForm.Required("name", "description")
		eventsForm.MinLength("name", 10)
		eventsForm.MinLength("description", 15)
		eventsForm.CSRF = token
		if !eventsForm.Valid() {
			ach.tmpl.ExecuteTemplate(w, "admin.events.new.layout", eventsForm)
			return
		}
		event := entity.Events{}
		event.Header = r.FormValue("name")
		event.Description = r.FormValue("description")
		mf, fh, err := r.FormFile("catimg")
		if err != nil {
			fmt.Println("67")
			eventsForm.VErrors.Add("catimg", "Cannot store your image")
			ach.tmpl.ExecuteTemplate(w, "admin.events.new.layout", eventsForm)
			return
		}
		defer mf.Close()

		event.Image = fh.Filename

		WriteFile(&mf, fh.Filename)

		_ ,err2:= ach.eventsService.StoreEvents(event)

		if  len(err2)>0 {
			fmt.Println("81")
			eventsForm.VErrors.Add("catimg", "Cannot store your image")
			ach.tmpl.ExecuteTemplate(w, "admin.events.new.layout", eventsForm)
			return
		}
		eventsForm.VErrors.Add("success", "Successfully Added")
		ach.tmpl.ExecuteTemplate(w, "admin.events.new.layout", eventsForm)

	}
	if r.Method==http.MethodGet {
		eventForm := struct {
			Values  url.Values
			VErrors form.ValidationErrors
			CSRF    string
		}{
			Values:  nil,
			VErrors: nil,
			CSRF:    token,
		}
		ach.tmpl.ExecuteTemplate(w, "admin.events.new.layout", eventForm)

	}

}

// AdminNewsUpdate handle requests on /admin/events/update
func (ach *AdminEventsHandler) AdminEventsUpdate(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(ach.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		cat, err2 := ach.eventsService.EventsById(id)
		if len(err2)>0{
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		newsUpdateForm := struct {
			FormInput form.Input
			Event     entity.Events
			CSRF      string
		}{

			Event: *cat,
			CSRF:  token,
		}

		ach.tmpl.ExecuteTemplate(w, "admin.events.update.layout", newsUpdateForm)

	} else if r.Method == http.MethodPost {
		v := url.Values{}
		v.Add("name", r.FormValue("name"))
		v.Add("description", r.FormValue("description"))
		eventsForm := form.Input{Values: v, VErrors: form.ValidationErrors{}}
		eventsForm.Required("name", "description")
		eventsForm.MinLength("name", 10)
		eventsForm.MinLength("description", 15)
		eventsForm.CSRF = token


		event := entity.Events{}
		event.Id, err = strconv.Atoi(r.FormValue("id"))
		if err!=nil{
			fmt.Println("152")
			eventsForm.VErrors.Add("generic","Cannot Update The Event")
		}
		event.Header = r.FormValue("name")
		event.Description = r.FormValue("description")
		event.Image = r.FormValue("image")

		newsUpdateForm := struct {
			FormInput form.Input
			Events    entity.Events
			CSRF      string
		}{
			FormInput: eventsForm,
			Events:    event,
			CSRF:      token,
		}
		if !eventsForm.Valid() {

			ach.tmpl.ExecuteTemplate(w, "admin.events.update.content", newsUpdateForm)
			return
		}
		mf, _, err := r.FormFile("catimg")
		if err!=nil{
			fmt.Println("174")
			eventsForm.VErrors.Add("catimg","Cannot Process This Image")
			newsUpdateForm = struct {
				FormInput form.Input
				Events entity.Events
				CSRF    string
			}{
				FormInput: eventsForm,
				Events:       event,
				CSRF:      token,
			}
			ach.tmpl.ExecuteTemplate(w, "admin.events.update.layout", newsUpdateForm)
			return
		}

		defer mf.Close()

		WriteFile(&mf, event.Image)

		_,err2 := ach.eventsService.UpdateEvents(event)

		if len(err2)>0 {
			fmt.Println("196")
			eventsForm.VErrors.Add("generic","Cannot Process This Image")
			eventsForm.VErrors.Add("catimg","Cannot Process This Image")
			newsUpdateForm = struct {
				FormInput form.Input
				Events entity.Events
				CSRF    string
			}{
				FormInput: eventsForm,
				Events:       event,
				CSRF:      token,
			}
			ach.tmpl.ExecuteTemplate(w, "admin.events.update.layout", newsUpdateForm)
			return
		}
		eventsForm.VErrors.Add("success","Successfully Updated")
		newsUpdateForm = struct {
			FormInput form.Input
			Events entity.Events
			CSRF    string
		}{
			FormInput: eventsForm,
			Events:       event,
			CSRF:      token,
		}
		ach.tmpl.ExecuteTemplate(w, "admin.events.update.layout", newsUpdateForm)

	} else {
		http.Redirect(w, r, "/admin/events", http.StatusSeeOther)
	}

}


// AdminNewsDelete handle requests on route /events/categories/delete
func (ach *AdminEventsHandler) AdminEventsDelete(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")

		id, err := strconv.Atoi(idRaw)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		_,err2 := ach.eventsService.DeleteEvents(id)

		if len(err2)>0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

	}

	http.Redirect(w, r, "/admin/events", http.StatusSeeOther)
}

func WriteEventsFile(mf *multipart.File, fname string) {

	wd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	path := filepath.Join(wd, "ui", "assets", "img", fname)

	image, err := os.Create(path)

	if err != nil {
		panic(err)
	}
	defer image.Close()
	io.Copy(image, *mf)
}

