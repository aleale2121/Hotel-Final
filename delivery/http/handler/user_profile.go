


package handler

import (
    "encoding/json"
    "fmt"
    "github.com/yuidegm/Hotel-Rental-Managemnet-System/entity"
    "html/template"
    "io/ioutil"
    "net/http"
)

type profileHandler struct {
    tmpl    *template.Template
    userHandler *UserHandler
}
func NewUserProfileHandler(T *template.Template,userHandler *UserHandler) *profileHandler {
    return &profileHandler{tmpl: T,userHandler:userHandler}
}
type UserData struct {
    UserInfo entity.User
    UserOrderInfo []entity.Order
}
func (ph *profileHandler) UserInfo(w http.ResponseWriter, r *http.Request) {
    id:=ph.userHandler.loggedInUser.Id
    cliet:=&http.Client{}
    url:=fmt.Sprintf("http://localhost:9090/user/reserves/%d",id)
    req, err := http.NewRequest(http.MethodGet,url, nil)
    response,err:=cliet.Do(req)
    if http.StatusNotFound == response.StatusCode {
        http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
        return
    }
    if http.StatusUnprocessableEntity == response.StatusCode{
        http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
        return

    }
    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
        return
    }

    var responseObject []entity.Order
    err = json.Unmarshal(responseData, &responseObject)
    if err != nil {
        http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
        return
    }
    fmt.Println(responseObject)
    info:=UserData{UserInfo:*ph.userHandler.loggedInUser,UserOrderInfo:responseObject}
    _ = ph.tmpl.ExecuteTemplate(w, "profile.layout", info)
}