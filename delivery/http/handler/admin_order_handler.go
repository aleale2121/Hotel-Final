package handler

import (
	"encoding/json"
	"fmt"
	"strconv"

	//pp "github.com/aleale2121/webProject2019/api/entity"
	"github.com/yuidegm/Hotel-Rental-Managemnet-System/entity"

	"html/template"
	"io/ioutil"
	"net/http"
)

// AdminRoomHandler handles category handler admin requests
type AdminOrderHandler struct {
	tmpl    *template.Template
}


// NewAdminRoomHandler initializes and returns new AdminCateogryHandler
func NewAdminOrderHandler(T *template.Template) *AdminRoomHandler {
	return &AdminRoomHandler{tmpl: T}
}

// AdminRooms handle requests on route /admin/orders
func (arh *AdminRoomHandler) AdminOrders(w http.ResponseWriter, r *http.Request) {
	client:=http.Client{}
	req,err:=http.NewRequest("GET","http://localhost:9090/user/reserve",nil)
	if err!=nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	response, err := client.Do(req)
	if http.StatusNotFound == response.StatusCode {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if http.StatusUnprocessableEntity == response.StatusCode {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return

	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var orderslist []entity.Order
	err = json.Unmarshal(responseData, &orderslist)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	arh.tmpl.ExecuteTemplate(w, "admin.order.layout", orderslist)

}
// AdminRooms handle requests on route /admin/rooms
func (arh *AdminRoomHandler) AdminOrderDelete(w http.ResponseWriter, r *http.Request) {
	idRaw := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idRaw)
	fmt.Println("Delete Order ",id)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	client:=http.Client{}
	mUrl:=fmt.Sprintf("http://localhost:9090/user/reserve/delete/%d",id)
	req,err:=http.NewRequest("DELETE",mUrl,nil)
	if err!=nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	response, err := client.Do(req)
	if err!=nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if http.StatusUnprocessableEntity == response.StatusCode {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return

	}
	http.Redirect(w,r,"/admin/orders",http.StatusSeeOther)
}
