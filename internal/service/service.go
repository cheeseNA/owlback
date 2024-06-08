package service

import (
	"context"
	"fmt"
	"github.com/cheeseNA/owlback/internal/funccall"
	api "github.com/cheeseNA/owlback/internal/ogen"
	"github.com/cheeseNA/owlback/internal/repository"
	"go.uber.org/zap"
	"net/url"
)

type Service struct {
	repo        repository.ITaskRepository
	funcService funccall.IFuncService
	logger      *zap.Logger // TODO: replace with context logger interface
}

func NewService(repo repository.ITaskRepository, funcService funccall.IFuncService, logger *zap.Logger) *Service {
	return &Service{
		repo:        repo,
		funcService: funcService,
		logger:      logger,
	}
}

func (s *Service) CrateTask(ctx context.Context, req api.OptTaskRequest) error {
	taskReq, ok := req.Get()
	if !ok {
		return fmt.Errorf("invalid request")
	}
	task := repository.Task{
		SiteURL:        taskReq.SiteURL.String(),
		ConditionQuery: taskReq.ConditionQuery,
		DurationDay:    taskReq.DurationDay,
		IsPublic:       taskReq.IsPublic,
		IsPaused:       false,
	}
	return s.repo.CreateTask(task)
}

func (s *Service) DeleteTaskByID(ctx context.Context, params api.DeleteTaskByIDParams) error {
	return s.repo.DeleteTaskByID(params.TaskId)
}

func (s *Service) GetTaskByID(ctx context.Context, params api.GetTaskByIDParams) (api.GetTaskByIDRes, error) {
	task, err := s.repo.GetTaskByID(params.TaskId)
	if err != nil {
		return &api.GetTaskByIDNotFound{}, err
	}
	siteUrl, err := url.Parse(task.SiteURL)
	if err != nil {
		return nil, err
	}
	var lastCrawledAt api.OptDateTime
	if task.LastCrawledAt != nil {
		lastCrawledAt = api.NewOptDateTime(*task.LastCrawledAt)
	} else {
		lastCrawledAt.Reset()
	}
	return &api.TaskResponse{
		SiteURL:        *siteUrl,
		ConditionQuery: task.ConditionQuery,
		DurationDay:    task.DurationDay,
		IsPublic:       task.IsPublic,
		ID:             task.ID,
		CreatedAt:      task.CreatedAt,
		CreatedBy:      task.CreatedBy,
		UpdatedAt:      task.UpdatedAt,
		LastCrawledAt:  lastCrawledAt,
		IsPaused:       task.IsPaused,
	}, nil
}

func (s *Service) GetTasks(ctx context.Context) ([]api.TaskResponse, error) {
	tasks, err := s.repo.GetTasks()
	if err != nil {
		return nil, err
	}
	res := make([]api.TaskResponse, len(tasks))
	for i, task := range tasks {
		siteUrl, err := url.Parse(task.SiteURL)
		var lastCrawledAt api.OptDateTime
		if task.LastCrawledAt != nil {
			lastCrawledAt = api.NewOptDateTime(*task.LastCrawledAt)
		} else {
			lastCrawledAt.Reset()
		}
		if err != nil {
			return nil, err
		}
		res[i] = api.TaskResponse{
			SiteURL:        *siteUrl,
			ConditionQuery: task.ConditionQuery,
			DurationDay:    task.DurationDay,
			IsPublic:       task.IsPublic,
			ID:             task.ID,
			CreatedAt:      task.CreatedAt,
			CreatedBy:      task.CreatedBy,
			UpdatedAt:      task.UpdatedAt,
			LastCrawledAt:  lastCrawledAt,
			IsPaused:       task.IsPaused,
		}
	}
	return res, nil
}

func (s *Service) Healthz(ctx context.Context) error {
	return nil
}

func (s *Service) PostCronWrpouiqjflsadkmxcvz780923(ctx context.Context) error {
	return nil
}
