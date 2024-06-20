package repository

import (
	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

type ITaskRepository interface {
	CreateTask(Task) error
	DeleteTaskByID(uuid.UUID) error
	GetTaskByID(uuid.UUID) (Task, error)
	UpdateTask(Task, Task) error
	GetTasks() ([]Task, error)
	GetTasksToCrawl() ([]Task, error)
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
	return r.db.Omit("ID", "LastCrawledAt", "LastContent").Create(&task).Error
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

func (r *TaskRepository) UpdateTask(targetTask Task, updateTask Task) error {
	r.logger.Info("Updating task", zap.Any("id", targetTask.ID))
	return r.db.Model(&targetTask).Updates(updateTask).Error
}

func (r *TaskRepository) GetTasks() ([]Task, error) {
	r.logger.Info("Getting tasks")
	var tasks []Task
	err := r.db.Where("is_public = ?", true).Find(&tasks).Error
	return tasks, err
}

func (r *TaskRepository) GetTasksToCrawl() ([]Task, error) {
	r.logger.Info("Getting tasks to crawl")
	var tasks []Task
	currentTime := time.Now()
	err := r.db.Where("last_crawled_at + duration_day * interval '1 day' < ?", currentTime).Or("last_crawled_at IS NULL").Find(&tasks).Error
	return tasks, err
}
