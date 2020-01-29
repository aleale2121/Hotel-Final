package handler

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/yuidegm/Hotel-Rental-Managemnet-System/api/entity"
	"github.com/yuidegm/Hotel-Rental-Managemnet-System/api/order"
	"github.com/yuidegm/Hotel-Rental-Managemnet-System/api/utils"
	"net/http"
	"strconv"
)

type RoomReserveHandler struct{

	reserveService order.OrderService
}
func NewRoomReserveHandler(reserveServices order.OrderService) *RoomReserveHandler {
	return &RoomReserveHandler{reserveService: reserveServices}
}
func (rrh *RoomReserveHandler) GetReservations(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	reserved, errs := rrh.reserveService.Orders()

	if errs!=nil{
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(reserved, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}
func (rrh *RoomReserveHandler) GetReservation(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	reserved, errs := rrh.reserveService.CustomerOrders(int32(id))

	if errs!=nil{
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(reserved, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}
func (rrh *RoomReserveHandler) GetRoomReservation(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	reserved, errs := rrh.reserveService.RoomOrder(uint32(id))

	if errs!=nil{
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(reserved, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}
func (rrh *RoomReserveHandler) PostOrder(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	body:=utils.BodyParser(r)
	var orderedRoom entity.Order
	err :=json.Unmarshal(body,&orderedRoom)
	if err!=nil {
		utils.ToJson(w,err.Error(),http.StatusUnprocessableEntity)
		return
	}


    fmt.Println("post order api",orderedRoom)
	_, errs := rrh.reserveService.StoreOrder(&orderedRoom)

	if errs!=nil {
		fmt.Println("gorm error")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	utils.ToJson(w,"post post successful",http.StatusCreated)
	return
}
func (rrh *RoomReserveHandler) PutReservationOrder(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	orders, errs := rrh.reserveService.Order(uint32(id))

	if errs!=nil{
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	l := r.ContentLength

	body := make([]byte, l)

	r.Body.Read(body)

	json.Unmarshal(body, &orders)

	orders, errs = rrh.reserveService.UpdateOrder(orders)

	if errs!=nil{
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(orders, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
	return
}
func (rrh *RoomReserveHandler) DeleteReservation(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, errs := rrh.reserveService.DeleteOrder(uint32(id))

	if errs!=nil{
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	//response, errr := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")


	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return
}