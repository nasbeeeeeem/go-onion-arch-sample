// Infrastructure
package repository

import (
	"go-onion-arch-sample/domain/model"
	"go-onion-arch-sample/domain/repository"
	"go-onion-arch-sample/infrastructure/database"
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
func (t *taskRepository) CreateTask(task *model.Task) (*model.Task, error) {
	result := t.dbClient.Client.Create(task)
	if result.Error != nil {
		return nil, result.Error
	}

	return task, nil
}

// タスクの取得
func (t *taskRepository) GetTaskById(id string) (*model.Task, error) {
	task := &model.Task{}
	result := t.dbClient.Client.Find(task, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return task, nil
}

// タスクの全件取得
func (t *taskRepository) GetTasks() ([]model.Task, error) {
	var tasks []model.Task
	result := t.dbClient.Client.Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}

	return tasks, nil
}

// タスクの更新
func (t *taskRepository) UpdateTask(task *model.Task, id string) (*model.Task, error) {
	result := t.dbClient.Client.Model(&task).Where("ID = ?", id).Updates(task)
	if result.Error != nil {
		return nil, result.Error
	}

	return task, nil
}

// タスクの削除
func (t *taskRepository) DeleteTask(id string) error {
	result := t.dbClient.Client.Delete(&model.Task{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
