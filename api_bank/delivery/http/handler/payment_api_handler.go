package handler

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/aleale2121/Hotel-Final/api_bank/bank"
	"github.com/aleale2121/Hotel-Final/api_bank/entity"
	"github.com/aleale2121/Hotel-Final/utils"
	"net/http"
	"strconv"
)

// AdminCommentHandler handles comment related http requests
type PaymentApiHandler struct {
	papi     bank.Services
}

func NewPaymentApiHandler(services bank.Services) *PaymentApiHandler{
	return &PaymentApiHandler{papi:services}
}


func (pay *PaymentApiHandler) ReturnSingleCustomer(w http.ResponseWriter, r *http.Request,ps httprouter.Params) {
	fmt.Println("Single User retrieval")
	id, err := strconv.Atoi(ps.ByName("id"))
	customer2, errs := pay.papi.RetrieveAccountFromBank(int64(id))

	if len(errs)>0{
		fmt.Println("gorm error")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return

	}
	output, err := json.MarshalIndent(customer2, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
	return

}

// PutPayments handles PUT /user/payments/:id request
func (pay *PaymentApiHandler) PutPayments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	body:=utils.BodyParser(r)
	var customer entity.Customer
	err :=json.Unmarshal(body,&customer)
	if err!=nil {

		utils.ToJson(w,err.Error(),http.StatusUnprocessableEntity)
		return
	}
	fmt.Println("payment api",customer)
	 errs1 := pay.papi.UpdateUserAccount(customer.AccountNumber,customer.AccountBalance)

	if errs1!=nil {
		fmt.Println("gorm error")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return

	}
	output, err := json.MarshalIndent(customer, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)
	return
}
