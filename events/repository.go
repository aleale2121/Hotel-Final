package events

import "github.com/aleale2121/Hotel-Final/entity"

// CategoryService specifies food menu category news_services
type EventRepository interface {
	Events() ([]entity.Events, []error)
	EventsById(id int) (*entity.Events, []error)
	UpdateEvents(news entity.Events) (*entity.Events, []error)
	DeleteEvents(id int) (*entity.Events, []error)
	StoreEvents(news entity.Events) (*entity.Events, []error)
}