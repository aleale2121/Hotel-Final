package repository
//
//import (
//	"database/sql"
//	"errors"
//	"fmt"
//	"github.com/aleale2121/webProject2019/entity"
//)
//
//// RoomRepositoryImpl implements the rooms.RoomRepository interface
//type RoomRepositoryImpl struct {
//	conn *sql.DB
//}
//
//// NewRoomRepositoryImpl will create an object of PsqlRoomRepository
//func NewRoomRepositoryImpl(Conn *sql.DB) *RoomRepositoryImpl {
//	return &RoomRepositoryImpl{conn: Conn}
//}
//
//// Rooms returns all rooms from the database
//func (rri *RoomRepositoryImpl) Rooms() ([]entity.Room, error) {
//
//	rows, err := rri.conn.Query("SELECT * FROM rooms;")
//	if err != nil {
//		return nil, errors.Event("Could not query the database")
//	}
//	defer rows.Close()
//
//	roomlist := []entity.Room{}
//
//	for rows.Next() {
//		room := entity.Room{}
//		err = rows.Scan(&room.ID, &room.RoomNumber,&room.Price, &room.Description, &room.Image, &room.Type)
//		if err != nil {
//			return nil, err
//		}
//		roomlist = append(roomlist, room)
//	}
//
//	return roomlist, nil
//}
//
//// Room returns a room with a given id
//func (rri *RoomRepositoryImpl) Room(id int) (entity.Room, error) {
//
//	row := rri.conn.QueryRow("SELECT * FROM rooms WHERE id = $1", id)
//
//	room := entity.Room{}
//
//	err := row.Scan(&room.ID, &room.RoomNumber,&room.Rating, &room.Description, &room.Image, &room.Type)
//	if err != nil {
//		return room, err
//	}
//
//	return room, nil
//}
//
//// UpdateRoom updates a given object with a new data
//func (rri *RoomRepositoryImpl) UpdateRoom(r entity.Room) error {
//
//	_, err := rri.conn.Exec("UPDATE rooms SET room_number=$1,price=$2,description=$3, image=$4,rmtype=$5 WHERE id=$6",
//		r.RoomNumber, r.Price, r.Description, r.Image, r.Type, r.ID)
//	if err != nil {
//		return errors.Event("Update has failed")
//	}
//
//	return nil
//}
//
//// DeleteRoom removes a room from a database by its id
//func (rri *RoomRepositoryImpl) DeleteRoom(id int) error {
//
//	_, err := rri.conn.Exec("DELETE FROM rooms WHERE id=$1", id)
//	if err != nil {
//		return errors.Event("Delete has failed")
//	}
//
//	return nil
//}
//
//// StoreRoom stores new room information to database
//func (rri *RoomRepositoryImpl) StoreRoom(r entity.Room) error {
//	fmt.Println("From Store Room Repository")
//	fmt.Print(r)
//	_, err := rri.conn.Exec("INSERT INTO rooms (room_number,price,description,image,rmtype) values($1, $2, $3,$4,$5)",
//		r.RoomNumber, r.Price, r.Description, r.Image, r.Type)
//	if err != nil {
//		fmt.Println(err," From Store Room ")
//		return errors.Event("Insertion has failed")
//	}
//
//	return nil
//}