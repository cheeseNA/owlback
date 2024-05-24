package service

import (
	"context"
	ogen "github.com/cheeseNA/owlback/ogen"
	"github.com/cheeseNA/owlback/repository"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/http"
	"net/url"
	"sync"
)

type Service struct {
	tasks map[uuid.UUID]ogen.Task
	repo  repository.ITaskRepository
	mux   sync.Mutex
}

func NewService(repo repository.ITaskRepository) *Service {
	return &Service{
		tasks: map[uuid.UUID]ogen.Task{},
		repo:  repo,
	}
}

func (s *Service) AddTask(ctx context.Context, task *ogen.Task) (*ogen.Task, error) {
	err := s.repo.AddTask(repository.Task{
		ID:           uuid.New(),
		SiteURL:      task.SiteURL.String(),
		Condition:    task.Condition,
		DurationDays: task.DurationDay,
	})
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (s *Service) GetTaskById(ctx context.Context, params ogen.GetTaskByIdParams) (ogen.GetTaskByIdRes, error) {
	task, err := s.repo.GetTaskById(params.TaskId)
	if err != nil {
		return &ogen.GetTaskByIdNotFound{}, err
	}
	retTask := ogen.Task{
		ID:          ogen.NewOptUUID(task.ID),
		SiteURL:     url.URL{Path: task.SiteURL},
		Condition:   task.Condition,
		DurationDay: task.DurationDays,
	}
	return &retTask, nil
}

func (s *Service) UpdateTask(ctx context.Context, params ogen.UpdateTaskParams) error {
	return http.ErrNotImplemented
}
