package news_repository


import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/yuidegm/Hotel-Rental-Managemnet-System/entity"
)

// RoomRepositoryImpl implements the rooms.RoomRepository interface
type GormNewsRepositoryImpl struct {
	Con *gorm.DB
}

// NewRoomRepositoryImpl will create an object of PsqlRoomRepository
func NewGormNewsRepositoryImpl(Conn *gorm.DB) *GormNewsRepositoryImpl {
	return &GormNewsRepositoryImpl{Con: Conn}
}
// News returns all rooms from the database
func (rri *GormNewsRepositoryImpl) News() ([]entity.News, []error) {
	cmnts := []entity.News{}
	fmt.Println("ents gormrepo")
	errs := rri.Con.Find(&cmnts).GetErrors()
	fmt.Println(cmnts)
	if len(errs)>0 {
		return nil, nil
	}
	return cmnts, nil
}
func (rri *GormNewsRepositoryImpl) NewsFive() ([]entity.News, []error) {
	cmnts := []entity.News{}
	fmt.Println("ents gormrepo")
	errs := rri.Con.Order("id desc").Limit(5).Find(&cmnts).GetErrors()
	fmt.Println(cmnts)
	if len(errs)>0 {
		return nil, nil
	}
	return cmnts, nil
}
// NewsById returns a News with a given id
func (rri *GormNewsRepositoryImpl) NewsById(id int) (*entity.News,[]error) {
	cmnt := &entity.News{}
	errs := rri.Con.First(cmnt, id).GetErrors()
	if len(errs)>0{
		fmt.Println("inside evid errr")
		return  nil,errs
	}
	return cmnt, nil
}
// UpdateNews updates a given object with a new data
func (rri *GormNewsRepositoryImpl) UpdateNews(r entity.News) (*entity.News,[]error) {
	cmnt := r
	errs := rri.Con.Save(&cmnt).GetErrors()
	if errs!=nil {
		return nil, errs
	}
	return nil, nil
}

// DeleteNews removes a News from a database by its id
func (rri *GormNewsRepositoryImpl) DeleteNews(id int) (*entity.News, []error) {
	cmnt, errs := rri.NewsById(id)

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
func (rri *GormNewsRepositoryImpl) StoreNews(r entity.News)  (*entity.News, []error) {
	cmnt := r
	errs := rri.Con.Create(&cmnt).GetErrors()
	if errs!=nil {
		return nil,errs
	}
	return nil, errs
}

