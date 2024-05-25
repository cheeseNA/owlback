package repository

import (
	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ITaskRepository interface {
	CreateTask(task Task) error
	DeleteTaskByID(id uuid.UUID) error
	GetTaskByID(id uuid.UUID) (Task, error)
	GetTasks() ([]Task, error)
}

type TaskRepository struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewTaskRepository(db *gorm.DB, logger *zap.Logger) *TaskRepository {
	return &TaskRepository{db: db, logger: logger}
}

func (r *TaskRepository) CreateTask(task Task) error {
	r.logger.Info("Creating task", zap.Any("task", task))
	task.CreatedBy = uuid.New()
	return r.db.Omit("ID").Create(&task).Error
}

func (r *TaskRepository) DeleteTaskByID(id uuid.UUID) error {
	r.logger.Info("Deleting task", zap.Any("id", id))
	tx := r.db.Delete(&Task{}, id)
	if tx.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return tx.Error
}

func (r *TaskRepository) GetTaskByID(id uuid.UUID) (Task, error) {
	r.logger.Info("Getting task by id", zap.Any("id", id))
	var task Task
	err := r.db.First(&task, id).Error
	return task, err
}

func (r *TaskRepository) GetTasks() ([]Task, error) {
	r.logger.Info("Getting tasks")
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}
