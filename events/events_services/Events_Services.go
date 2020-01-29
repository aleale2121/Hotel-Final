package events_services
//
//import (
//	"github.com/getme04/original/entity"
//	"github.com/getme04/original/events"
//)
//
//// RoomServiceImpl implements rooms.RoomService interface
//type EventsServiceImpl struct {
//	eventsServ events.EventsService
//}
//
//// NewNewsServiceImpl will create new RoomService object
//func NewEventsServiceImpl(eventRepo events.EventsService) *EventsServiceImpl {
//	return &EventsServiceImpl{eventsServ: eventRepo}
//}
//
//// News returns list of all rooms
//func (rs *EventsServiceImpl) Events() ([]entity.Events, error) {
//
//	newss, err := rs.eventsServ.Events()
//
//	if err != nil {
//		return nil, err
//	}
//	return newss, nil
//}
//
//// StoreNews persists new room information
//func (rs *EventsServiceImpl) StoreEvent(event entity.Events) error {
//
//	err := rs.eventsServ.StoreEvent(event)
//
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//// NewById returns a room object with a given id
//func (rs *EventsServiceImpl) EventById(id int) (entity.Events, error) {
//
//	r, err := rs.eventsServ.EventById(id)
//
//	if err != nil {
//		return r, err
//	}
//
//	return r, nil
//}
//
//// UpdateNews updates a cateogory with new data
//func (rs *EventsServiceImpl) UpdateEvent(event entity.Events) error {
//
//	err := rs.eventsServ.UpdateEvent(event)
//
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//// DeleteNews delete a room by its id
//func (rs *EventsServiceImpl) DeleteEvent(id int) error {
//
//	err := rs.eventsServ.DeleteEvent(id)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//
//
