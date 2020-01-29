package events_repository

import (


"fmt"
"github.com/yuidegm/Hotel-Rental-Managemnet-System/entity"
"github.com/jinzhu/gorm"
)

// RoomRepositoryImpl implements the rooms.RoomRepository interface
type GormEventsRepositoryImpl struct {
	Con *gorm.DB
}

// NewRoomRepositoryImpl will create an object of PsqlRoomRepository
func NewGormEventsRepositoryImpl(Conn *gorm.DB) *GormEventsRepositoryImpl {
	return &GormEventsRepositoryImpl{Con: Conn}
}

// News returns all rooms from the database
func (rri *GormEventsRepositoryImpl) Events() ([]entity.Events, []error) {
	cmnts := []entity.Events{}
	fmt.Println("ents gormrepo")
	errs := rri.Con.Find(&cmnts).GetErrors()
	fmt.Println(cmnts)
	if len(errs)>0 {
		return nil, nil
	}
	return cmnts, nil
}

// NewsById returns a News with a given id
func (rri *GormEventsRepositoryImpl) EventsById(id int) (*entity.Events,[]error) {
	cmnt := &entity.Events{}
	errs := rri.Con.First(cmnt, id).GetErrors()
	if len(errs)>0{
		fmt.Println("inside evid errr")
		return  nil,errs
	}
	return cmnt, nil
}

// UpdateNews updates a given object with a new data
func (rri *GormEventsRepositoryImpl) UpdateEvents(r entity.Events) (*entity.Events,[]error) {
	cmnt := r
	errs := rri.Con.Save(&cmnt).GetErrors()
	if errs!=nil {
		return nil, errs
	}
	return nil, nil
}

// DeleteNews removes a News from a database by its id
func (rri *GormEventsRepositoryImpl) DeleteEvents(id int) (*entity.Events, []error) {
	cmnt, errs := rri.EventsById(id)

	if errs!=nil {
		return nil, errs
	}
	errs = rri.Con.Delete(&cmnt, id).GetErrors()
	if errs!=nil {
		return nil, errs
	}
	return cmnt, errs
}

// StoreNews stores new News information to database
func (rri *GormEventsRepositoryImpl) StoreEvents(r entity.Events)  (*entity.Events, []error) {
	cmnt := r
	errs := rri.Con.Create(&cmnt).GetErrors()
	if errs!=nil {
		return nil,errs
	}
	return nil, errs
}


