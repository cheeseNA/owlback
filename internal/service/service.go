package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/cheeseNA/owlback/internal/funccall"
	api "github.com/cheeseNA/owlback/internal/ogen"
	"github.com/cheeseNA/owlback/internal/repository"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"net/url"
	"time"
)

type Service struct {
	repo        repository.ITaskRepository
	funcService funccall.IFuncService
	logger      *zap.Logger // TODO: replace with context logger interface
}

const scrapeCooldown = 1 * time.Minute

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
		CreatedBy:      uuid.New(), // TODO: replace with user id
	}
	return s.repo.CreateTask(task)
}

func (s *Service) DeleteTaskByID(ctx context.Context, params api.DeleteTaskByIDParams) error {
	err := s.repo.DeleteTaskByID(params.TaskId)
	//if err != nil {
	//	if errors.Is(err, gorm.ErrRecordNotFound) { // todo: not depend on actual impl.
	//		return &api.DeleteTaskByIDNotFound // todo: re-design api
	//	}
	//}
	return err
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
	tasks, err := s.repo.GetTasksToCrawl()
	s.logger.Info("Crawling tasks")
	if err != nil {
		return err
	}
	for _, task := range tasks {
		s.logger.Info("Crawling task", zap.Any("task", task))
		lastContent := ""
		if task.LastContent != nil {
			lastContent = *task.LastContent
		}
		now := time.Now()
		req := funccall.Request{
			Url: task.SiteURL,
			//OpenaiKey:   "sk-proj-md1SX1jfiAVlxu7oMH3FT3BlbkFJbFtePRkUFXMkoC7VGDWf", // TODO: add openai key to model
			GeminiKey:   "AIzaSyCHyWg2J9Ntgb6YZ-p9f_Kmezo0vO9xv7I",
			Query:       task.ConditionQuery,
			LastContent: lastContent,
			IsStrict:    false, // TODO: add strict to model
		}
		res, e := s.funcService.CallFunc(req)
		time.Sleep(scrapeCooldown)
		if e != nil {
			s.logger.Error("Failed to call function", zap.Error(e))
			err = errors.Join(err, e)
			continue
		}
		if *res.IsTriggered {
			s.logger.Info("Triggered", zap.Any("task", task))
			// TODO: send notification
		}
		// TODO: lock & transaction
		err = errors.Join(err, s.repo.UpdateTask(task, repository.Task{LastContent: res.NewContent, LastCrawledAt: &now}))
	}
	return err
}
