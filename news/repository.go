package news

import "github.com/yuidegm/Hotel-Rental-Managemnet-System/entity"

// CategoryService specifies food menu category news_services
type NewsRepository interface {
	News() ([]entity.News, []error)
	NewsById(id int) (*entity.News, []error)
	UpdateNews(news entity.News) (*entity.News, []error)
	DeleteNews(id int) (*entity.News, []error)
	StoreNews(news entity.News) (*entity.News, []error)
	NewsFive() ([]entity.News, []error)

}

