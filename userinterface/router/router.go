package router

import (
	"go-onion-arch-sample/userinterface/handler"

	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo, taskHandler *handler.TaskHandler) {
	e.POST("/tasks", taskHandler.CreateTask)
	e.GET("/tasks/:id", taskHandler.GetTaskByID, MiddlewareWithID)
	e.GET("/tasks", taskHandler.GetTasks)
	e.PUT("/tasks/:id", taskHandler.UpdateTask, MiddlewareWithID)
	e.DELETE("/tasks/:id", taskHandler.DeleteTask, MiddlewareWithID)
}

func MiddlewareWithID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		c.Set("id", id)
		return next(c)
	}
}
