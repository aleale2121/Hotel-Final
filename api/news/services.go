package news

import "github.com/aleale2121/Hotel-Final/api/entity"

// NewsService specifies News menu News news_services
type NewsService interface {
	News() ([]entity.News, []error)
	NewsById(id uint) (*entity.News, []error)
	UpdateNews(comment *entity.News) (*entity.News, []error)
	DeleteNews(id uint) (*entity.News,[]error)
	StoreNews(comment *entity.News) (*entity.News, []error)
}
