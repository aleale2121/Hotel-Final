
package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/aleale2121/Hotel-Final/comment/com_Repository"
	"github.com/aleale2121/Hotel-Final/comment/com_Service"
	"github.com/aleale2121/Hotel-Final/delivery/http/handler"
	"github.com/aleale2121/Hotel-Final/entity"
	"github.com/aleale2121/Hotel-Final/events/events_repository"
	"github.com/aleale2121/Hotel-Final/events/events_services"
	"github.com/aleale2121/Hotel-Final/news/news_repository"
	"github.com/aleale2121/Hotel-Final/news/news_services"
	rrepim "github.com/aleale2121/Hotel-Final/rate/repository"
	rsrvim "github.com/aleale2121/Hotel-Final/rate/services"
	"github.com/aleale2121/Hotel-Final/rtoken"
	"html/template"
	"net/http"
	"time"
)
var templ *template.Template
var admintempl *template.Template
var usertempl *template.Template

var err error
func createTables(dbconn *gorm.DB) []error {


	dbconn.Debug().DropTableIfExists(&entity.Role{},&entity.User{},&entity.Order{},&entity.Type{}, &entity.Room{})
	errs := dbconn.Debug().CreateTable(&entity.Role{},&entity.User{},&entity.Order{},&entity.Type{}, &entity.Room{},
	&entity.Events{},&entity.News{},&entity.Comments{}).GetErrors()
	dbconn.Debug().Model(&entity.Room{}).AddForeignKey("type_id","types(Id)","cascade","cascade")
	dbconn.Debug().Model(&entity.Order{}).AddForeignKey("user_id","users(Id)","cascade","cascade")
	dbconn.Debug().Model(&entity.Order{}).AddForeignKey("room_id","rooms(Id)","cascade","cascade")
	if errs != nil {
		return errs
	}
	return nil
}
func init(){

}


func main() {
	dbconn, err := gorm.Open("postgres", "postgres://postgres:root@localhost/Hotel?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer dbconn.Close()
	fmt.Println("connected")
     //createTables(dbconn)
	csrfSignKey := []byte(rtoken.GenerateRandomID(32))
	templ = template.Must(template.ParseGlob("ui/templates/user_template/*"))
	admintempl=template.Must(template.ParseGlob("ui/templates/admin_template/*"))
	usertempl=template.Must(template.ParseGlob("ui/templates/common_template/*"))
	fs := http.FileServer(http.Dir("ui/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))



	//rate
	dbRepor:=rrepim.NewRateGormRepo(dbconn)
	rateSrv:=rsrvim.NewRateService(dbRepor)

	///user handlers //menuhandler
	ComRepo := com_Repository.NewGormComRepositoryImpl(dbconn)
	ComServ :=com_Service.NewGormComServiceImpl(ComRepo)
	eventsRepo := events_repository.NewGormEventsRepositoryImpl(dbconn)
	eventsServ :=events_services.NewGormNewsServiceImpl(eventsRepo)
	newsRepo := news_repository.NewGormNewsRepositoryImpl(dbconn)
	newsServ := news_services.NewGormNewsServiceImpl(newsRepo)


	sess := configSess()

	userhandler:=handler.NewUserHandler(usertempl,sess,csrfSignKey)
	adminUserHandler:=handler.NewUserHandler(admintempl,sess,csrfSignKey)
	menuHandler:=handler.NewCMenuHandler(templ,userhandler,rateSrv,csrfSignKey)
	rateHandler:=handler.NewRatingHandler(templ,userhandler,rateSrv)
	usernewsHandler := handler.NewMenuHandler(templ, newsServ,userhandler)
	userEventHandler := handler.NewEventsHandler(templ, eventsServ,userhandler)
	userComHandler := handler.NewUserComHandler(templ, ComServ,userhandler,csrfSignKey)

	http.HandleFunc("/login",userhandler.LoginGetHandler)
	http.HandleFunc("/signup",userhandler.SignupGetHandler)
	http.HandleFunc("/",menuHandler.Home)
	http.HandleFunc("/newss",usernewsHandler.News_page)
	http.HandleFunc("/event",userEventHandler.Event_page)
	http.Handle("/rooms",userhandler.Authenticated(http.HandlerFunc(menuHandler.Rooms)))
	http.HandleFunc("/contact",userComHandler.UserCom)
	http.HandleFunc("/map",menuHandler.Maps)
	http.Handle("/rate",userhandler.Authenticated(http.HandlerFunc(rateHandler.AddRate)))
	http.Handle("/logout",userhandler.Authenticated(http.HandlerFunc(userhandler.Logout)))
	/////...............USER ORDERS HANDLER............./////
	reserveHandler:=handler.NewRoomReserveHandler(usertempl,userhandler,csrfSignKey)
	http.Handle("/user/reserve", userhandler.Authenticated(http.HandlerFunc(reserveHandler.GetReservations)))
	http.Handle("/user/reserve/new", userhandler.Authenticated(http.HandlerFunc(reserveHandler.PostOrder)))
	http.Handle("/user/reserve/update/:id", userhandler.Authenticated(http.HandlerFunc(reserveHandler.PutReservationOrder)))
	http.Handle("/user/reserve/delete/:id", userhandler.Authenticated(http.HandlerFunc(reserveHandler.DeleteReservation)))

	/////...............USER PROFILE HANDLER............./////
	userprofile:=handler.NewUserProfileHandler(templ,userhandler)
	http.Handle("/profile",userhandler.Authenticated(http.HandlerFunc(userprofile.UserInfo)))
	///.................ADMIN ROOM HANDLER................//////
	adminRoomHandler := handler.NewAdminRoomHandler(admintempl,csrfSignKey)
	http.Handle("/admin/rooms", userhandler.Authenticated(userhandler.Authorized(http.HandlerFunc( adminRoomHandler.AdminRooms))))
	http.Handle("/admin/rooms/new",  userhandler.Authenticated(userhandler.Authorized(http.HandlerFunc(adminRoomHandler.AdminRoomNew))))
	http.Handle("/admin/rooms/update", userhandler.Authenticated(userhandler.Authorized(http.HandlerFunc( adminRoomHandler.AdminRoomsUpdate))))
	http.Handle("/admin/rooms/delete", userhandler.Authenticated(userhandler.Authorized(http.HandlerFunc( adminRoomHandler.AdminRoomsDelete))))

	//////////-----------ADMIN EVENTS HANDLER-----------------////////////

	adminEventsHandler := handler.NewAdminEventsHandler(admintempl, eventsServ,csrfSignKey)
	http.Handle("/admin/events",  userhandler.Authenticated(userhandler.Authorized(http.HandlerFunc(adminEventsHandler.AdminEvents))))
	http.Handle("/admin/events/new",  userhandler.Authenticated(userhandler.Authorized(http.HandlerFunc(adminEventsHandler.AdminEventsNew))))
	http.Handle("/admin/events/update",  userhandler.Authenticated(userhandler.Authorized(http.HandlerFunc(adminEventsHandler.AdminEventsUpdate))))
	http.Handle("/admin/events/delete",  userhandler.Authenticated(userhandler.Authorized(http.HandlerFunc(adminEventsHandler.AdminEventsDelete))))

	///////////............ADMIN COMMENTS HANDLER............../////////

	adminComHandler := handler.NewAdminComHandler(admintempl, ComServ)
	http.Handle("/admin/com",  userhandler.Authenticated(userhandler.Authorized(http.HandlerFunc(adminComHandler.AdminCom))))
	http.Handle("/admin/com/delete",  userhandler.Authenticated(userhandler.Authorized(http.HandlerFunc(adminComHandler.AdminComDelete))))

	//////// ...........ADMIN NEWS HANDLER.........//////////////

	adminNewsHandler := handler.NewAdminNewsHandler(admintempl, newsServ,csrfSignKey)
	http.Handle("/admin/newss",  userhandler.Authenticated(userhandler.Authorized(http.HandlerFunc(adminNewsHandler.AdminNews))))
	http.Handle("/admin/newss/new",  userhandler.Authenticated(userhandler.Authorized(http.HandlerFunc(adminNewsHandler.AdminNewsNew))))
	http.Handle("/admin/newss/update", userhandler.Authenticated(userhandler.Authorized(http.HandlerFunc( adminNewsHandler.AdminNewsUpdate))))
	http.Handle("/admin/newss/delete",  userhandler.Authenticated(userhandler.Authorized(http.HandlerFunc(adminNewsHandler.AdminNewsDelete))))
	//////// ...........ADMIN ORDERS HANDLER.........//////////////
	adminOrderHandler :=handler.NewAdminOrderHandler(admintempl)
	http.Handle("/admin/orders", userhandler.Authenticated(userhandler.Authorized(http.HandlerFunc(adminOrderHandler.AdminOrders))))
	http.Handle("/admin/orders/delete", userhandler.Authenticated(userhandler.Authorized(http.HandlerFunc(adminOrderHandler.AdminOrderDelete))))


	///////................ADMIN USERS HANDLER..........////////

	http.Handle("/admin/users", userhandler.Authenticated(userhandler.Authorized(http.HandlerFunc(adminUserHandler.AdminUsers))))
	http.Handle("/admin/users/new", userhandler.Authenticated(userhandler.Authorized(http.HandlerFunc(adminUserHandler.AdminUsersNew))))
	http.Handle("/admin/users/update", userhandler.Authenticated(userhandler.Authorized(http.HandlerFunc(adminUserHandler.AdminUsersUpdate))))
	http.Handle("/admin/users/delete", userhandler.Authenticated(userhandler.Authorized(http.HandlerFunc(adminUserHandler.AdminUsersDelete))))



	_ = http.ListenAndServe(":8080", nil)


}

func configSess() *entity.Session {
	tokenExpires := time.Now().Add(time.Minute * 30).Unix()
	sessionID := rtoken.GenerateRandomID(32)
	signingString, err := rtoken.GenerateRandomString(32)
	if err != nil {
		panic(err)
	}
	signingKey := []byte(signingString)

	return &entity.Session{
		Expires:    tokenExpires,
		SigningKey: signingKey,
		UUID:       sessionID,
	}
}



