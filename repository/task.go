package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"

	"gorm.io/gorm"
)

type TaskRepository interface {
	GetTasks(ctx context.Context, id int) ([]entity.Task, error)
	StoreTask(ctx context.Context, task *entity.Task) (taskId int, err error)
	GetTaskByID(ctx context.Context, id int) (entity.Task, error)
	GetTasksByCategoryID(ctx context.Context, catId int) ([]entity.Task, error)
	UpdateTask(ctx context.Context, task *entity.Task) error
	DeleteTask(ctx context.Context, id int) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db}
}

func (r *taskRepository) GetTasks(ctx context.Context, id int) ([]entity.Task, error) {
	results := []entity.Task{}
	rows, err := r.db.Table("tasks").Where("user_id = ?", id).Select("*").Rows()

	if err != nil {
		return []entity.Task{}, err
	}

	defer rows.Close()

	for rows.Next() {
		r.db.ScanRows(rows, &results)
	}

	return results, nil
}

func (r *taskRepository) StoreTask(ctx context.Context, task *entity.Task) (taskId int, err error) {
	result := r.db.Create(&task)

	if result.Error != nil {
		return 0, result.Error
	}

	return task.ID, nil
}

func (r *taskRepository) GetTaskByID(ctx context.Context, id int) (entity.Task, error) {
	result := entity.Task{}
	response := r.db.Where("id = ?", id).Model(&entity.Task{}).First(&result)

	if response.Error != nil {
		return entity.Task{}, response.Error
	}

	if response.RowsAffected == 0 {
		return entity.Task{}, nil
	}

	return result, nil // TODO: replace this
}

func (r *taskRepository) GetTasksByCategoryID(ctx context.Context, catId int) ([]entity.Task, error) {
	results := []entity.Task{}
	rows, err := r.db.Table("tasks").Where("category_id = ?", catId).Select("*").Rows()

	if err != nil {
		return []entity.Task{}, err
	}

	defer rows.Close()

	for rows.Next() {
		r.db.ScanRows(rows, &results)
	}

	return results, nil
}

func (r *taskRepository) UpdateTask(ctx context.Context, task *entity.Task) error {
	result := r.db.Model(&task).Updates(task)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *taskRepository) DeleteTask(ctx context.Context, id int) error {
	result := r.db.Delete(&entity.Task{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
