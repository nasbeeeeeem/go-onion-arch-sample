package main

import (
	"go-onion-arch-sample/application/service"
	"go-onion-arch-sample/infrastructure/database"
	"go-onion-arch-sample/infrastructure/repository"
	"go-onion-arch-sample/userinterface/handler"
	"go-onion-arch-sample/userinterface/router"

	"github.com/labstack/echo/v4"
)

func main() {
	db, err := database.NewDBClient()
	if err != nil {
		panic(err)
	}

	taskRepo := repository.NewTaskRepository(db)

	taskService := service.NewTaskService(taskRepo)

	taskHandler := handler.NewTaskHandler(taskService)

	profileRepo := repository.NewProfileRepository(db)

	profileService := service.NewProfileService(profileRepo)

	profileHandler := handler.NewProfileHandler(profileService)

	e := echo.New()

	router.NewRouter(e, taskHandler, profileHandler)

	e.Start("localhost:8080")
}
