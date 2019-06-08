package main

import (
	"os"

	"github.com/NasSilverBullet/pom/application/usecase"
	"github.com/NasSilverBullet/pom/infrastructure/file"
	"github.com/NasSilverBullet/pom/presentation/api/handler"
	"github.com/NasSilverBullet/pom/presentation/api/router"
)

func main() {
	r := file.NewTimerSetRepository()
	u := usecase.NewTimerSetUsecase(r)
	h := handler.NewTimerSetHandler(u)
	root, err := router.NewRouter(h)
	if err != nil {
		os.Exit(-1)
	}
	if err := root.Execute(); err != nil {
		os.Exit(-1)
	}
}
