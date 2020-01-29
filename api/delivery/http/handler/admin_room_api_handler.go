package handler

import (
	"encoding/json"
	"fmt"
	"github.com/aleale2121/Hotel-Final/api/entity"
	"github.com/aleale2121/Hotel-Final/api/rooms"
	"github.com/aleale2121/Hotel-Final/api/utils"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

// AdminRoomHandler handles category handler admin requests
type AdminRoomHandler struct {
	roomSrv rooms.RoomServices
}
// NewAdminRoomHandler initializes and returns new AdminCateogryHandler
func NewAdminRoomHandler( CS rooms.RoomServices) *AdminRoomHandler {
	return &AdminRoomHandler{ roomSrv: CS}
}

// AdminRooms handle requests on route /admin/rooms
func (arh *AdminRoomHandler) AdminRooms(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	room_list, errs := arh.roomSrv.Rooms()

	if errs!=nil{
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(room_list, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}

// AdminRoomNew hanlde requests on route /admin/rooms/new
func (arh *AdminRoomHandler) AdminRoomNew(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	body:=utils.BodyParser(r)
	var room1 entity.Room
	err :=json.Unmarshal(body,&room1)
	if err!=nil {
		utils.ToJson(w,err.Error(),http.StatusUnprocessableEntity)
		return
	}


	fmt.Println(room1)
	room2, errs := arh.roomSrv.StoreRoom(room1)

	if errs!=nil {
		fmt.Println("gorm error")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return

	}
	output, err := json.MarshalIndent(room2, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
	return
}

// AdminRoomsUpdate handle requests on /admin/rooms/update
func (arh *AdminRoomHandler) AdminRoomsUpdate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))
	fmt.Println("Server Updating ",id)
	if err != nil {
		fmt.Println("88")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	room1, errs := arh.roomSrv.Room(id)
    fmt.Println(room1)
	if len(errs)>0{
		fmt.Println("97")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	l := r.ContentLength

	body := make([]byte, l)

	r.Body.Read(body)

	json.Unmarshal(body, &room1)

	room1, errs = arh.roomSrv.UpdateRoom(room1)

	if len(errs)>0{
		fmt.Println("114")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(room1, "", "\t\t")

	if err != nil {
		fmt.Println("23")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
	return

}

// AdminRoomsDelete handle requests on route /admin/rooms/delete
func (arh *AdminRoomHandler) AdminRoomsDelete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	fmt.Println("Server Deleting ",id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, errs := arh.roomSrv.DeleteRoom(int(uint32(id)))

	if errs!=nil{
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return

}

func (arh *AdminRoomHandler) AdminRoomTypes(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	roomTypes, errs := arh.roomSrv.RoomTypes()

	if errs!=nil{
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(roomTypes, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}