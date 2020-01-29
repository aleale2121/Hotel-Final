package events_services


import (
	"github.com/aleale2121/Hotel-Final/entity"
	"github.com/aleale2121/Hotel-Final/events"
)

// RoomServiceImpl implements rooms.RoomService interface
type GormEventsServiceImpl struct {
	newsRepo events.EventRepository
}

// NewNewsServiceImpl will create new RoomService object
func NewGormNewsServiceImpl(NewRepo events.EventRepository) *GormEventsServiceImpl {
	return &GormEventsServiceImpl{newsRepo: NewRepo}
}

// News returns list of all rooms
func (rs *GormEventsServiceImpl) Events() ([]entity.Events, []error) {

	news, err := rs.newsRepo.Events()

	if err != nil {
		return nil, err
	}

	return news, nil
}

// StoreNews persists new room information
func (rs *GormEventsServiceImpl) StoreEvents(neww entity.Events) (*entity.Events, []error) {

	r,err := rs.newsRepo.StoreEvents(neww)

	if err != nil {
		return nil,err
	}

	return r,nil
}

// NewById returns a room object with a given id
func (rs *GormEventsServiceImpl)EventsById(id int) (*entity.Events, []error) {

	r, err := rs.newsRepo.EventsById(int(id))

	if err != nil {
		return r, err
	}

	return r, nil
}

// UpdateNews updates a cateogory with new data
func (rs *GormEventsServiceImpl) UpdateEvents(neww entity.Events) (*entity.Events, []error){

	r,err := rs.newsRepo.UpdateEvents(neww)

	if err != nil {
		return r,err
	}

	return r,nil
}

// DeleteNews delete a room by its id
func (rs *GormEventsServiceImpl) DeleteEvents(id int) (*entity.Events, []error) {

	r,err := rs.newsRepo.DeleteEvents(int(id))
	if err != nil {
		return r,err
	}
	return r,nil
}



