package comment
import "github.com/aleale2121/Hotel-Final/entity"

// CategoryService specifies food menu category news_services
type CommentServices interface {
	Comment() ([]entity.Comments, []error)
	CommentsById(id int) (*entity.Comments, []error)
	//UpdateEvents(newss entity.Events) (*entity.Events, []error)
	DeleteCom(id int) (*entity.Comments, []error)
	StoreCom(news entity.Comments) (*entity.Comments, []error)
}
