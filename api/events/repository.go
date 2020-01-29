package events

import (
	"github.com/yuidegm/Hotel-Rental-Managemnet-System/api/entity"
)

// CategoryService specifies food menu category news_services
type EventRepository interface {
	Events() ([]entity.Events, []error)
	EventById(id uint) (*entity.Events, []error)
	UpdateEvent(comment *entity.Events) (*entity.Events, []error)
	DeleteEvent(id uint) (*entity.Events,[]error)
	StoreEvent(comment *entity.Events) (*entity.Events, []error)
}
