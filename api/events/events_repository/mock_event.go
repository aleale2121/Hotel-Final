package events_repository

import (
	"errors"

	"github.com/aleale2121/Hotel-Final/api/entity"
	"github.com/aleale2121/Hotel-Final/api/events"
	"github.com/jinzhu/gorm"
)

// MockEventsRepo implements the menu.CategoryRepository interface
type MockEventsRepo struct {
	conn *gorm.DB
}

// NewMockEventsRepo will create a new object of MockEventsRepo
func NewMockEventsRepo(db *gorm.DB) events.EventRepository {
	return &MockEventsRepo{conn: db}
}
// Categories returns all fake categories
func (mCatRepo *MockEventsRepo) Events() ([]entity.Events, []error) {
	ctgs := []entity.Events{entity.EventsMock}
	return ctgs, nil
}

func (mCatRepo *MockEventsRepo) EventById(id uint) (*entity.Events, []error) {
	ctg := entity.EventsMock
	if id == 1 {
		return &ctg, nil
	}
	return nil, []error{errors.New("Not found")}
}

func (mCatRepo *MockEventsRepo) UpdateEvent(news *entity.Events) (*entity.Events, []error) {
	cat := entity.EventsMock
	return &cat, nil
}

func (mCatRepo *MockEventsRepo) DeleteEvent(id uint) (*entity.Events, []error) {
	cat := entity.EventsMock
	if id != 1 {
		return nil, []error{errors.New("Not found")}
	}
	return &cat, nil
}

func (mCatRepo *MockEventsRepo) StoreEvent(news *entity.Events) (*entity.Events, []error) {
	//cat := newss
	return nil, nil
}


