// Domain Service
// DIコンテナ
// infrastructure/repository/repository.goから依存性逆転の原則が適用される
package repository

import (
	"go-onion-arch-sample/ent"
)

type TaskRepository interface {
	CreateTask(task ent.Task) (*ent.Task, error)
	GetTaskById(taskID int) (*ent.Task, error)
	GetTasks(profileID string) ([]*ent.Task, error)
	UpdateTask(task ent.Task, taskID int) (*ent.Task, error)
	DeleteTask(taskID int) error
}
