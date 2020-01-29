package handler

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/aleale2121/Hotel-Final/api/entity"
	us "github.com/aleale2121/Hotel-Final/api/user"
	"github.com/aleale2121/Hotel-Final/api/utils"
	"net/http"
)

type SessionApiHandler struct{
	sessionService us.SessionService
}
func NewSessionApiHandler(sesServices us.SessionService) *SessionApiHandler {
	return &SessionApiHandler{sessionService: sesServices}
}
func (uph *SessionApiHandler) GetSessionByName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := ps.ByName("id")
	sessions, errs := uph.sessionService.Session(name)

	if errs!=nil{
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(sessions, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}
func (uph *SessionApiHandler) DeleteSession(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := ps.ByName("id")
	_, errs := uph.sessionService.DeleteSession(name)

	if errs!=nil{
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return

}
func (uph *SessionApiHandler) PostSession(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	session1 := &entity.Session{}

	err := json.Unmarshal(body, session1)

	if err!=nil {
		utils.ToJson(w,err.Error(),http.StatusUnprocessableEntity)
		return
	}

    fmt.Println("posting session")
	fmt.Println(session1)
	fmt.Println("posting session")
	_, errs := uph.sessionService.StoreSession(session1)

	if len(errs)>0 {
		fmt.Println("gorm error")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return

	}
	output, err := json.MarshalIndent(*session1, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
	return
}