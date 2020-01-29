package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"github.com/yuidegm/Hotel-Rental-Managemnet-System/api_bank/bank/repository"
	"github.com/yuidegm/Hotel-Rental-Managemnet-System/api_bank/bank/services"
	"github.com/yuidegm/Hotel-Rental-Managemnet-System/api_bank/delivery/http/handler"
	"net/http"
)

var dbconn *gorm.DB
var err error

func init(){

	dbconn,err= gorm.Open("postgres", "postgres://postgres:root@localhost/Bank?sslmode=disable")
	if err != nil {
		panic(err)
	}
	database:=dbconn.DB()

	err2:=database.Ping()
	if err2!=nil{
		panic(err2.Error())
	}
	fmt.Println("bank database  connected")

}

func main() {
	//dbconn.DropTableIfExists(&entity.Customer{})
    //dbconn.AutoMigrate(&entity.Customer{})
	//Payments
	bankRepo:=repository.NewBankGormRepo(dbconn)
    bankServ:=services.NewBankService(bankRepo)
    bankHandler:=handler.NewPaymentApiHandler(bankServ)
	// payments api router
	router := httprouter.New()
	router.PUT("/bank/customer/pay", bankHandler.PutPayments)
	router.GET("/bank/customer/:id", bankHandler.ReturnSingleCustomer)

	_ = http.ListenAndServe(":8181", router)
}
