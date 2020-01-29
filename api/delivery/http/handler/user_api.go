package handler
import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/aleale2121/Hotel-Final/api/entity"
	user "github.com/aleale2121/Hotel-Final/api/user"
	"github.com/aleale2121/Hotel-Final/api/utils"
	"net/http"
	"strconv"
)

type UserApiHandler struct{
	userService user.UserService
}

func NewUserApiHandler(userServices user.UserService) *UserApiHandler {
	return &UserApiHandler{userService: userServices}
}
func (uph *UserApiHandler) GetUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	users, errs := uph.userService.User(uint32(id))

	if errs!=nil{
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(users, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}
func (uph *UserApiHandler) GetUserRoles(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	body:=utils.BodyParser(r)
	var user1 entity.User
	err :=json.Unmarshal(body,&user1)
	if err!=nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return

	}
	userRoles, errs := uph.userService.UserRoles(&user1)
	if len(errs)>0{

		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(userRoles, "", "\t\t")

	if err != nil {

		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}
func (uph *UserApiHandler) IsEmailExists(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	email:= ps.ByName("email")

	ok := uph.userService.EmailExists(email)
	output, err := json.MarshalIndent(ok, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}
func (uph *UserApiHandler) IsPhoneExists(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	phone:= ps.ByName("phone")

	ok := uph.userService.PhoneExists(phone)
	output, err := json.MarshalIndent(ok, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}
func (uph *UserApiHandler) GetUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	users, errs := uph.userService.Users()

	if errs!=nil{
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(users, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}
func (uph *UserApiHandler) GetUserByUsernameAndPassword(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {


	body:=utils.BodyParser(r)
	var user1 entity.User
	err :=json.Unmarshal(body,&user1)
	if err!=nil {
		utils.ToJson(w,err.Error(),http.StatusUnprocessableEntity)
		return
	}
	user2, errs := uph.userService.UserByUserName(user1)

	if errs!=nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return

	}
	output, err := json.MarshalIndent(*user2, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	fmt.Println("user api jso",string(output))
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
	return

}
func (uph *UserApiHandler) PostUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	body:=utils.BodyParser(r)
	var user1 entity.User
	err :=json.Unmarshal(body,&user1)
	if err!=nil {
		utils.ToJson(w,err.Error(),http.StatusUnprocessableEntity)
		return
	}
	user2, errs := uph.userService.StoreUser(&user1)

	if errs!=nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return

	}
	output, err := json.MarshalIndent(user2, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
	return
}
func (uph *UserApiHandler) PutUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
   fmt.Println("user")
	id, err := strconv.Atoi(ps.ByName("id"))
	fmt.Println(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	fmt.Println(id)
	user1, errs := uph.userService.User(uint32(id))

	if errs!=nil{
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	fmt.Println(user1)
	l := r.ContentLength

	body := make([]byte, l)

	_, _ = r.Body.Read(body)

	_ = json.Unmarshal(body, &user1)
	fmt.Println(user1)
	user1, errs = uph.userService.UpdateUser(user1)

	if errs!=nil{
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	fmt.Println(user1)
	output, err := json.MarshalIndent(user1, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	fmt.Println(string(output))
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
	return
}
func (uph *UserApiHandler) DeleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, errs := uph.userService.DeleteUser(uint32(id))

	if errs!=nil{
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return
}
