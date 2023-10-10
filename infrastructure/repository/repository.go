// Infrastructure
package repository

import (
	"context"
	"go-onion-arch-sample/domain/repository"
	"go-onion-arch-sample/ent"
	"go-onion-arch-sample/ent/task"
	"go-onion-arch-sample/infrastructure/database"

	"github.com/labstack/echo/v4"
)

type taskRepository struct {
	dbClient *database.DBClient
}

func NewTaskRepository(db *database.DBClient) repository.TaskRepository {
	return &taskRepository{
		dbClient: db,
	}
}

// タスクの登録
func (t *taskRepository) CreateTask(task *ent.Task) (*ent.Task, error) {
	newTask, err := t.dbClient.Client.Task.Create().SetTitele(task.Titele).SetCompleted(task.Completed).Save(context.Background())
	if err != nil {
		return nil, err
	}

	return newTask, nil
}

// タスクの取得
func (t *taskRepository) GetTaskById(taskID int) (*ent.Task, error) {
	task, err := t.dbClient.Client.Task.Query().Where(task.ID(uint(taskID))).Only(context.Background())
	if err != nil {
		return nil, err
	}
	return task, nil
}

// タスクの全件取得
func (t *taskRepository) GetTasks(c echo.Context) ([]*ent.Task, error) {
	tasks, err := t.dbClient.Client.Task.Query().All(context.Background())
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

// タスクの更新
func (t *taskRepository) UpdateTask(task *ent.Task, taskID int) (*ent.Task, error) {
	updateTask, err := t.dbClient.Client.Task.UpdateOneID(uint(taskID)).SetTitele(task.Titele).SetCompleted(task.Completed).Save(context.Background())
	if err != nil {
		return nil, err
	}

	return updateTask, nil
}

// タスクの削除
func (t *taskRepository) DeleteTask(taskID int) error {
	err := t.dbClient.Client.Task.DeleteOneID(uint(taskID)).Exec(context.Background())
	if err != nil {
		return err
	}

	return nil
}
