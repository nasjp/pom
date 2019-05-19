package main

import (
	"os"

	"github.com/YukihiroTaniguchi/pom/application/usecase"
	"github.com/YukihiroTaniguchi/pom/infrastructure/file"
	"github.com/YukihiroTaniguchi/pom/presentation/api/handler"
	"github.com/YukihiroTaniguchi/pom/presentation/api/router"
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
