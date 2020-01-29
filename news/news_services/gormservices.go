package news_services

import (
	"github.com/aleale2121/Hotel-Final/entity"
	"github.com/aleale2121/Hotel-Final/news"
)

// RoomServiceImpl implements rooms.RoomService interface
type GormNewsServiceImpl struct {
	newsRepo news.NewsRepository
}

// NewNewsServiceImpl will create new RoomService object
func NewGormNewsServiceImpl(NewRepo news.NewsRepository) *GormNewsServiceImpl {
	return &GormNewsServiceImpl{newsRepo: NewRepo}
}

// News returns list of all rooms
func (rs *GormNewsServiceImpl) News() ([]entity.News, []error) {

	news, err := rs.newsRepo.News()

	if err != nil {
		return nil, err
	}

	return news, nil
}
func (rri *GormNewsServiceImpl) NewsFive() ([]entity.News, []error) {
         news,errs:=rri.newsRepo.NewsFive()
	if len(errs)>0 {
		return nil, nil
	}
	return news, nil
}
// StoreNews persists new room information
func (rs *GormNewsServiceImpl) StoreNews(neww entity.News) (*entity.News, []error) {

	r,err := rs.newsRepo.StoreNews(neww)

	if err != nil {
		return nil,err
	}

	return r,nil
}

// NewById returns a room object with a given id
func (rs *GormNewsServiceImpl)NewsById(id int) (*entity.News, []error) {

	r, err := rs.newsRepo.NewsById(int(id))

	if err != nil {
		return r, err
	}

	return r, nil
}

// UpdateNews updates a cateogory with new data
func (rs *GormNewsServiceImpl) UpdateNews(neww entity.News) (*entity.News, []error){

	r,err := rs.newsRepo.UpdateNews(neww)

	if err != nil {
		return r,err
	}

	return r,nil
}

// DeleteNews delete a room by its id
func (rs *GormNewsServiceImpl) DeleteNews(id int) (*entity.News, []error) {

	r,err := rs.newsRepo.DeleteNews(int(id))
	if err != nil {
		return r,err
	}
	return r,nil
}


