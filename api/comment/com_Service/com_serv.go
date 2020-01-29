package com_Service
import (
	"github.com/aleale2121/Hotel-Final/comment"
	"github.com/aleale2121/Hotel-Final/api/entity"
)

// RoomServiceImpl implements rooms.RoomService interface
type GormComServiceImpl struct {
	newsRepo comment.CommentRepository
}



// NewNewsServiceImpl will create new RoomService object
func NewGormComServiceImpl(NewRepo comment.CommentRepository) comment.CommentServices {
	return &GormComServiceImpl{newsRepo: NewRepo}
}

// News returns list of all rooms
func (rs *GormComServiceImpl) Comments() ([]entity.Comments, []error) {

	news, err := rs.newsRepo.Comment()

	if err != nil {
		return nil, err
	}

	return news, nil
}

// StoreNews persists new room information
func (rs *GormComServiceImpl) StoreCom(neww entity.Comments) (*entity.Comments, []error) {

	r,err := rs.newsRepo.StoreCom(neww)

	if err != nil {
		return nil,err
	}

	return r,nil
}

// NewById returns a room object with a given id
func (rs *GormComServiceImpl)CommentsById(id int) (*entity.Comments, []error) {

	r, err := rs.newsRepo.CommentsById(int(id))

	if err != nil {
		return r, err
	}

	return r, nil
}

// UpdateNews updates a cateogory with new data
//func (rs *GormComServiceImpl) UpdateEvents(neww entity.Comments) (*entity.Comments, []error){
//
//	r,err := rs.newsRepo.UpdateEvents(neww)
//
//	if err != nil {
//		return r,err
//	}
//
//	return r,nil
//}

// DeleteNews delete a room by its id
func (rs *GormComServiceImpl) DeleteCom(id int) (*entity.Comments, []error) {

	r,err := rs.newsRepo.DeleteCom(int(id))
	if err != nil {
		return r,err
	}
	return r,nil
}



