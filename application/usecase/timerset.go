package usecase

import (
	"github.com/NasSilverBullet/pom/domain/model"
	"github.com/NasSilverBullet/pom/domain/repository"
)

// TimerSetUsecase ...
type TimerSetUsecase interface {
	Create(*model.TimerSet, string) error
	Get(string) (*model.TimerSet, error)
	Update(*model.TimerSet, string) error
	Delete(string) error
	Find() ([]string, error)
}

// NewTimerSetUsecase ...
func NewTimerSetUsecase(r repository.TimerSetRepository) TimerSetUsecase {
	return &timerSetUsecase{r}
}

type timerSetUsecase struct {
	repository.TimerSetRepository
}

func (u *timerSetUsecase) Create(t *model.TimerSet, fn string) (err error) {
	return u.TimerSetRepository.Create(t, fn)
}

func (u *timerSetUsecase) Get(fn string) (t *model.TimerSet, err error) {
	t, err = u.TimerSetRepository.Get(fn)
	return
}

func (u *timerSetUsecase) Update(t *model.TimerSet, fn string) (err error) {
	return u.TimerSetRepository.Update(t, fn)
}

func (u *timerSetUsecase) Delete(fn string) (err error) {
	return u.TimerSetRepository.Delete(fn)
}

func (u *timerSetUsecase) Find() (l []string, err error) {
	l, err = u.TimerSetRepository.Find()
	return
}
