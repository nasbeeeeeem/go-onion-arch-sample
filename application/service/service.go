// Application Service
package service

import (
	"go-onion-arch-sample/domain/model"
	"go-onion-arch-sample/domain/repository"
)

type TaskService interface {
	CreateTask(task *model.Task) (*model.Task, error)
	GetTaskById(id string) (*model.Task, error)
	GetTasks() ([]model.Task, error)
	UpdateTask(task *model.Task, id string) (*model.Task, error)
	DeleteTask(id string) error
}

type taskService struct {
	taskRepo repository.TaskRepository
}

func NewTaskService(taskRepo repository.TaskRepository) TaskService {
	return &taskService{
		taskRepo: taskRepo,
	}
}

// タスクの登録
func (t *taskService) CreateTask(task *model.Task) (*model.Task, error) {
	return t.taskRepo.CreateTask(task)
}

// タスクの取得
func (t *taskService) GetTaskById(id string) (*model.Task, error) {
	return t.taskRepo.GetTaskById(id)
}

// タスクの一覧取得
func (t *taskService) GetTasks() ([]model.Task, error) {
	return t.taskRepo.GetTasks()
}

// タスクの更新
func (t *taskService) UpdateTask(task *model.Task, id string) (*model.Task, error) {
	return t.taskRepo.UpdateTask(task, id)
}

// タスクの削除
func (t *taskService) DeleteTask(id string) error {
	return t.taskRepo.DeleteTask(id)
}
