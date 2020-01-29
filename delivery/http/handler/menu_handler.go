package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/aleale2121/Hotel-Final/entity"
	"github.com/aleale2121/Hotel-Final/form"
	"github.com/aleale2121/Hotel-Final/rate"
	"github.com/aleale2121/Hotel-Final/rtoken"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)
// MenuHandler handles menu related requests
type MenuHandler struct {
	temp        *template.Template
	userHandler *UserHandler
	rSrv rate.RateService
	csrfSignKey []byte
}
type RoomData struct {
	Roomslist []entity.Room
	FormInput  form.Input
	AverageRating int
	IsLogged string
	ActionMode string
}
func NewCMenuHandler(T *template.Template,userHandler *UserHandler,rSrv rate.RateService,csrfSignKey []byte) *MenuHandler {
	return &MenuHandler{temp: T,userHandler:userHandler,rSrv:rSrv,csrfSignKey:csrfSignKey}
}

func(rrh *MenuHandler) Home(w http.ResponseWriter, r *http.Request) {
	isLogged :="false"
	if rrh.userHandler.loggedIn(r) {
		isLogged="true"
	}
	var v *entity.Rating
	allRatings,ee:=rrh.rSrv.GetAllRatings(v)
	if len(ee)>0{
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	var total int
	for _, rate :=range allRatings {
		total+= rate.RateValue
	}
	var avg=float32(total)/float32(len(allRatings))
	roomData:=RoomData{AverageRating:int(avg),IsLogged:isLogged}

	_ = rrh.temp.ExecuteTemplate(w, "index.layout",roomData)
}
func(rrh *MenuHandler) Contactus(w http.ResponseWriter, r *http.Request) {
	_ = rrh.temp.ExecuteTemplate(w, "contact.layout", nil)
}
func(rrh *MenuHandler) Rooms(w http.ResponseWriter, r *http.Request) {

	token, err := rtoken.CSRFToken(rrh.csrfSignKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	isLogged :="false"
	if rrh.userHandler.loggedIn(r) {
		isLogged="true"
	}
	if r.Method == http.MethodPost {
		v := url.Values{}
		v.Add("arrival", r.FormValue("arrival"))
		v.Add("departure", r.FormValue("departure"))
		v.Add("adults", r.FormValue("adults"))
		v.Add("childs", r.FormValue("childs"))
		v.Add("account", r.FormValue("account"))
		v.Add("room_price1", r.FormValue("room_price1"))

		RoomReserveForm := form.Input{Values: v, VErrors: form.ValidationErrors{}}
		RoomReserveForm.Required("arrival", "departure","adults","childs","account",)
		RoomReserveForm.CSRF = token
		rooms:=rrh.GetRooms2()
		roomRes:=RoomData{Roomslist: rooms, FormInput: RoomReserveForm,ActionMode:"RoomReserve",IsLogged:isLogged}
		if !RoomReserveForm.Valid() {
			rrh.temp.ExecuteTemplate(w, "room_index2.layout", roomRes)
			return
		}

		order :=entity.Order{}
		timeStampString1 :=r.FormValue("arrival")
		timeStampString2 :=r.FormValue("departure")
		timeStampArrival, _ := time.Parse("2006-01-02 15:04", timeStampString1)
		timeStampDest, _ := time.Parse("2006-01-02 15:04", timeStampString2)

		hr, min, sec := timeStampArrival.Clock()
		year, month, day := timeStampArrival.Date()

		hr2, min2, sec2 := timeStampDest.Clock()
		year2, month2, day2 := timeStampDest.Date()

		seconds1:=year*365*24*60*60+ int(month)*30*24*60*60+day*24*60*60+hr*60*60 +min*60+sec
		seconds2:=year2*365*24*60*60+ int(month2)*30*24*60*60+day2*24*60*60+hr2*60*60 +min2*60+sec2

		if(seconds2-seconds1)<3600{
			RoomReserveForm.VErrors.Add("generics", " You should stay at least one hour in our room ")
			roomRes=RoomData{Roomslist: rooms, FormInput: RoomReserveForm,ActionMode:"RoomReserve",IsLogged:isLogged}
			rrh.temp.ExecuteTemplate(w, "room_index2.layout", roomRes)
			return
		}

		hrnow, minnow,secnow:=time.Now().Clock()
		yrnow,mnnow,daynow:=time.Now().Date()
		totalsecondsnow:=yrnow*365*24*60*60+ int(mnnow)*30*24*60*60+daynow*24*60*60+hrnow*60*60 +minnow*60+secnow
		fmt.Println("106")
		if(totalsecondsnow>seconds1)||(totalsecondsnow>seconds2){
			RoomReserveForm.VErrors.Add("generics", "please choose a date and time now or in the then time")
			roomRes=RoomData{Roomslist: rooms, FormInput: RoomReserveForm,ActionMode:"RoomReserve",IsLogged:isLogged}
			rrh.temp.ExecuteTemplate(w, "room_index2.layout", roomRes)
			return
		}
		accountNumber, _ := strconv.Atoi(r.FormValue("account"))
		price, _ :=strconv.ParseFloat(r.FormValue("room_price1"),64)

		roomid, _ :=strconv.Atoi(r.FormValue("roomId"))
		order.RoomId= uint32(roomid)
		order.UserId= rrh.userHandler.loggedInUser.Id
		adults,_:=strconv.Atoi(r.FormValue("adults"))
		childs,_:=strconv.Atoi(r.FormValue("childs"))
		order.Adults = uint(adults)
		order.Child= uint(childs)
		order.ArrivalDate=timeStampArrival
		order.DepartureDate=timeStampDest

		client:=&http.Client{}
		url:=fmt.Sprintf("http://localhost:8181/bank/customer/%d",int64(accountNumber))
		request, _:=http.NewRequest("GET",url,nil)
		response, err := client.Do(request)
		//response, err := client.Do(request)
		fmt.Println("130")
		if err !=nil{
			RoomReserveForm.VErrors.Add("generics", "Temporary server error")
			roomRes=RoomData{Roomslist: rooms, FormInput: RoomReserveForm,ActionMode:"RoomReserve",IsLogged:isLogged}
			rrh.temp.ExecuteTemplate(w, "room_index2.layout", roomRes)
			return
		}
		fmt.Println("137")
		if response.StatusCode==http.StatusNotFound  {
			RoomReserveForm.VErrors.Add("account", "Account Does not Exist")
			roomRes=RoomData{Roomslist: rooms, FormInput: RoomReserveForm,ActionMode:"RoomReserve",IsLogged:isLogged}
			rrh.temp.ExecuteTemplate(w, "room_index2.layout", roomRes)
			return
		}
		fmt.Println("144")
		if http.StatusUnprocessableEntity == response.StatusCode {
			RoomReserveForm.VErrors.Add("generics", "Invalid  inputs ")
			roomRes=RoomData{Roomslist: rooms, FormInput: RoomReserveForm,ActionMode:"RoomReserve",IsLogged:isLogged}
			rrh.temp.ExecuteTemplate(w, "room_index2.layout", roomRes)
			return

		}
		fmt.Println("152")
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			RoomReserveForm.VErrors.Add("generics", "Temporary Server Error ")
			roomRes=RoomData{Roomslist: rooms, FormInput: RoomReserveForm,ActionMode:"RoomReserve",IsLogged:isLogged}
			rrh.temp.ExecuteTemplate(w, "room_index2.layout", roomRes)
			return
		}
		fmt.Println("160")
		var customer entity.Customer
		err = json.Unmarshal(responseData, &customer)
		if err != nil {
			RoomReserveForm.VErrors.Add("generics", "Temporary Server Error")
			roomRes=RoomData{Roomslist: rooms, FormInput: RoomReserveForm,ActionMode:"RoomReserve",IsLogged:isLogged}
			rrh.temp.ExecuteTemplate(w, "room_index2.layout", roomRes)
			return
		}
		fmt.Println("169")
		if float32(price )>customer.AccountBalance {
			RoomReserveForm.VErrors.Add("generics", "There is no enough money in your account")
			roomRes=RoomData{Roomslist: rooms, FormInput: RoomReserveForm,ActionMode:"RoomReserve",IsLogged:isLogged}
			rrh.temp.ExecuteTemplate(w, "room_index2.layout", roomRes)
			return
		}
		fmt.Println("176")
		//check if the room is reserved
		url=fmt.Sprintf("http://localhost:9090/user/reserved/%d",roomid)
		req,err:=http.NewRequest("GET",url,nil)
		response, err=client.Do(req)
		if response.StatusCode==http.StatusNotFound  {
			RoomReserveForm.VErrors.Add("generics", "Temporary Server Error ")
			roomRes=RoomData{Roomslist: rooms, FormInput: RoomReserveForm,ActionMode:"RoomReserve",IsLogged:isLogged}
			rrh.temp.ExecuteTemplate(w, "room_index2.layout", roomRes)
			return
		}
		fmt.Println("187")
		responseData, err = ioutil.ReadAll(response.Body)
		if err != nil {
			RoomReserveForm.VErrors.Add("generics", "Temporary Server Error ")
			roomRes=RoomData{Roomslist: rooms, FormInput: RoomReserveForm,ActionMode:"RoomReserve",IsLogged:isLogged}
			rrh.temp.ExecuteTemplate(w, "room_index2.layout", roomRes)
			return
		}
		fmt.Println("195")
		var orders []entity.Order
		err =json.Unmarshal(responseData,&orders)
		if err!=nil{
			RoomReserveForm.VErrors.Add("generics", "Temporary Server Error ")
			roomRes=RoomData{Roomslist: rooms, FormInput: RoomReserveForm,ActionMode:"RoomReserve",IsLogged:isLogged}
			rrh.temp.ExecuteTemplate(w, "room_index2.layout", roomRes)
			return
		}
		fmt.Println("204")
		for _, order := range orders {
			hr3, min3, sec3 := order.ArrivalDate.Clock()
			year3, month3, day3 := order.ArrivalDate.Date()

			hr4, min4, sec4 := order.DepartureDate.Clock()
			year4, month4, day4 := order.DepartureDate.Date()

			seconds3:=year3*365*24*60*60+ int(month3)*30*24*60*60+day3*24*60*60+hr3*60*60 +min3*60+sec3
			seconds4:=year4*365*24*60*60+ int(month4)*30*24*60*60+day4*24*60*60+hr4*60*60 +min4*60+sec4

			if!(((seconds1<seconds3)&&(seconds2<seconds4))||((seconds1>seconds3)&&(seconds2>seconds4))){
				RoomReserveForm.VErrors.Add("generics", "The room has already reserved please try another time")
				roomRes=RoomData{Roomslist: rooms, FormInput: RoomReserveForm,ActionMode:"RoomReserve",IsLogged:isLogged}
				rrh.temp.ExecuteTemplate(w, "room_index2.layout", roomRes)
				return
			}
		}
		fmt.Println("204")
		///bank transfer
		url=fmt.Sprintf("http://localhost:8181/bank/customer/pay")

		usermoney:=customer.AccountBalance-float32(price )
		customer.AccountBalance=usermoney
		output,err :=json.MarshalIndent(customer,"", "\t\t")
		if err!=nil{
			RoomReserveForm.VErrors.Add("generics", "Internal Server Error")
			roomRes=RoomData{Roomslist: rooms, FormInput: RoomReserveForm,ActionMode:"RoomReserve",IsLogged:isLogged}
			rrh.temp.ExecuteTemplate(w, "room_index2.layout", roomRes)
			return
		}
		fmt.Println("235")
		req,err=http.NewRequest("PUT",url,bytes.NewBuffer(output))
		response, err=client.Do(req)

		if response.StatusCode==http.StatusNotFound  {
			RoomReserveForm.VErrors.Add("generics", "Internal Server Error")
			roomRes=RoomData{Roomslist: rooms, FormInput: RoomReserveForm,ActionMode:"RoomReserve",IsLogged:isLogged}
			rrh.temp.ExecuteTemplate(w, "room_index2.layout", roomRes)
			return
		}
		fmt.Println("245")
		if http.StatusUnprocessableEntity == response.StatusCode {
			RoomReserveForm.VErrors.Add("generics", "Internal Server Error")
			roomRes=RoomData{Roomslist: rooms, FormInput: RoomReserveForm,ActionMode:"RoomReserve",IsLogged:isLogged}
			rrh.temp.ExecuteTemplate(w, "room_index2.layout", roomRes)
			return

		}
		///order
		output,err =json.MarshalIndent(order,"", "\t\t")
		if err!=nil{
			RoomReserveForm.VErrors.Add("generics", "Internal Server Error")
			roomRes=RoomData{Roomslist: rooms, FormInput: RoomReserveForm,ActionMode:"RoomReserve",IsLogged:isLogged}
			rrh.temp.ExecuteTemplate(w, "room_index2.layout", roomRes)
			return
		}
		fmt.Println("261")
		url=fmt.Sprintf("http://localhost:9090/user/reserve/new")
		req,err=http.NewRequest("POST",url,bytes.NewBuffer(output))
		response, err=client.Do(req)

		if response.StatusCode==http.StatusNotFound  {
			RoomReserveForm.VErrors.Add("generics", "Internal Server Error")
			roomRes=RoomData{Roomslist: rooms, FormInput: RoomReserveForm,ActionMode:"RoomReserve",IsLogged:isLogged}
			rrh.temp.ExecuteTemplate(w, "room_index2.layout", roomRes)
			return
		}
		fmt.Println("272")
		if http.StatusCreated != response.StatusCode {
			RoomReserveForm.VErrors.Add("generics", "Internal Server Error")
			roomRes=RoomData{Roomslist: rooms, FormInput: RoomReserveForm,ActionMode:"RoomReserve",IsLogged:isLogged}
			rrh.temp.ExecuteTemplate(w, "room_index2.layout", roomRes)
			return

		}
		fmt.Println("280")
		RoomReserveForm.VErrors.Add("success", "You have successfully reserved the room")
		roomRes=RoomData{Roomslist: rooms, FormInput: RoomReserveForm,ActionMode:"RoomSee",IsLogged:isLogged}
		rrh.temp.ExecuteTemplate(w, "room_index2.layout", roomRes)
		fmt.Println("284")
	}
	if r.Method==http.MethodGet {
		response, err := http.Get("http://localhost:9090/room/rooms")
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

		var roomslist []entity.Room
		err = json.Unmarshal(responseData, &roomslist)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		roomFormInput := form.Input{
			Values:  nil,
			VErrors: nil,
			CSRF:    token,
		}
		roomdata := RoomData{Roomslist: roomslist, FormInput: roomFormInput,IsLogged:isLogged}
		_ = rrh.temp.ExecuteTemplate(w, "room_index2.layout", roomdata)
	}

}
func(rrh *MenuHandler) News(response http.ResponseWriter, request *http.Request) {
	_ = rrh.temp.ExecuteTemplate(response, "notification.html", nil)
}
func(rrh *MenuHandler) Maps(response http.ResponseWriter, request *http.Request) {
	_ = rrh.temp.ExecuteTemplate(response, "map_index.layout", nil)
}
func (rrh *MenuHandler) GetRooms2() []entity.Room {
	client:=http.Client{}
	req, err:=http.NewRequest(http.MethodGet,"http://localhost:9090/room/rooms",nil)
	if err != nil {
		return nil
	}
	response, err :=client.Do(req)
	if response.StatusCode == http.StatusNotFound {
		return nil
	}
	if http.StatusUnprocessableEntity == response.StatusCode {
		return nil
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil
	}
	var roomslist []entity.Room
	err = json.Unmarshal(responseData, &roomslist)
	if err != nil {
		return nil
	}
	return roomslist

}
