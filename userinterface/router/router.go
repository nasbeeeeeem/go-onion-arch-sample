package router

import (
	"go-onion-arch-sample/userinterface/handler"

	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo, taskHandler *handler.TaskHandler, profileHandler *handler.ProfileHandler) {

	e.POST("/tasks", taskHandler.CreateTask)
	e.GET("/tasks/:id", taskHandler.GetTaskByID, MiddlewareWithID)
	e.GET("/tasks", taskHandler.GetTasks)
	e.PUT("/tasks/:id", taskHandler.UpdateTask, MiddlewareWithID)
	e.DELETE("/tasks/:id", taskHandler.DeleteTask, MiddlewareWithID)

	e.POST("/profiles", profileHandler.CreateProfile)
	e.GET("/profiles/:id", profileHandler.GetProfileById)
	e.GET("/profiles", profileHandler.GetProfiles)
	e.PUT("/profiles/:id", profileHandler.UpdateProfile)
	e.DELETE("/profiles/:id", profileHandler.DeleteProfile)
	// task := e.Group("/tasks")
	// task.POST("/", taskHandler.CreateTask)
	// task.GET("/:id", taskHandler.GetTaskByID, MiddlewareWithID)
	// task.GET("/", taskHandler.GetTasks)
	// task.PUT("/:id", taskHandler.UpdateTask, MiddlewareWithID)
	// task.DELETE("/:id", taskHandler.DeleteTask, MiddlewareWithID)

	// profile := e.Group("/profiles")
	// profile.POST("/", profileHandler.CreateProfile)
	// profile.GET("/:id", profileHandler.GetProfileById)
	// profile.GET("/", profileHandler.GetProfiles)
	// profile.PUT("/:id", profileHandler.UpdateProfile)
	// profile.DELETE("/:id", profileHandler.DeleteProfile)
}

func MiddlewareWithID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		c.Set("id", id)
		return next(c)
	}
}
