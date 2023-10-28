package router

import (
	"go-onion-arch-sample/userinterface/handler"

	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo, taskHandler *handler.TaskHandler, profileHandler *handler.ProfileHandler) {

	// e.POST("/tasks", taskHandler.CreateTask)
	e.GET("/tasks/:id", taskHandler.GetTaskByID, MiddlewareWithID)
	e.GET("/tasks", taskHandler.GetTasks)
	e.POST("/tasks/:id", taskHandler.UpdateTask, MiddlewareWithID)
	e.DELETE("/tasks/:id", taskHandler.DeleteTask, MiddlewareWithID)

	e.POST("/profiles", profileHandler.CreateProfile)
	e.GET("/profiles/:id", profileHandler.GetProfileById)
	// e.GET("/profiles", profileHandler.GetProfiles)
	e.POST("/profiles/:id", profileHandler.UpdateProfile)
	// e.DELETE("/profiles/:id", profileHandler.DeleteProfile)
}

func MiddlewareWithID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		c.Set("id", id)
		return next(c)
	}
}
