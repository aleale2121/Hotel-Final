package news_services

import (
	"fmt"
	"github.com/aleale2121/Hotel-Final/api/entity"
	"github.com/aleale2121/Hotel-Final/api/news"
)

// CommentService implements menu.CommentService interface
type EventsService struct {
	EventsRepo news.NewsRepository
}

// NewCommentService returns a new CommentService object
func NewNewsService(commRepo news.NewsRepository) news.NewsService {
	return &EventsService{EventsRepo: commRepo}
}

// Comments returns all stored comments
func (cs *EventsService) News() ([]entity.News, []error) {
	cmnts, errs := cs.EventsRepo.News()
	fmt.Println("ents gorm serv",cmnts,errs)

	if len(errs)>0 {
		fmt.Println("ents gorm serv",cmnts,errs)
		return nil, errs
	}
	return cmnts, nil
}
// Comments retrieves stored comment by its id
func (cs *EventsService) NewsById(id uint) (*entity.News, []error) {
	cmnt, errs := cs.EventsRepo.NewsById(id)
	if len(errs)>0 {
		return nil, errs
	}
	return cmnt, nil
}
// UpdateComment updates a given comment
func (cs *EventsService) UpdateNews(comment *entity.News) (*entity.News, []error) {
	cmnt, errs := cs.EventsRepo.UpdateNews(comment)
	if errs!=nil {
		return nil, errs
	}
	return cmnt, errs
}
// DeleteEvents deletes a given comment
func (cs *EventsService) DeleteNews(id uint) (*entity.News, []error) {
	cmnt, errs := cs.EventsRepo.DeleteNews(id)
	if errs!=nil {
		return nil, errs
	}
	return cmnt, errs
}
// StoreComment stores a given comment
func (cs *EventsService) StoreNews(comment *entity.News) (*entity.News, []error) {
	cmnt, errs := cs.EventsRepo.StoreNews(comment)
	if errs!=nil {
		return nil, errs
	}
	return cmnt, errs
}
