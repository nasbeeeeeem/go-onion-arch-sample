// Domain Service
// DIコンテナ
// infrastructure/repository/repository.goから依存性逆転の原則が適用される
package repository

import "go-onion-arch-sample/domain/model"

type TaskRepository interface {
	CreateTask(task *model.Task) (*model.Task, error)
	GetTaskById(id string) (*model.Task, error)
	GetTasks() ([]model.Task, error)
	UpdateTask(task *model.Task, id string) (*model.Task, error)
	DeleteTask(id string) error
}
