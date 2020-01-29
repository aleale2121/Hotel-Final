package news_repository

import (
	"errors"
	"github.com/aleale2121/Hotel-Final/api/news"

	"github.com/jinzhu/gorm"
	"github.com/aleale2121/Hotel-Final/api/entity"
)

// MockEventsRepo implements the menu.CategoryRepository interface
type MockNewsRepo struct {
	conn *gorm.DB
}

// NewMockEventsRepo will create a new object of MockEventsRepo
func NewMockNewsRepo(db *gorm.DB) news.NewsRepository {
	return &MockNewsRepo{conn: db}
}
// Categories returns all fake categories
func (mCatRepo *MockNewsRepo) News() ([]entity.News, []error) {
	ctgs := []entity.News{entity.NewsMock}
	return ctgs, nil
}

func (mCatRepo *MockNewsRepo) NewsById(id uint) (*entity.News, []error) {
	ctg := entity.NewsMock
	if id == 1 {
		return &ctg, nil
	}
	return nil, []error{errors.New("Not found")}
}

func (mCatRepo *MockNewsRepo) UpdateNews(news *entity.News) (*entity.News, []error) {
	cat := entity.NewsMock
	return &cat, nil
}

func (mCatRepo *MockNewsRepo) DeleteNews(id uint) (*entity.News, []error) {
	cat := entity.NewsMock
	if id != 1 {
		return nil, []error{errors.New("Not found")}
	}
	return &cat, nil
}

func (mCatRepo *MockNewsRepo) StoreNews(news *entity.News) (*entity.News, []error) {
	//cat := newss
	return nil, nil
}


