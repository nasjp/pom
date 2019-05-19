package repository

import "github.com/YukihiroTaniguchi/pom/domain/model"

// TimerSetRepository ...
type TimerSetRepository interface {
	Create(*model.TimerSet, string) error
	Get(string) (*model.TimerSet, error)
	Update(*model.TimerSet, string) error
	Delete(string) error
	Find() ([]string, error)
}
