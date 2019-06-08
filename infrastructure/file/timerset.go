package file

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/NasSilverBullet/pom/domain/model"
	"github.com/NasSilverBullet/pom/domain/repository"
)

type timerSetRepository struct {
}

// NewTimerSetRepository ...
func NewTimerSetRepository() repository.TimerSetRepository {
	return &timerSetRepository{}
}

// Create ...
func (r *timerSetRepository) Create(t *model.TimerSet, fn string) (err error) {
	if strings.Contains(fn, "/") || strings.Contains(fn, ".") {
		err = fmt.Errorf("\"%s\" is not permitted file name, please confirm", fn)
		return
	}

	ep, err := model.GetGOPATH()
	if err != nil {
		return
	}

	p := ep + model.GetAppDir() + fn
	j, err := json.Marshal(t)
	if err != nil {
		return
	}
	err = create(j, p)
	return
}

// Get ...
func (r *timerSetRepository) Get(fn string) (t *model.TimerSet, err error) {
	ep, err := model.GetGOPATH()
	if err != nil {
		return
	}
	p := ep + model.GetAppDir() + fn
	j, err := get(p)
	if err != nil {
		return
	}
	err = json.Unmarshal(j, &t)
	return
}

// Update ...
func (r *timerSetRepository) Update(t *model.TimerSet, fn string) (err error) {
	ep, err := model.GetGOPATH()
	if err != nil {
		return
	}
	p := ep + model.GetAppDir() + fn

	j, err := json.Marshal(t)
	if err != nil {
		return
	}
	err = update(j, p)
	return
}

// Delete ...
func (r *timerSetRepository) Delete(fn string) (err error) {
	ep, err := model.GetGOPATH()
	if err != nil {
		return
	}
	p := ep + model.GetAppDir() + fn
	err = remove(p)
	return
}

// Find ...
func (r *timerSetRepository) Find() (l []string, err error) {
	return
}
