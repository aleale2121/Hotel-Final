//
package events_repository
//
//import (
//	"database/sql"
//	"errors"
//	"fmt"
//	"github.com/getme04/original/entity"
//)
//
//// RoomRepositoryImpl implements the rooms.RoomRepository interface
//type EventsRepositoryImpl struct {
//	Con *sql.DB
//}
//
//// NewRoomRepositoryImpl will create an object of PsqlRoomRepository
//func NewEventsRepositoryImpl(Conn *sql.DB) *EventsRepositoryImpl {
//	return &EventsRepositoryImpl{Con: Conn}
//}
//
//// News returns all rooms from the database
//func (rri *EventsRepositoryImpl) Events() ([]entity.Events, error) {
//
//	rows, err := rri.Con.Query("SELECT * FROM events;")
//	if err != nil {
//
//		return nil, errors.Event("Could not query the database")
//	}
//	fmt.Printf("data fetched exactlly")
//	defer rows.Close()
//
//	newsList := []entity.Events{}
//
//	for rows.Next() {
//		event := entity.Events{}
//		err = rows.Scan(&event.Id, &event.Header,&event.Description, &event.Image)
//		if err != nil {
//			return nil, err
//		}
//		newsList = append(newsList, event)
//	}
//
//	return newsList, nil
//}
//
//// NewsById returns a News with a given id
//func (rri *EventsRepositoryImpl) EventById(id int) (entity.Events, error) {
//
//	row := rri.Con.QueryRow("SELECT * FROM events WHERE id = $1", id)
//
//	Events := entity.Events{}
//
//	err := row.Scan(&Events.Id, &Events.Header, &Events.Description, &Events.Image)
//	if err != nil {
//		return Events, err
//	}
//
//	return Events, nil
//}
//
//// UpdateNews updates a given object with a new data
//func (rri *EventsRepositoryImpl) UpdateEvent(r entity.Events) error {
//
//	_, err := rri.Con.Exec("UPDATE events SET header=$1,description=$2,image=$3 WHERE id=$4",
//		r.Header, r.Description, r.Image, r.Id)
//	if err != nil {
//		return errors.Event("Update has failed")
//	}
//
//	return nil
//}
//
//// DeleteNews removes a News from a database by its id
//func (rri *EventsRepositoryImpl) DeleteEvent(id int) error {
//
//	_, err := rri.Con.Exec("DELETE FROM events WHERE id=$1", id)
//	if err != nil {
//		return errors.Event("Delete has failed")
//	}
//
//	return nil
//}
//
//// StoreNews stores new News information to database
//func (rri *EventsRepositoryImpl) StoreEvent(r entity.Events) error {
//
//	_, err := rri.Con.Exec("INSERT INTO events ( header,description,image) values($1, $2, $3)",
//		r.Header,r.Description, r.Image)
//	if err != nil {
//		return errors.Event("Insertion has failed")
//	}
//	//dd:="gg";
//	//_, errr := rri.Con.Exec("INSERT INTO dt ( dat) values($1)",dd);
//	//if errr != nil {
//	//	return errors.Event("Insertion has failed")
//	//}
//	return nil
//}
//
