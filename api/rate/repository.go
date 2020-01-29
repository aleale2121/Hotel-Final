package rate
import "github.com/aleale2121/Hotel-Final/api/entity"




type RateRepository interface {
	AddRate(updaterate *entity.Rating) (*entity.Rating, []error)
	//api
	Rate() ([]entity.Rating, []error)
	Rates(id uint) (*entity.Rating, []error)
	UpdateRates(rpay *entity.Rating) (*entity.Rating, []error)
	DeleteRates(id uint) (*entity.Rating, []error)
}