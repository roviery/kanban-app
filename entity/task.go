package entity

import "time"

type Task struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	Title       string    `json:"title" gorm:"type:varchar(255);not null"`
	Description string    `json:"description" gorm:"type:text;not null"`
	CategoryID  int       `json:"category_id" gorm:"type:int;not null"`
	UserID      int       `json:"user_id" gorm:"type:int;not null"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

type TaskRequest struct {
	ID          int    `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	CategoryID  int    `json:"category_id"`
}

type TaskCategoryRequest struct {
	ID         int `json:"id"`
	CategoryID int `json:"category_id" binding:"required"`
}

type TaskResponse struct {
	UserID  int    `json:"user_id"`
	TaskID  int    `json:"task_id"`
	Message string `json:"message"`
}

func NewTaskResponse(userId int, taskId int, msg string) TaskResponse {
	return TaskResponse{
		UserID:  userId,
		TaskID:  taskId,
		Message: msg,
	}
}
