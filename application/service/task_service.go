// Application Service
package service

import (
	"go-onion-arch-sample/domain/repository"
	"go-onion-arch-sample/ent"
)

type TaskService interface {
	CreateTask(task ent.Task) (*ent.Task, error)
	GetTaskById(taskID int) (*ent.Task, error)
	GetTasks(profileID string) ([]*ent.Task, error)
	UpdateTask(task ent.Task, taskID int) (*ent.Task, error)
	DeleteTask(taskID int) error
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
func (t *taskService) CreateTask(task ent.Task) (*ent.Task, error) {
	return t.taskRepo.CreateTask(task)
}

// タスクの取得
func (t *taskService) GetTaskById(taskID int) (*ent.Task, error) {
	return t.taskRepo.GetTaskById(taskID)
}

// タスクの全件取得
func (t *taskService) GetTasks(profileID string) ([]*ent.Task, error) {
	return t.taskRepo.GetTasks(profileID)
}

// タスクの更新
func (t *taskService) UpdateTask(task ent.Task, taskID int) (*ent.Task, error) {
	return t.taskRepo.UpdateTask(task, taskID)
}

// タスクの削除
func (t *taskService) DeleteTask(taskID int) error {
	return t.taskRepo.DeleteTask(taskID)
}
