
package main

import (
	"fmt"
	"github.com/aleale2121/Hotel-Final/api/delivery/http/handler"
	"github.com/aleale2121/Hotel-Final/api/entity"
	"github.com/aleale2121/Hotel-Final/api/events/events_repository"
	"github.com/aleale2121/Hotel-Final/api/events/events_services"
	"github.com/aleale2121/Hotel-Final/api/news/news_repository"
	"github.com/aleale2121/Hotel-Final/api/news/news_services"
	or "github.com/aleale2121/Hotel-Final/api/order/repository"
	os "github.com/aleale2121/Hotel-Final/api/order/services"
	rrepim "github.com/aleale2121/Hotel-Final/api/rate/repository"
	rsrvim "github.com/aleale2121/Hotel-Final/api/rate/services"
	rp "github.com/aleale2121/Hotel-Final/api/rooms/repository"
	rs "github.com/aleale2121/Hotel-Final/api/rooms/services"
	up "github.com/aleale2121/Hotel-Final/api/user/repository"
	us "github.com/aleale2121/Hotel-Final/api/user/services"
	//_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"html/template"
	"net/http"
)
var tmpl *template.Template
var admintempl *template.Template

func createTables(dbconn *gorm.DB) []error {


	dbconn.Debug().DropTableIfExists(&entity.Role{},&entity.User{},&entity.Order{},&entity.Type{}, &entity.Room{},&entity.Rating{})
	errs := dbconn.Debug().CreateTable(&entity.Role{},&entity.User{},&entity.Order{},&entity.Type{}, &entity.Room{},
		&entity.Events{},&entity.News{},&entity.Comments{},&entity.Rating{}).GetErrors()
	dbconn.Debug().Model(&entity.Room{}).AddForeignKey("type_id","types(Id)","cascade","cascade")
	dbconn.Debug().Model(&entity.Order{}).AddForeignKey("user_id","users(Id)","cascade","cascade")
	dbconn.Debug().Model(&entity.Order{}).AddForeignKey("room_id","rooms(Id)","cascade","cascade")
	if errs != nil {
		return errs
	}
	return nil
}

func main() {

	dbconn, err := gorm.Open("postgres",
		"postgres://postgres:root@localhost/Hotel?sslmode=disable")
	if dbconn !=nil {
		defer dbconn.Close()

	}
   // dbconn.Debug().DropTableIfExists(&entity.Order{})
  // dbconn.Debug().CreateTable(&entity.Order{})
	// createTables(dbconn)
	if err != nil {
		panic(err)
	}

    ///////////////..... EVENTS API...........////////////
	eventRepo := events_repository.NewEventsGormRepo(dbconn)
	eventSrv := events_services.NewEventService(eventRepo)
	adminEventHandler := handler.NewAdminEventsHandlerApi(eventSrv)

	//rate
	dbRepor := rrepim.NewRateGormRepo(dbconn)
	rateSrv := rsrvim.NewRateService(dbRepor)
	rateHandler := handler.NewRateApiHandler(rateSrv)

	router := httprouter.New()

	//rates api router
	router.GET("/users/rates", rateHandler.GetAllRates)
	router.GET("/users/rates/:id", rateHandler.ReturnSingleRates)
	router.POST("/users/rates", rateHandler.CreateRates)
	router.PUT("/users/rates/:id", rateHandler.PutRates)
	router.DELETE("/users/rates/:id", rateHandler.DeleteRates)

	router.GET("/hotel/events/:id", adminEventHandler.GetEventById)
	router.GET("/hotel/events", adminEventHandler.GetEvents)
	router.PUT("/hotel/events/:id", adminEventHandler.PutEvents)
	router.POST("/hotel/events", adminEventHandler.PostEvents)
	router.DELETE("/hotel/events/:id", adminEventHandler.DeleteEvents)

	//////////..........NEWS API ..............////////////////
	newsRepo := news_repository.NewNewssGormRepo(dbconn)
	newsSrv := news_services.NewNewsService(newsRepo)
	adminNewsHandler := handler.NewAdminNewsHandlerApi(newsSrv)

	router.GET("/hotel/newss/:id", adminNewsHandler.GetNewsById)
	router.GET("/hotel/newss", adminNewsHandler.GetNews)
	router.PUT("/hotel/newss/:id", adminNewsHandler.PutNews)
	router.POST("/hotel/newss", adminNewsHandler.PostNews)
	router.DELETE("/hotel/newss/:id", adminNewsHandler.DeleteNews)
	//
	reserveRep:=or.NewOrderGormRepo(dbconn)
	reserveServ:=os.NewOrderService(reserveRep)
	reserveHandler:=handler.NewRoomReserveHandler(reserveServ)
	//order handler
	router.GET("/user/reserve", reserveHandler.GetReservations)
	router.POST("/user/reserve/new", reserveHandler.PostOrder)
	router.GET("/user/reserves/:id", reserveHandler.GetReservation)
	router.GET("/user/reserved/:id", reserveHandler.GetRoomReservation)
	router.PUT("/user/reserve/update/:id", reserveHandler.PutReservationOrder)
	router.DELETE("/user/reserve/delete/:id", reserveHandler.DeleteReservation)

	//user handler
	userRep:=up.NewUserGormRepo(dbconn)
	userServ:=us.NewUserService(userRep)
	userHandler:=handler.NewUserApiHandler(userServ)

	router.GET("/user/users", userHandler.GetUsers)
	router.GET("/user/user/:id", userHandler.GetUser)
	router.GET("/user/roles", userHandler.GetUserRoles)
	router.GET("/user/email/:email", userHandler.IsEmailExists)
	router.GET("/user/phone/:phone", userHandler.IsPhoneExists)
	router.POST("/user/check",userHandler.GetUserByUsernameAndPassword)
	router.POST("/user/new", userHandler.PostUser)
	router.PUT("/user/update/:id", userHandler.PutUser)
	router.DELETE("/user/delete/:id", userHandler.DeleteUser)
    //user session handler
	sessionRep:=up.NewSessionGormRepo(dbconn)
	sessionServ:=us.NewSessionService(sessionRep)
	sessionHandler:=handler.NewSessionApiHandler(sessionServ)
	router.GET("/user/session/:id", sessionHandler.GetSessionByName)
	router.POST("/user/session/new", sessionHandler.PostSession)
	router.DELETE("/user/session/delete/:id", sessionHandler.DeleteSession)
	//user role handler
	roleRep:=up.NewRoleGormRepo(dbconn)
	roleServ:=us.NewRoleService(roleRep)
	roleHandler:=handler.NewRoleApiHandler(roleServ)
	router.GET("/roles", roleHandler.GetRoles)
	router.GET("/roles/:name", roleHandler.GetRoleByName)
	router.GET("/role/:id",roleHandler.GetRoleByID)
	router.POST("/role/new", roleHandler.PostRole)
	router.PUT("/role/update/:id", roleHandler.PutRole)
	router.DELETE("/role/delete/:id", roleHandler.DeleteRole)
	//admin room handler
	roomRep:=rp.NewRoomGormRepo(dbconn)
	roomSer:=rs.NewRoomServiceImpl(roomRep)
	roomHandler:=handler.NewAdminRoomHandler(roomSer)
	router.GET("/room/rooms", roomHandler.AdminRooms)
	router.GET("/room/types", roomHandler.AdminRoomTypes)
	router.POST("/room/new", roomHandler.AdminRoomNew)
	router.PUT("/room/update/:id", roomHandler.AdminRoomsUpdate)
	router.DELETE("/room/delete/:id", roomHandler.AdminRoomsDelete)

	fmt.Println("server started at port 9090")

	_ = http.ListenAndServe(":9090", router)
}