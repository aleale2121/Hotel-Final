package handler
import (
	"encoding/json"
	"fmt"
	"github.com/aleale2121/Hotel-Final/api/entity"
	"github.com/aleale2121/Hotel-Final/api/rate"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"

)

// RateApiHandler handles rate related http requests
type RateApiHandler struct {
	rhSrv     rate.RateService
}



// NewRateApiHandler returns new NewRateApiHandler object
func NewRateApiHandler(cmn rate.RateService) *RateApiHandler {
	return &RateApiHandler{rhSrv: cmn}
}
// GetAllRates handles GET /v1/user/rates request
func (rh *RateApiHandler) GetAllRates(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rates, errs := rh.rhSrv.Rate()

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(rates, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}

func (rh *RateApiHandler) ReturnSingleRates(w http.ResponseWriter, r *http.Request,ps httprouter.Params) {
	fmt.Println("in returning single paymenting method")
	// string to int
	id, err := strconv.Atoi(ps.ByName("id"))
	w.Header().Set("Content-Type", "applicatiuon/json")
	rates, errs := rh.rhSrv.Rate()
	if len(errs) > 0 {
		println(errs)
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	fmt.Println("in 60")


	for _, rating := range rates {

		if err == nil {
			if rating.UserId == uint(id) {
				output, err := json.MarshalIndent(rating, "", "\t\t")
				if err != nil {
					w.Header().Set("Content-Type", "application/json")
					http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				w.Write(output)
				return

			}

		}

	}

}

func (rh *RateApiHandler) CreateRates(w http.ResponseWriter, r *http.Request,ps httprouter.Params) {
	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	rates := &entity.Rating{}
	fmt.Println("create api")
	err := json.Unmarshal(body, rates)
	if err != nil {
		fmt.Println("create api",err)
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, errs := rh.rhSrv.AddRate(rates)
	
	if len(errs) > 0 {
		fmt.Println("create api ",105)
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	
	return

}
// PutRates handles PUT /v1/user/rates/:id request
func (rh *RateApiHandler) PutRates(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	rates, errs := rh.rhSrv.Rates(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	json.Unmarshal(body, &rates)
	rates, errs = rh.rhSrv.UpdateRates(rates)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	output, err := json.MarshalIndent(rates, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}



//DeleteRates handles DELETE /v1/user/rates/:id request
func (rh *RateApiHandler) DeleteRates(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, errs := rh.rhSrv.DeleteRates(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return
}