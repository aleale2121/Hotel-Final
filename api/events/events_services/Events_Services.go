package events_services

import (
	"fmt"
	"github.com/yuidegm/Hotel-Rental-Managemnet-System/api/entity"
	"github.com/yuidegm/Hotel-Rental-Managemnet-System/api/events"
)

// CommentService implements menu.CommentService interface
type EventsService struct {
	EventsRepo events.EventRepository
}
// NewCommentService returns a new CommentService object
func NewEventService(commRepo events.EventRepository) events.EventsService {
	return &EventsService{EventsRepo: commRepo}
}

// Comments returns all stored comments
func (cs *EventsService) Events() ([]entity.Events, []error) {
	cmnts, errs := cs.EventsRepo.Events()
	fmt.Println("ents gorm serv",cmnts,errs)

	if len(errs)>0 {
		fmt.Println("ents gorm serv",cmnts,errs)
		return nil, errs
	}
	return cmnts, nil
}
// Comments retrieves stored comment by its id
func (cs *EventsService) EventById(id uint) (*entity.Events, []error) {
	cmnt, errs := cs.EventsRepo.EventById(id)
	if len(errs)>0 {
		return nil, errs
	}
	return cmnt, nil
}
// UpdateComment updates a given comment
func (cs *EventsService) UpdateEvent(comment *entity.Events) (*entity.Events, []error) {
	cmnt, errs := cs.EventsRepo.UpdateEvent(comment)
	if errs!=nil {
		return nil, errs
	}
	return cmnt, errs
}
// DeleteEvents deletes a given comment
func (cs *EventsService) DeleteEvent(id uint) (*entity.Events, []error) {
	cmnt, errs := cs.EventsRepo.DeleteEvent(id)
	if errs!=nil {
		return nil, errs
	}
	return cmnt, errs
}
// StoreComment stores a given comment
func (cs *EventsService) StoreEvent(comment *entity.Events) (*entity.Events, []error) {
	cmnt, errs := cs.EventsRepo.StoreEvent(comment)
	if errs!=nil {
		return nil, errs
	}
	return cmnt, errs
}
