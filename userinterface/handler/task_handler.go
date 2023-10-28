package handler

import (
	"go-onion-arch-sample/application/service"
	"go-onion-arch-sample/ent"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	TaskService service.TaskService
}

// TaskServiceを実装したコンストラクター
func NewTaskHandler(taskService service.TaskService) *TaskHandler {
	return &TaskHandler{
		TaskService: taskService,
	}
}

// タスクの登録
func (h *TaskHandler) CreateTask(c echo.Context) error {
	task := ent.Task{}
	if err := c.Bind(&task); err != nil {
		return err
	}

	createdTask, err := h.TaskService.CreateTask(task)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, createdTask)
}

// タスクの取得
func (h *TaskHandler) GetTaskByID(c echo.Context) error {
	id := c.Param("id")
	taskID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	task, err := h.TaskService.GetTaskById(taskID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, task)
}

// タスクの全件取得
func (h *TaskHandler) GetTasks(c echo.Context) error {
	req := c.Request()
	authorization := req.Header.Get("Authorization")
	arr := strings.Split(authorization, " ")

	tasks, err := h.TaskService.GetTasks(arr[1])
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, tasks)
}

// タスクの更新
func (h *TaskHandler) UpdateTask(c echo.Context) error {
	id := c.Param("id")
	taskID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	var task ent.Task
	if err := c.Bind(&task); err != nil {
		return err
	}

	updateTask, err := h.TaskService.UpdateTask(task, taskID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, updateTask)
}

// タスクの削除
func (h *TaskHandler) DeleteTask(c echo.Context) error {
	id := c.Param("id")
	taskID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	if err := h.TaskService.DeleteTask(taskID); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
