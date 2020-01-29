
package handler

import (
	"github.com/aleale2121/Hotel-Final/entity"
	"github.com/aleale2121/Hotel-Final/form"
	"github.com/aleale2121/Hotel-Final/rtoken"
	"net/url"

	"github.com/aleale2121/Hotel-Final/comment"
	"html/template"

	"net/http"
)

// AdminNewsHandler handles category handler admin requests
type UserComHandler struct {
	tmpl        *template.Template
	newsService comment.CommentServices
	userHandler *UserHandler
	csrfSignKey    []byte
}
type commentData struct {
	IsLogged string
	FormInput form.Input
}

// NewAdminNewsHandler initializes and returns new AdminCateogryHandler
func NewUserComHandler(T *template.Template, NS comment.CommentServices,userHandler *UserHandler,csrfSignKey    []byte) *UserComHandler {
	return &UserComHandler{tmpl: T, newsService: NS,	csrfSignKey:csrfSignKey,userHandler:userHandler}
}


func (ach *UserComHandler) UserCom(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(ach.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	isLogged :="false"
	if ach.userHandler.loggedIn(r) {
		isLogged="true"
	}
	if r.Method == http.MethodPost {
		v := url.Values{}
		v.Add("fname", r.FormValue("fname"))
		v.Add("email", r.FormValue("email"))
		v.Add("comment", r.FormValue("comment"))
		ctg := entity.Comments{}
		//ctg.Publish=tutu
		ctg.Name = r.FormValue("fname")
		ctg.Email = r.FormValue("email")
		ctg.Comment = r.FormValue("comment")
		commentsForm := form.Input{Values: v, VErrors: form.ValidationErrors{}}
		commentsForm.Required("fname", "email","comment")
		commentsForm.MinLength("comment", 10)
		commentsForm.MatchesPattern("email",form.EmailRX)
		commentsForm.CSRF = token
		cData:=commentData{
			IsLogged:  isLogged,
			FormInput: commentsForm,
		}
		if !commentsForm.Valid() {
			ach.tmpl.ExecuteTemplate(w, "contact.layout", cData)
			return
		}
		_ ,err2:= ach.newsService.StoreCom(ctg)
		if  len(err2)>0 {
			commentsForm.VErrors.Add("generics","Sorry We Cannot Send Your Comment")
			ach.tmpl.ExecuteTemplate(w, "contact.layout", cData)
			return
		}
		commentsForm.VErrors.Add("success","Feedback sent successfully!! check your email for reply")

		ach.tmpl.ExecuteTemplate(w, "contact.layout", cData)

		return
	}
	if  r.Method == http.MethodGet {
		commentsForm := struct {
			Values  url.Values
			VErrors form.ValidationErrors
			CSRF    string
		}{
			Values:  nil,
			VErrors: nil,
			CSRF:    token,
		}
		cData:=commentData{
			IsLogged:  isLogged,
			FormInput: commentsForm,
		}
		ach.tmpl.ExecuteTemplate(w, "contact.layout", cData)

	}
}

