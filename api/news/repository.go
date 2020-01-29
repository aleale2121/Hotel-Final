package news

import "github.com/yuidegm/Hotel-Rental-Managemnet-System/api/entity"

// CategoryService specifies food menu category news_services
type NewsRepository interface {
	News() ([]entity.News, []error)
	NewsById(id uint) (*entity.News, []error)
	UpdateNews(comment *entity.News) (*entity.News, []error)
	DeleteNews(id uint) (*entity.News,[]error)
	StoreNews(comment *entity.News) (*entity.News, []error)
}

