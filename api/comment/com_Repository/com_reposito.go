package com_Repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/yuidegm/Hotel-Rental-Managemnet-System/api/comment"
	"github.com/yuidegm/Hotel-Rental-Managemnet-System/api/entity"
)

// RoomRepositoryImpl implements the rooms.RoomRepository interface
type GormComRepositoryImpl struct {
	Con *gorm.DB
}



// NewRoomRepositoryImpl will create an object of PsqlRoomRepository
func NewGormComRepositoryImpl(Conn *gorm.DB) comment.CommentRepository{
	return &GormComRepositoryImpl{Con: Conn}
}

// News returns all rooms from the database
func (rri *GormComRepositoryImpl) Comment() ([]entity.Comments, []error) {
	cmnts := []entity.Comments{}
	fmt.Println("ents gormrepo")
	errs := rri.Con.Find(&cmnts).GetErrors()
	fmt.Println(cmnts)
	if len(errs)>0 {
		return nil, nil
	}
	return cmnts, nil
}

// NewsById returns a News with a given id
func (rri *GormComRepositoryImpl) CommentsById(id int) (*entity.Comments, []error) {
	cmnt := &entity.Comments{}
	errs := rri.Con.First(cmnt, id).GetErrors()
	if len(errs)>0{
		fmt.Println("inside evid errr")
		return  nil,errs
	}
	return cmnt, nil
}

// DeleteNews removes a News from a database by its id
func (rri *GormComRepositoryImpl) DeleteCom(id int) (*entity.Comments, []error) {
	cmnt, errs := rri.CommentsById(id)

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
func (rri *GormComRepositoryImpl) StoreCom(r entity.Comments)  (*entity.Comments, []error) {
	cmnt := r
	errs := rri.Con.Create(&cmnt).GetErrors()
	if errs!=nil {
		return nil,errs
	}
	return nil, errs
}


