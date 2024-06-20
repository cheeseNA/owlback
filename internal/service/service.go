package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/cheeseNA/owlback/internal/funccall"
	"github.com/cheeseNA/owlback/internal/middleware"
	api "github.com/cheeseNA/owlback/internal/ogen"
	"github.com/cheeseNA/owlback/internal/repository"
	"go.uber.org/zap"
	"gorm.io/gorm"
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

func (s *Service) CrateTask(ctx context.Context, req api.OptTaskRequest) (api.CrateTaskRes, error) {
	s.logger.Info("Create task")
	taskReq, ok := req.Get()
	if !ok {
		s.logger.Error("Invalid request")
		return &api.CrateTaskBadRequest{}, fmt.Errorf("invalid request")
	}
	user := middleware.GetUser(ctx)
	if user == nil {
		s.logger.Error("Unauthorized")
		return &api.CrateTaskUnauthorized{}, fmt.Errorf("unauthorized")
	}
	task := repository.Task{
		SiteURL:        taskReq.SiteURL.String(),
		ConditionQuery: taskReq.ConditionQuery,
		DurationDay:    taskReq.DurationDay,
		IsPublic:       taskReq.IsPublic,
		IsPaused:       false,
		User:           repository.TokenToUserModel(user),
	}
	if err := s.repo.CreateTask(task); err != nil {
		s.logger.Error("Failed to create task", zap.Error(err))
		return &api.CrateTaskInternalServerError{}, err
	}
	return &api.CrateTaskCreated{}, nil
}

func (s *Service) DeleteTaskByID(ctx context.Context, params api.DeleteTaskByIDParams) (api.DeleteTaskByIDRes, error) {
	s.logger.Info("Delete task by id: ", zap.Any("id", params.TaskId))
	task, err := s.repo.GetTaskByID(params.TaskId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { // todo: not depend on actual impl.
			s.logger.Error("Task not found", zap.Error(err))
			return &api.DeleteTaskByIDNotFound{}, err
		}
		s.logger.Error("Failed to get task", zap.Error(err))
		return &api.DeleteTaskByIDInternalServerError{}, err
	}
	if task.UserID != middleware.GetUser(ctx).UID {
		s.logger.Error("Unauthorized")
		return &api.DeleteTaskByIDUnauthorized{}, fmt.Errorf("unauthorized or forbidden")
	}
	err = s.repo.DeleteTaskByID(params.TaskId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { // TODO: use transaction
			s.logger.Error("Task not found", zap.Error(err))
			return &api.DeleteTaskByIDNotFound{}, err
		}
		s.logger.Error("Failed to delete task", zap.Error(err))
		return &api.DeleteTaskByIDInternalServerError{}, err
	}
	return &api.DeleteTaskByIDOK{}, nil
}

func (s *Service) GetTaskByID(ctx context.Context, params api.GetTaskByIDParams) (api.GetTaskByIDRes, error) {
	s.logger.Info("Get task by id: ", zap.Any("id", params.TaskId))
	task, err := s.repo.GetTaskByID(params.TaskId)
	if err != nil {
		s.logger.Error("Failed to get task", zap.Error(err))
		return &api.GetTaskByIDNotFound{}, err
	}
	if !task.IsPublic && task.UserID != middleware.GetUser(ctx).UID {
		s.logger.Error("Unauthorized")
		return &api.GetTaskByIDUnauthorized{}, fmt.Errorf("unauthorized or forbidden")
	}

	siteUrl, err := url.Parse(task.SiteURL)
	if err != nil {
		s.logger.Error("Failed to parse site url", zap.Error(err))
		return &api.GetTaskByIDInternalServerError{}, err
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
		UserID:         task.UserID,
		UpdatedAt:      task.UpdatedAt,
		LastCrawledAt:  lastCrawledAt,
		IsPaused:       task.IsPaused,
	}, nil
}

func (s *Service) GetTasks(ctx context.Context) (api.GetTasksRes, error) {
	s.logger.Info("Get tasks")
	tasks, err := s.repo.GetTasks()
	if err != nil {
		s.logger.Error("Failed to get tasks", zap.Error(err))
		return &api.GetTasksInternalServerError{}, err
	}
	res := make([]api.TaskResponse, len(tasks))
	for i, task := range tasks {
		siteUrl, err := url.Parse(task.SiteURL)
		if err != nil {
			s.logger.Error("Failed to parse site url", zap.Error(err))
			return &api.GetTasksInternalServerError{}, err
		}
		var lastCrawledAt api.OptDateTime
		if task.LastCrawledAt != nil {
			lastCrawledAt = api.NewOptDateTime(*task.LastCrawledAt)
		} else {
			lastCrawledAt.Reset()
		}

		res[i] = api.TaskResponse{
			SiteURL:        *siteUrl,
			ConditionQuery: task.ConditionQuery,
			DurationDay:    task.DurationDay,
			IsPublic:       task.IsPublic,
			ID:             task.ID,
			CreatedAt:      task.CreatedAt,
			UserID:         task.UserID,
			UpdatedAt:      task.UpdatedAt,
			LastCrawledAt:  lastCrawledAt,
			IsPaused:       task.IsPaused,
		}
	}
	okRes := api.GetTasksOKApplicationJSON(res)
	return &okRes, nil
}

func (s *Service) Healthz(ctx context.Context) error {
	s.logger.Info("healthz")
	return nil
}

func (s *Service) PostCronWrpouiqjflsadkmxcvz780923(ctx context.Context) error {
	s.logger.Info("Crawling tasks")
	tasks, err := s.repo.GetTasksToCrawl()
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
