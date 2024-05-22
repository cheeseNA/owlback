package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"net/url"
	"sync"

	ogen "github.com/cheeseNA/owlback/ogen"
)

type service struct {
	tasks map[uuid.UUID]ogen.Task
	mux   sync.Mutex
}

func (s *service) AddTask(ctx context.Context, task *ogen.Task) (*ogen.Task, error) {
	s.mux.Lock()
	defer s.mux.Unlock()
	taskId := uuid.New()
	s.tasks[taskId] = *task
	task.ID.SetTo(taskId)
	return task, nil
}

func (s *service) GetTaskById(ctx context.Context, params ogen.GetTaskByIdParams) (ogen.GetTaskByIdRes, error) {
	s.mux.Lock()
	defer s.mux.Unlock()
	task, ok := s.tasks[params.TaskId]
	if !ok {
		return &ogen.GetTaskByIdNotFound{}, nil
	}
	return &task, nil
}

func (s *service) UpdateTask(ctx context.Context, params ogen.UpdateTaskParams) error {
	s.mux.Lock()
	defer s.mux.Unlock()
	task, ok := s.tasks[params.TaskId]
	if !ok {
		return fmt.Errorf("task not found")
	}
	if val, ok := params.SiteURL.Get(); ok {
		url, err := url.Parse(val)
		if err != nil {
			return err
		}
		task.SiteURL = *url
	}
	if val, ok := params.Condition.Get(); ok {
		task.Condition = val
	}
	if val, ok := params.DurationDay.Get(); ok {
		task.DurationDay = val
	}
	s.tasks[params.TaskId] = task
	return nil
}
