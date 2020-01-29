package handler
import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/aleale2121/Hotel-Final/api/entity"
	us "github.com/aleale2121/Hotel-Final/api/user"
	"github.com/aleale2121/Hotel-Final/api/utils"
	"net/http"
	"strconv"
)
type RoleApiHandler struct{
	roleService us.RoleService
}
func NewRoleApiHandler(userServices us.RoleService) *RoleApiHandler {
	return &RoleApiHandler{roleService: userServices}
}
func (uph *RoleApiHandler) GetRoleByName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := ps.ByName("name")
	roles, errs := uph.roleService.RoleByName(name)

	if len(errs)>0{
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(roles, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}
func (uph *RoleApiHandler) GetRoleByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id,errs :=strconv.Atoi( ps.ByName("id"))

	if errs!=nil{
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	roles, err := uph.roleService.Role(uint(id))

	if len(err)>0{
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, errs := json.MarshalIndent(roles, "", "\t\t")

	if errs != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}
func (uph *RoleApiHandler) GetRoles(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	roles, errs := uph.roleService.Roles()

	if len(errs)>0{
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(roles, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}
func (uph *RoleApiHandler) PostRole(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	body:=utils.BodyParser(r)
	var role entity.Role
	err :=json.Unmarshal(body,&role)
	if err!=nil {
		utils.ToJson(w,err.Error(),http.StatusUnprocessableEntity)
		return
	}
	role1, errs := uph.roleService.StoreRole(&role)
	if len(errs)>0 {
		fmt.Println("gorm error")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return

	}
	output, err := json.MarshalIndent(role1, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
	return
}
func (uph *RoleApiHandler) PutRole(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	user1, errs := uph.roleService.Role(uint(id))

	if errs!=nil{
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	l := r.ContentLength

	body := make([]byte, l)

	r.Body.Read(body)

	json.Unmarshal(body, &user1)

	user1, errs = uph.roleService.UpdateRole(user1)

	if errs!=nil{
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(user1, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
	return
}
func (uph *RoleApiHandler) DeleteRole(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, errs := uph.roleService.DeleteRole(uint(id))

	if len(errs)>0{
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return
}
