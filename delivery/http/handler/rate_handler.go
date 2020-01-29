package handler

import (
	"github.com/aleale2121/Hotel-Final/entity"
	"github.com/aleale2121/Hotel-Final/rate"
	"html/template"
	"net/http"
	"strconv"
)

// UserRateHandler handles category handler admin requests
type UserRateHandler struct {
	tmpl    *template.Template
	userHandler *UserHandler
	rateSrv rate.RateService
}


// NewRatingHandler initializes and returns new UserRateHandler
func NewRatingHandler(template *template.Template,userHandler *UserHandler, rs rate.RateService) *UserRateHandler {
	return &UserRateHandler{tmpl: template,userHandler:userHandler, rateSrv: rs}
}

// AddRate handle requests on route /rate
func (rth *UserRateHandler) AddRate(w http.ResponseWriter, r *http.Request)  {
	rate :=entity.Rating{}
    id:= rth.userHandler.loggedInUser.Id

	rateValueFromHidden:=r.FormValue("hidden_rate_value_container")
	if r.Method == "POST" && rateValueFromHidden!="0"{
		rateValue,err:=strconv.Atoi(rateValueFromHidden)
		if err!=nil{
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		userPreviousRate, _:= rth.rateSrv.GetUserRateValue(uint(id))
		if userPreviousRate ==nil{
			 rate.RateValue= rateValue
			 rate.UserId= uint(id)
			 _, err:= rth.rateSrv.AddRate(&rate)
			 if len(err)>0 {
				 http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				 return
			 }

		 } else{
			  err:= rth.rateSrv.UpdateUserRateValue(uint(id), rateValue)
			  if len(err)>0{
				  http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				  return
			  }
		 }
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}




