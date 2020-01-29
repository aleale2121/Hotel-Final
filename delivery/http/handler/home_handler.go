package handler

import (
	"fmt"
	"github.com/yuidegm/Hotel-Rental-Managemnet-System/entity"
	"github.com/yuidegm/Hotel-Rental-Managemnet-System/rate"
	"html/template"
	"net/http"
)

// HomePageHandler handles menu related requests
type HomePageHandler struct {
	tmpl        *template.Template
	rSrv rate.RateService
}
// HomePageHandler initializes and returns new HomePageHandler
func NewHomePageHandler(T *template.Template, CS rate.RateService) *HomePageHandler {
	return &HomePageHandler{tmpl: T, rSrv: CS}
}


// this handle home page on route
func (entitypointer *HomePageHandler) Index(w http.ResponseWriter, r *http.Request) {

	var v *entity.Rating
	val,ee:=entitypointer.rSrv.GetAllRatings(v)
	fmt.Println(val,"home index")
	if len(ee)>0{
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	var total int
    for _,valueofrate:=range val{
    	total+=valueofrate.RateValue
	}
	  var avg float32=float32(total)/float32(len(val))
	 averagefin:=int(avg)
	entitypointer.tmpl.ExecuteTemplate(w, "index.layout", averagefin)
}



