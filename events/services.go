package events

import "github.com/aleale2121/Hotel-Final/entity"

// NewsService specifies News menu News news_services
type EventsService interface {
	Events() ([]entity.Events, []error)
	EventsById(id int) (*entity.Events, []error)
	UpdateEvents(news entity.Events) (*entity.Events, []error)
	DeleteEvents(id int) (*entity.Events, []error)
	StoreEvents(news entity.Events) (*entity.Events, []error)
}