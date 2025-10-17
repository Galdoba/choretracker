package services

import (
	"errors"
	"fmt"
	"os/user"
	"path/filepath"
	"time"

	"github.com/Galdoba/choretracker/internal/constants"
	"github.com/Galdoba/choretracker/internal/core/domain"
	"github.com/Galdoba/choretracker/internal/core/dto"
	"github.com/Galdoba/choretracker/internal/core/ports"
	"github.com/Galdoba/choretracker/internal/utils"
	"github.com/Galdoba/choretracker/pkg/cronexpr"
)

type TaskService struct {
	Storage   ports.Storage
	Validator ports.Validator
	Logger    ports.Logger
}

func NewTaskService(storage ports.Storage, validator ports.Validator, logger ports.Logger) *TaskService {
	return &TaskService{
		Storage:   storage,
		Validator: validator,
		Logger:    logger,
	}
}

func (ts *TaskService) ServeRequest(req dto.ToServiceRequest) (*domain.Chore, error) {
	switch req.Action {
	case dto.Create:
		return ts.CreateTask(req)
	case dto.Read:
		return ts.ReadTask(req)
	case dto.Update:
		return ts.UpdateTask(req)
	case dto.Delete:
		return ts.DeleteTask(req)
	default:
		return nil, fmt.Errorf("unknown request type: %v", req.RequestType())
	}
}

func (ts *TaskService) CreateTask(req dto.ToServiceRequest) (*domain.Chore, error) {
	ch := domain.Chore{}
	if err := ts.Validator.ValidateRequest(req); err != nil {
		return nil, utils.LogError(ts.Logger, "invalid request", err)
	}
	ch, err := newChoreFromRequest(req)
	if err != nil {
		return nil, utils.LogError(ts.Logger, "failed to create new chore", err)
	}
	if err := ch.Validate(); err != nil {
		return nil, utils.LogError(ts.Logger, "failed to validate chore", err)
	}
	if err := updateNotificationTime(&ch); err != nil {
		return nil, utils.LogError(ts.Logger, "failed to update notification time", err)
	}
	if err := ts.Storage.Create(ch); err != nil {
		return nil, utils.LogError(ts.Logger, "failed to create chore in storage", err)
	}
	ts.Logger.Infof("new chore created: %v (id=%v)", ch.Title, ch.ID)
	return &ch, nil
}

func newChoreFromRequest(req dto.ToServiceRequest) (domain.Chore, error) {
	ch := domain.Chore{}
	if req.Action != dto.Create {
		return ch, fmt.Errorf("request type is not '%v'", dto.Create)
	}
	openTime := time.Now()
	ch.ID = openTime.Unix()
	ch.Opened = openTime

	updateChore(&ch, req.Fields.Content())

	if ch.Author == "" {
		currentUser, err := user.Current()
		if err != nil {
			return ch, fmt.Errorf("failed to get current username: %v", err)
		}
		ch.Author = filepath.Base(currentUser.Username)
	}

	return ch, nil
}

func (ts TaskService) logError(msg string, err error) error {
	ts.Logger.Errorf(msg)
	if err == nil {
		return errors.New(msg)
	}
	return fmt.Errorf("%v: %v", msg, err)
}

func loadFromStorage(ts *TaskService, id int64) (domain.Chore, error) {
	if id == 0 {
		return domain.Chore{}, utils.LogError(ts.Logger, "id can't be equal to 0", nil)
	}

	ch, err := ts.Storage.Read(id)
	if err != nil {
		return domain.Chore{}, utils.LogError(ts.Logger, "failed to read from storage", err)
	}
	return ch, nil
}

func (ts *TaskService) ReadTask(req dto.ToServiceRequest) (*domain.Chore, error) {
	ch := domain.Chore{}
	if err := ts.Validator.ValidateRequest(req); err != nil {
		return nil, utils.LogError(ts.Logger, "invalid request", err)
	}
	ch, err := loadFromStorage(ts, *req.Identity.ID)
	if err != nil {
		return nil, err
	}
	ts.Logger.Infof("chore read: '%v' (id=%v)", ch.Title, ch.ID)
	return &ch, nil
}

func (ts *TaskService) UpdateTask(req dto.ToServiceRequest) (*domain.Chore, error) {
	ch := domain.Chore{}
	if err := ts.Validator.ValidateRequest(req); err != nil {
		return nil, utils.LogError(ts.Logger, "invalid request", err)
	}
	ch, err := loadFromStorage(ts, *req.Identity.ID)
	if err != nil {
		return nil, err
	}

	updateChore(&ch, req.Fields.Content())

	if err := updateNotificationTime(&ch); err != nil {
		return nil, ts.logError("failed to update notification time", err)
	}
	if err := ts.Storage.Update(ch); err != nil {
		return nil, ts.logError("failed to update chore in storage", err)
	}

	ts.Logger.Infof("chore updated: '%v' (id=%v)", ch.Title, ch.ID)
	return &ch, nil
}

func (ts *TaskService) DeleteTask(req dto.ToServiceRequest) (*domain.Chore, error) {
	if err := ts.Validator.ValidateRequest(req); err != nil {
		return nil, utils.LogError(ts.Logger, "invalid request", err)
	}
	ch, err := loadFromStorage(ts, *req.Identity.ID)
	if err != nil {
		return nil, utils.LogError(ts.Logger, "failed to load chore from storage", err)
	}
	title := ch.Title
	id := ch.ID
	if err := ts.Storage.Delete(ch.ID); err != nil {
		return nil, utils.LogError(ts.Logger, "failed to delete chore", err)
	}
	ts.Logger.Infof("chore deleted: '%v' (id=%v)", title, id)
	return nil, nil
}

func updateNotificationTime(ch *domain.Chore) error {
	exp, err := cronexpr.Parse(ch.Schedule)
	if err != nil {
		return fmt.Errorf("failed to parse chore shedule: %v", err)
	}
	ch.NextNotification = exp.Next(time.Now())
	return nil
}

func updateChore(ch *domain.Chore, c map[string]string) {
	if title, ok := c[constants.Fld_Title]; ok {
		ch.Title = setUpdated(ch.Title, title)
	}
	if desc, ok := c[constants.Fld_Descr]; ok {
		ch.Description = setUpdated(ch.Description, desc)
	}
	if author, ok := c[constants.Fld_Author]; ok {
		ch.Author = setUpdated(ch.Author, author)
	}
	if schedule, ok := c[constants.Fld_Schedule]; ok {
		ch.Schedule = setUpdated(ch.Schedule, schedule)
	}
	if comment, ok := c[constants.Fld_Comment]; ok {
		ch.Comment = setUpdated(ch.Comment, comment)
	}
}

func setUpdated(old string, new string) string {
	return utils.SetUpdatedField(old, &new)
}
