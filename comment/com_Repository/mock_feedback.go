package com_Repository

import (
	"errors"

	"github.com/aleale2121/Hotel-Final/entity"
	"github.com/aleale2121/Hotel-Final/comment"
	"github.com/jinzhu/gorm"
)

// MockComRepo implements the menu.CategoryRepository interface
type MockComRepo struct {
	conn *gorm.DB
}


// NewMockEventsRepo will create a new object of MockComRepo
func NewMockComRepo(db *gorm.DB) comment.CommentRepository{
	return &MockComRepo{conn: db}
}
// Categories returns all fake categories
func (mCatRepo *MockComRepo) Comment() ([]entity.Comments, []error) {
	ctgs := []entity.Comments{entity.CommentsMock}
	return ctgs, nil
}

func (mCatRepo *MockComRepo) CommentsById(id int) (*entity.Comments, []error) {
	ctg := entity.CommentsMock
	if id == 1 {
		return &ctg, nil
	}
	return nil, []error{errors.New("Not found")}
}


func (mCatRepo *MockComRepo) DeleteCom(id int) (*entity.Comments, []error) {
	cat := entity.CommentsMock
	if id != 1 {
		return nil, []error{errors.New("Not found")}
	}
	return &cat, nil
}

func (mCatRepo *MockComRepo) StoreCom(com entity.Comments) (*entity.Comments, []error) {
	cat := com
	return &cat, nil
}


