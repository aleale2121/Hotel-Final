package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/yuidegm/Hotel-Rental-Managemnet-System/entity"
	"github.com/yuidegm/Hotel-Rental-Managemnet-System/form"
	"github.com/yuidegm/Hotel-Rental-Managemnet-System/rtoken"
	"html/template"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
)

// AdminRoomHandler handles category handler admin requests
type AdminRoomHandler struct {
	tmpl    *template.Template
	csrfSignKey []byte
}
type RoomAndRoomType struct {
	AllRooms []entity.Room
	AllRoomCategory []entity.Type
	FormInput form.Input
	ActionMode string
	RoomOnAction entity.Room
}
var roomcar RoomAndRoomType
// NewAdminRoomHandler initializes and returns new AdminCateogryHandler
func NewAdminRoomHandler(T *template.Template, csKey []byte) *AdminRoomHandler {
	return &AdminRoomHandler{tmpl: T ,csrfSignKey: csKey}
}
// AdminRooms handle requests on route /admin/rooms
func (arh *AdminRoomHandler) AdminRooms(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(arh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if r.Method==http.MethodGet{
		response, err := http.Get("http://localhost:9090/room/rooms")
		if response.StatusCode == http.StatusNotFound {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		}
		if http.StatusUnprocessableEntity == response.StatusCode {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		}
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		}
		var roomslist []entity.Room
		err = json.Unmarshal(responseData, &roomslist)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		}
		//fmt.Println(roomslist)
		//for room types
		client:=&http.Client{}
		req, err := http.NewRequest("GET","http://localhost:9090/room/types",nil)
		response2, err:=client.Do(req)
		if http.StatusNotFound == response2.StatusCode {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		}
		if http.StatusUnprocessableEntity == response2.StatusCode {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		}
		responseData2, err := ioutil.ReadAll(response2.Body)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		}
		var roomTypes []entity.Type
		//var roomTypes2 []entity.Type
		err = json.Unmarshal(responseData2, &roomTypes)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		}
		//fmt.Println(roomTypes)
		RoomForm := form.Input{Values: nil, VErrors: nil,CSRF:token}
		roomcar=RoomAndRoomType{AllRooms: roomslist,AllRoomCategory: roomTypes,FormInput:RoomForm}
		//fmt.Println(roomcar)
		arh.tmpl.ExecuteTemplate(w, "room_index2.layout", roomcar)
		return
	}
	if r.Method==http.MethodPost{
		v := url.Values{}
		v.Add("roomNum", r.FormValue("roomNum"))
		v.Add("price", r.FormValue("price"))
		v.Add("description", r.FormValue("description"))
		v.Add("catimg", r.FormValue("catimg"))
		newRoomForm := form.Input{Values: v, VErrors: form.ValidationErrors{}}
		newRoomForm.Required("roomNum", "price","description")
		newRoomForm.MinLength("description", 10)
		newRoomForm.MatchesPattern("price", form.PriceRX)
		newRoomForm.CSRF = token
		rooms,types:=arh.GetRoomAndRoomCategories()


		room := entity.Room{}
		room.RoomNumber, _ = strconv.Atoi( r.FormValue("roomNum"))
		id, _ := strconv.Atoi( r.FormValue("roomCategory"))
		room.TypeId= uint32(id)
		room.Price,_=strconv.ParseFloat( r.FormValue("price"),64)
		room.Description = r.FormValue("description")
		mf, fh, err := r.FormFile("catimg")
		if err != nil {
			rooms,types=arh.GetRoomAndRoomCategories()
			//Println("131")
			newRoomForm.VErrors.Add("catimg", "Cannot Add  Room")
			roomcar=RoomAndRoomType{AllRooms: rooms,AllRoomCategory: types,FormInput:newRoomForm,ActionMode:"postRoom",RoomOnAction:room}
			arh.tmpl.ExecuteTemplate(w, "room_index2.layout", roomcar)
			return
		}


		defer mf.Close()
		room.Image = fh.Filename
		writeFile(&mf, fh.Filename)
		if !newRoomForm.Valid() {
			//fmt.Println(roomcar)
			roomcar=RoomAndRoomType{AllRooms: rooms,AllRoomCategory: types,FormInput:newRoomForm,ActionMode:"postRoom",RoomOnAction:room}
			arh.tmpl.ExecuteTemplate(w, "room_index2.layout", roomcar)
			//http.Redirect(w, r, "/admin/rooms", 300)
			return
		}
		//_, _ = arh.roomSrv.StoreRoom(room)

		//fmt.Println(err2)
		output, err := json.MarshalIndent(room, "", "\t\t")
		if err != nil {
			//fmt.Println("145")
			rooms,types=arh.GetRoomAndRoomCategories()
			newRoomForm.VErrors.Add("generic", "Cannot Add Event Room")
			roomcar=RoomAndRoomType{AllRooms: rooms,AllRoomCategory: types,FormInput:newRoomForm,ActionMode:"postRoom",RoomOnAction:room}
			arh.tmpl.ExecuteTemplate(w, "room_index2.layout", roomcar)
			return
		}

		response, err := http.Post("http://localhost:9090/room/new","application/json", bytes.NewBuffer(output))
		if http.StatusNotFound == response.StatusCode {
			//fmt.Println("154")
			rooms,types=arh.GetRoomAndRoomCategories()
			roomcar=RoomAndRoomType{AllRooms: rooms,AllRoomCategory: types,FormInput:newRoomForm,RoomOnAction:room}
			arh.tmpl.ExecuteTemplate(w, "room_index2.layout", roomcar)
			return
		}
		if http.StatusUnprocessableEntity == response.StatusCode {
			//fmt.Println("159")
			rooms,types=arh.GetRoomAndRoomCategories()
			newRoomForm.VErrors.Add("generic", "Cannot Add Event Room")
			roomcar=RoomAndRoomType{AllRooms: rooms,AllRoomCategory: types,FormInput:newRoomForm,ActionMode:"postRoom",RoomOnAction:room}
			arh.tmpl.ExecuteTemplate(w, "room_index2.layout", roomcar)
			return
		}
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			//fmt.Println("166")
			rooms,types=arh.GetRoomAndRoomCategories()
			newRoomForm.VErrors.Add("generic", "Cannot Add Event Room")
			roomcar=RoomAndRoomType{AllRooms: rooms,AllRoomCategory: types,FormInput:newRoomForm,ActionMode:"postRoom",RoomOnAction:room}
			arh.tmpl.ExecuteTemplate(w, "room_index2.layout", roomcar)
			return
		}

		var responseObject entity.Room
		err = json.Unmarshal(responseData, &responseObject)
		if err != nil {
			//fmt.Println("175")
			rooms,types=arh.GetRoomAndRoomCategories()
			roomcar=RoomAndRoomType{AllRooms: rooms,AllRoomCategory: types,FormInput:newRoomForm,ActionMode:"postRoom"}
			arh.tmpl.ExecuteTemplate(w, "room_index2.layout", roomcar)
			return
		}
		rooms,types=arh.GetRoomAndRoomCategories()
		roomcar=RoomAndRoomType{AllRooms: rooms,AllRoomCategory: types,FormInput:newRoomForm,ActionMode:"postRoom"}
		newRoomForm.VErrors.Add("success", "Successfully Added")
		arh.tmpl.ExecuteTemplate(w, "room_index2.layout", roomcar)
		//http.Redirect(w, r, "/admin/rooms", 300)
	}



}

// AdminRoomNew handle requests on route /admin/rooms/new
func (arh *AdminRoomHandler) AdminRoomNew(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(arh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if r.Method == http.MethodPost {
		v := url.Values{}
		v.Add("roomNum", r.FormValue("roomNum"))
		v.Add("price", r.FormValue("price"))
		v.Add("description", r.FormValue("description"))
		v.Add("catimg", r.FormValue("catimg"))
		newRoomForm := form.Input{Values: v, VErrors: form.ValidationErrors{}}
		newRoomForm.Required("roomNum", "price","description")
		newRoomForm.MinLength("description", 10)
		newRoomForm.MatchesPattern("price", form.PriceRX)
		newRoomForm.CSRF = token
		rooms,types:=arh.GetRoomAndRoomCategories()
		if !newRoomForm.Valid() {
			//fmt.Println(roomcar)
			roomcar=RoomAndRoomType{AllRooms: rooms,AllRoomCategory: types,FormInput:newRoomForm}
			arh.tmpl.ExecuteTemplate(w, "room_index2.layout", roomcar)
			//http.Redirect(w, r, "/admin/rooms", 300)
			return
		}


		room := entity.Room{}
		room.RoomNumber, _ = strconv.Atoi( r.FormValue("roomNum"))
		id, _ := strconv.Atoi( r.FormValue("roomCategory"))
		room.TypeId= uint32(id)
		room.Price,_=strconv.ParseFloat( r.FormValue("price"),64)
		room.Description = r.FormValue("description")
		mf, fh, err := r.FormFile("catimg")

		if err != nil {
			//Println("131")
			newRoomForm.VErrors.Add("catimg", "Cannot Add  Room")
			roomcar=RoomAndRoomType{AllRooms: rooms,AllRoomCategory: types,FormInput:newRoomForm}
			arh.tmpl.ExecuteTemplate(w, "room_index2.layout", roomcar)
			return
		}
		defer mf.Close()
		room.Image = fh.Filename
		writeFile(&mf, fh.Filename)
		//_, _ = arh.roomSrv.StoreRoom(room)

		//fmt.Println(err2)
		output, err := json.MarshalIndent(room, "", "\t\t")
		if err != nil {
			//fmt.Println("145")
			newRoomForm.VErrors.Add("generic", "Cannot Add Event Room")
			roomcar=RoomAndRoomType{AllRooms: rooms,AllRoomCategory: types,FormInput:newRoomForm}
			arh.tmpl.ExecuteTemplate(w, "room_index2.layout", roomcar)
			return
		}

		response, err := http.Post("http://localhost:9090/room/new","application/json", bytes.NewBuffer(output))
		if http.StatusNotFound == response.StatusCode {
			//fmt.Println("154")
			roomcar=RoomAndRoomType{AllRooms: rooms,AllRoomCategory: types,FormInput:newRoomForm}
			arh.tmpl.ExecuteTemplate(w, "room_index2.layout", roomcar)
			return
		}
		if http.StatusUnprocessableEntity == response.StatusCode {
			//fmt.Println("159")
			roomcar=RoomAndRoomType{AllRooms: rooms,AllRoomCategory: types,FormInput:newRoomForm}
			arh.tmpl.ExecuteTemplate(w, "room_index2.layout", roomcar)
			return
		}
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			//fmt.Println("166")
			newRoomForm.VErrors.Add("generic", "Cannot Add Event Room")
			roomcar=RoomAndRoomType{AllRooms: rooms,AllRoomCategory: types,FormInput:newRoomForm}
			arh.tmpl.ExecuteTemplate(w, "room_index2.layout", roomcar)
			return
		}

		var responseObject entity.Room
		err = json.Unmarshal(responseData, &responseObject)
		if err != nil {
			//fmt.Println("175")
			roomcar=RoomAndRoomType{AllRooms: rooms,AllRoomCategory: types,FormInput:newRoomForm}
			arh.tmpl.ExecuteTemplate(w, "room_index2.layout", roomcar)
			return
		}
		newRoomForm.VErrors.Add("success", "Successfully Added")
		arh.tmpl.ExecuteTemplate(w, "room_index2.layout", roomcar)
		return

	}
}

// AdminRoomsUpdate handle requests on /admin/rooms/update
func (arh *AdminRoomHandler) AdminRoomsUpdate(w http.ResponseWriter, r *http.Request) {
	token, err := rtoken.CSRFToken(arh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
    if r.Method == http.MethodPost {
		v := url.Values{}
		v.Add("roomNumUpdated", r.FormValue("roomNumUpdated"))
		v.Add("priceUpdated", r.FormValue("priceUpdated"))
		v.Add("descriptionUpdated", r.FormValue("descriptionUpdated"))
		v.Add("catimgUpdated", r.FormValue("catimgUpdated"))
		newRoomForm := form.Input{Values: v, VErrors: form.ValidationErrors{}}
		newRoomForm.Required("roomNumUpdated", "priceUpdated","descriptionUpdated","descriptionUpdated")
		newRoomForm.MinLength("descriptionUpdated", 10)
		newRoomForm.MatchesPattern("priceUpdated", form.PriceRX)
		newRoomForm.CSRF = token
		rooms,types:=arh.GetRoomAndRoomCategories()
		if !newRoomForm.Valid() {
			roomcar=RoomAndRoomType{AllRooms: rooms,AllRoomCategory: types,FormInput:newRoomForm,ActionMode:"updateRoom"}
			arh.tmpl.ExecuteTemplate(w, "room_index2.layout", roomcar)
			return
		}
		room := entity.Room{}
		vx:=r.FormValue("id")
		idd,_:=strconv.Atoi(vx)
		room.Id= uint32(idd)
		room.RoomNumber, _ = strconv.Atoi( r.FormValue("roomNumUpdated"))
		rtid, _ := strconv.Atoi( r.FormValue("roomCategoryUpdated"))
		room.TypeId= uint32(rtid)
		room.Price,_=strconv.ParseFloat( r.FormValue("priceUpdated"),64)
		room.Description = r.FormValue("descriptionUpdated")
		mf, fh, err := r.FormFile("catimgUpdated")
		if err != nil {
			newRoomForm.VErrors.Add("generics","Sorry Internal Server Error Has Occurred")
			roomcar=RoomAndRoomType{AllRooms: rooms,AllRoomCategory: types,FormInput:newRoomForm,ActionMode:"updateRoom"}
			arh.tmpl.ExecuteTemplate(w, "room_index2.layout", roomcar)
			return
		}
		defer mf.Close()
		room.Image = fh.Filename
		writeFile(&mf, fh.Filename)

		output, err := json.MarshalIndent(room, "", "\t\t")
		if err != nil {
			newRoomForm.VErrors.Add("generics2","Sorry Internal Server Error Has Occurred")
			roomcar=RoomAndRoomType{AllRooms: rooms,AllRoomCategory: types,FormInput:newRoomForm,ActionMode:"updateRoom"}
			arh.tmpl.ExecuteTemplate(w, "room_index2.layout", roomcar)
			return
		}
		fmt.Println("341")
        fmt.Println(room)
        fmt.Println(room.Id)
		client :=&http.Client{}
        urlRoomUpdate :=fmt.Sprintf("http://localhost:9090/room/update/%d",room.Id)
		req, err := http.NewRequest(http.MethodPut, urlRoomUpdate, bytes.NewBuffer(output))
		response,err:= client.Do(req)
		fmt.Println(response)
		if response.StatusCode==http.StatusNotFound  {
			newRoomForm.VErrors.Add("roomNumUpdated","Sorry  Room Number Already not Exists ")
			roomcar=RoomAndRoomType{AllRooms: rooms,AllRoomCategory: types,FormInput:newRoomForm,ActionMode:"updateRoom"}
			arh.tmpl.ExecuteTemplate(w, "room_index2.layout", roomcar)
			return
		}
		if http.StatusUnprocessableEntity == response.StatusCode{
			newRoomForm.VErrors.Add("generics2","Sorry Internal Server Error Has Occurred")
			roomcar=RoomAndRoomType{AllRooms: rooms,AllRoomCategory: types,FormInput:newRoomForm,ActionMode:"updateRoom"}
			arh.tmpl.ExecuteTemplate(w, "room_index2.layout", roomcar)
			return

		}
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			newRoomForm.VErrors.Add("generics2","Sorry Internal Server Error Has Occurred")
			roomcar=RoomAndRoomType{AllRooms: rooms,AllRoomCategory: types,FormInput:newRoomForm,ActionMode:"updateRoom"}
			arh.tmpl.ExecuteTemplate(w, "room_index2.layout", roomcar)
			return
		}

		var responseObject entity.Room
		err = json.Unmarshal(responseData, &responseObject)
		if err != nil {
			newRoomForm.VErrors.Add("generics2","Sorry Internal Server Error Has Occurred")
			roomcar=RoomAndRoomType{AllRooms: rooms,AllRoomCategory: types,FormInput:newRoomForm,ActionMode:"updateRoom"}
			arh.tmpl.ExecuteTemplate(w, "room_index2.layout", roomcar)
			return
		}
		newRoomForm.VErrors.Add("success2","Room Successfully Updated")
		rooms,types=arh.GetRoomAndRoomCategories()
		roomcar=RoomAndRoomType{AllRooms: rooms,AllRoomCategory: types,FormInput:newRoomForm,ActionMode:"updateRoom"}
		arh.tmpl.ExecuteTemplate(w, "room_index2.layout", roomcar)
		return

	} else {
		http.Redirect(w, r, "/admin/rooms", http.StatusSeeOther)
	}

}

// AdminRoomsDelete handle requests on route /admin/rooms/delete
func (arh *AdminRoomHandler) AdminRoomsDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")

		id, err := strconv.Atoi(idRaw)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		client :=&http.Client{}
		url:=fmt.Sprintf("http://localhost:9090/room/delete/%d",id)
		req, err := http.NewRequest(http.MethodDelete,url, nil)
		response,err:= client.Do(req)
		fmt.Println(response)
		if http.StatusNotFound == response.StatusCode {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		if http.StatusUnprocessableEntity == response.StatusCode{
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return

		}
		if err!=nil{
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/admin/rooms", 302)
	}

	http.Redirect(w, r, "/admin/rooms", 302)

}
//getting the datas
func (arh *AdminRoomHandler) GetRoomAndRoomCategories()([]entity.Room,[]entity.Type){
	response, err := http.Get("http://localhost:9090/room/rooms")
	if response.StatusCode == http.StatusNotFound {

		return nil, nil
	}
	if http.StatusUnprocessableEntity == response.StatusCode {
		return nil, nil
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, nil
	}
	var roomslist []entity.Room
	err = json.Unmarshal(responseData, &roomslist)
	if err != nil {
		return nil, nil
	}
	//for room types
	client:=&http.Client{}
	req, err := http.NewRequest("GET","http://localhost:9090/room/types",nil)
	response2, err:=client.Do(req)
	if http.StatusNotFound == response2.StatusCode {
		return roomslist, nil
	}
	if http.StatusUnprocessableEntity == response2.StatusCode {
		return roomslist, nil
	}
	responseData2, err := ioutil.ReadAll(response2.Body)
	if err != nil {
		return roomslist, nil
	}
	var roomTypes []entity.Type
	//var roomTypes2 []entity.Type
	err = json.Unmarshal(responseData2, &roomTypes)
	if err != nil {
		return roomslist, nil
	}
	return roomslist, roomTypes



}
func writeFile(mf *multipart.File, fname string) {

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
	_, _ = io.Copy(image, *mf)
}