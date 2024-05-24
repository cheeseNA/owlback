package repository

import (
	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ITaskRepository interface {
	AddTask(task Task) error
	GetTaskById(id uuid.UUID) (Task, error)
}

type TaskRepository struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewTaskRepository(db *gorm.DB, logger *zap.Logger) *TaskRepository {
	return &TaskRepository{db: db, logger: logger}
}

func (r *TaskRepository) AddTask(task Task) error {
	r.logger.Info("Adding task", zap.Any("task", task))
	return r.db.Create(&task).Error
}

func (r *TaskRepository) GetTaskById(id uuid.UUID) (Task, error) {
	var task Task
	err := r.db.First(&task, id).Error
	return task, err
}
