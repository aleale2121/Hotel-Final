package events_repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/aleale2121/Hotel-Final/api/entity"
	"github.com/aleale2121/Hotel-Final/api/events"
)

// CommentGormRepo implements menu.CommentRepository interface
type EventGormRepo struct {
	conn *gorm.DB
}

// NewCommentGormRepo returns new object of CommentGormRepo
func NewEventsGormRepo(db *gorm.DB) events.EventRepository {
	return &EventGormRepo{conn: db}
}

// Comments returns all customer comments stored in the database
func (cmntRepo *EventGormRepo) Events() ([]entity.Events, []error) {
	cmnts := []entity.Events{}
	fmt.Println("ents gormrepo")

	errs := cmntRepo.conn.Find(&cmnts).GetErrors()
	fmt.Println(cmnts)
	if len(errs)>0 {
		return nil, errs
	}
	return cmnts, nil
}
// Comments retrieves a customer comment from the database by its id
func (cmntRepo *EventGormRepo) EventById(id uint) (*entity.Events, []error) {
	cmnt := &entity.Events{}
	errs := cmntRepo.conn.First(cmnt, id).GetErrors()
	if len(errs)>0{
		fmt.Println("inside evid errr")
		return nil, errs
	}
	return cmnt, nil
}

// UpdateComment updates a given customer comment in the database
func (cmntRepo *EventGormRepo) UpdateEvent(event *entity.Events) (*entity.Events, []error) {
	cmnt := event
	errs := cmntRepo.conn.Save(cmnt).GetErrors()
	if errs!=nil {
		return nil, errs
	}
	return cmnt, errs
}

// DeleteEvents deletes a given customer comment from the database
func (cmntRepo *EventGormRepo) DeleteEvent(id uint) (*entity.Events, []error) {
	cmnt, errs := cmntRepo.EventById(id)

	if errs!=nil {
		return nil, errs
	}
	errs = cmntRepo.conn.Delete(cmnt, id).GetErrors()
	if errs!=nil {
		return nil, errs
	}
	return cmnt, errs
}

// StoreComment stores a given customer comment in the database
func (cmntRepo *EventGormRepo) StoreEvent(event *entity.Events) (*entity.Events, []error) {
	cmnt := event
	errs := cmntRepo.conn.Create(cmnt).GetErrors()
	if errs!=nil {
		return nil,errs
	}
	return cmnt, errs
}

