package events

import "github.com/yuidegm/Hotel-Rental-Managemnet-System/api/entity"

// NewsService specifies News menu News news_services
type EventsService interface {
	Events() ([]entity.Events, []error)
	EventById(id uint) (*entity.Events, []error)
	UpdateEvent(comment *entity.Events) (*entity.Events, []error)
	DeleteEvent(id uint) (*entity.Events,[]error)
	StoreEvent(comment *entity.Events) (*entity.Events, []error)
}
