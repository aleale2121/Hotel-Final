package handler

import (
	"fmt"
	"github.com/aleale2121/Hotel-Final/entity"
	"github.com/aleale2121/Hotel-Final/form"
	"github.com/aleale2121/Hotel-Final/news"
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
type AdminNewsHandler struct {
	tmpl        *template.Template
	newsService news.NewsService
	csrfSignKey    []byte
}

// NewAdminNewsHandler initializes and returns new AdminCategoryHandler
func NewAdminNewsHandler(T *template.Template, NS news.NewsService,csrfSignKey    []byte) *AdminNewsHandler {
	return &AdminNewsHandler{tmpl: T, newsService: NS,csrfSignKey:csrfSignKey}
}

// AdminNews handle requests on route /admin/newss
func (ach *AdminNewsHandler) AdminNews(w http.ResponseWriter, r *http.Request) {
	news, err := ach.newsService.News()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	ach.tmpl.ExecuteTemplate(w, "admin.newss.layout", news)
}

// AdminNewsNew hanlde requests on route /admin/new
func (ach *AdminNewsHandler) AdminNewsNew(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(ach.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if r.Method == http.MethodPost {
		v := url.Values{}
		v.Add("name", r.FormValue("name"))
		v.Add("description", r.FormValue("description"))
		newsForm := form.Input{Values: v, VErrors: form.ValidationErrors{}}
		newsForm.Required("name", "description")
		newsForm.MinLength("name", 10)
		newsForm.MinLength("description", 15)
		newsForm.CSRF = token
		if !newsForm.Valid() {
			ach.tmpl.ExecuteTemplate(w, "admin.newss.new.layout", newsForm)
			return
		}
		ctg := entity.News{}
		ctg.Header = r.FormValue("name")
		ctg.Description = r.FormValue("description")
		mf, fh, err := r.FormFile("catimg")
		if err != nil {
			fmt.Println("67")
			newsForm.VErrors.Add("catimg", "Cannot store your image")
			ach.tmpl.ExecuteTemplate(w, "admin.newss.new.layout", newsForm)
			return
		}
		defer mf.Close()

		ctg.Image = fh.Filename

		WriteFile(&mf, fh.Filename)

		_ ,err2:= ach.newsService.StoreNews(ctg)

		if  len(err2)>0 {
			fmt.Println("81")
			newsForm.VErrors.Add("catimg", "Cannot store your image")
			ach.tmpl.ExecuteTemplate(w, "admin.newss.new.layout", newsForm)
			return
		}
		newsForm.VErrors.Add("success", "Successfully Added")
		ach.tmpl.ExecuteTemplate(w, "admin.newss.new.layout", newsForm)

	}
	if r.Method==http.MethodGet {
		newsForm := struct {
			Values  url.Values
			VErrors form.ValidationErrors
			CSRF    string
		}{
			Values:  nil,
			VErrors: nil,
			CSRF:    token,
		}
		ach.tmpl.ExecuteTemplate(w, "admin.newss.new.layout", newsForm)

	}
}

// AdminNewsUpdate handle requests on /admin/categories/update
func (ach *AdminNewsHandler) AdminNewsUpdate(w http.ResponseWriter, r *http.Request) {
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

		cat, err2 := ach.newsService.NewsById(id)
		if len(err2)>0{
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		newsUpdateForm := struct {
			FormInput form.Input
			New  entity.News
			CSRF    string
		}{

			New: *cat,
			CSRF:    token,
		}

		ach.tmpl.ExecuteTemplate(w, "admin.newss.update.layout", newsUpdateForm)

	} else if r.Method == http.MethodPost {
		v := url.Values{}
		v.Add("name", r.FormValue("name"))
		v.Add("description", r.FormValue("description"))
		newsForm := form.Input{Values: v, VErrors: form.ValidationErrors{}}
		newsForm.Required("name", "description")
		newsForm.MinLength("name", 10)
		newsForm.MinLength("description", 15)
		newsForm.CSRF = token


		new := entity.News{}
		new.Id, err = strconv.Atoi(r.FormValue("id"))
		if err!=nil{
			newsForm.VErrors.Add("generic","Cannot Update The Room")
		}
		new.Header = r.FormValue("name")
		new.Description = r.FormValue("description")
		new.Image = r.FormValue("image")
		newsUpdateForm := struct {
			FormInput form.Input
			New entity.News
			CSRF    string
		}{
			FormInput:newsForm,
			New:new,
			CSRF:    token,
		}
		if !newsForm.Valid() {
			ach.tmpl.ExecuteTemplate(w, "admin.newss.update.layout", newsUpdateForm)
			return
		}
		mf, _, err := r.FormFile("catimg")
		if err!=nil{
			newsForm.VErrors.Add("catimg","Cannot Process This Image")
			newsUpdateForm = struct {
				FormInput form.Input
				New entity.News
				CSRF    string
			}{
				FormInput:newsForm,
				New:new,
				CSRF:    token,
			}
			ach.tmpl.ExecuteTemplate(w, "admin.newss.update.layout", newsUpdateForm)
			return
		}

		defer mf.Close()

		WriteFile(&mf, new.Image)

		_,err2 := ach.newsService.UpdateNews(new)

		if len(err2)>0 {
			newsForm.VErrors.Add("generic","Cannot Process This Image")
			newsUpdateForm = struct {
				FormInput form.Input
				New entity.News
				CSRF    string
			}{
				FormInput:newsForm,
				New:new,
				CSRF:    token,
			}
			ach.tmpl.ExecuteTemplate(w, "admin.newss.update.layout", newsUpdateForm)
			return
		}
		newsForm.VErrors.Add("success","Successfully Updated")
		newsUpdateForm = struct {
			FormInput form.Input
			New entity.News
			CSRF    string
		}{
			FormInput:newsForm,
			New:new,
			CSRF:    token,
		}
		ach.tmpl.ExecuteTemplate(w, "admin.newss.update.layout", newsUpdateForm)

	} else {
		http.Redirect(w, r, "/admin/newss", http.StatusSeeOther)
	}

}


// AdminNewsDelete handle requests on route /admin/categories/delete
func (ach *AdminNewsHandler) AdminNewsDelete(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")

		id, err := strconv.Atoi(idRaw)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		_,err2 := ach.newsService.DeleteNews(id)

		if len(err2)>0 {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

	}

	http.Redirect(w, r, "/admin/newss", http.StatusSeeOther)
}

func WriteFile(mf *multipart.File, Frame string) {

	wd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	path := filepath.Join(wd, "ui", "assets", "img", Frame)
	image, err := os.Create(path)

	if err != nil {
		panic(err)
	}
	defer image.Close()
	io.Copy(image, *mf)
}

