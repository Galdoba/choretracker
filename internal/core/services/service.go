package services

import (
	"errors"
	"fmt"
	"os/user"
	"path/filepath"
	"time"

	"github.com/Galdoba/choretracker/internal/core/domain"
	"github.com/Galdoba/choretracker/internal/core/dto"
	"github.com/Galdoba/choretracker/internal/core/ports"
	"github.com/Galdoba/choretracker/internal/helpers"
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

type TaskRequest interface {
	dto.CreateRequest | dto.ReadRequest | dto.UpdateRequest | dto.DeleteRequest
}

func HandleRequest[T TaskRequest](ts *TaskService, req T) (domain.Chore, error) {
	switch r := any(req).(type) {
	case dto.CreateRequest:
		return domain.Chore{}, ts.CreateTask(r)
	case dto.ReadRequest:
		return ts.ReadTask(r)
	case dto.UpdateRequest:
		return domain.Chore{}, ts.UpdateTask(r)
	case dto.DeleteRequest:
		return domain.Chore{}, ts.DeleteTask(r)
	}
	return domain.Chore{}, ts.logError("request is of unknown type", nil)
}

func (ts *TaskService) CreateTask(r dto.CreateRequest) error {
	ch, err := newChore(r)
	if err != nil {
		ts.logError("failed to create new chore", err)
	}
	if err := ts.Validator.Validate(ch); err != nil {
		return ts.logError("failed to validate chore", err)
	}
	if err := updateNotificationTime(&ch); err != nil {
		return ts.logError("failed to update notification time", err)
	}
	if err := ts.Storage.Create(ch); err != nil {
		return ts.logError("failed to create chore in storage", err)
	}
	ts.Logger.Infof("new chore created: %v", ch.Key())
	return nil
}

func newChore(cr dto.CreateRequest) (domain.Chore, error) {
	ch := domain.Chore{}
	openTime := time.Now()
	ch.ID = openTime.Unix()
	ch.Opened = openTime
	updateChore(&ch, cr.Content())

	if ch.Author == "" {
		currentUser, err := user.Current()
		if err != nil {
			return ch, fmt.Errorf("failed to get current username: %v", err)
		}
		ch.Author = filepath.Base(currentUser.Username)
	}

	return ch, nil
}

type idRetriever interface {
	getID() *int64
}

func (ts TaskService) logError(msg string, err error) error {
	ts.Logger.Errorf(msg)
	if err == nil {
		return errors.New(msg)
	}
	return fmt.Errorf("%v: %v", msg, err)
}

func loadFromStorage(ts *TaskService, id *int64) (domain.Chore, error) {
	if id == nil {
		return domain.Chore{}, utils.LogError(ts.Logger, "chore id not provided", nil)
	}

	for _, err := range []bool{*id == 0} {
		if err {
			return domain.Chore{}, utils.LogError(ts.Logger, "id can't be equal to 0", nil)
		}
	}

	ch, err := ts.Storage.Read(*id)
	if err != nil {
		return domain.Chore{}, utils.LogError(ts.Logger, "failed to read from storage", err)
	}
	return ch, nil
}

func (ts *TaskService) ReadTask(r dto.ReadRequest) (domain.Chore, error) {
	ch, err := loadFromStorage(ts, r.ID)
	if err != nil {
		return domain.Chore{}, err
	}
	ts.Logger.Infof("chore read: %v", ch.Key())
	return ch, nil
}

func (ts *TaskService) UpdateTask(r dto.UpdateRequest) error {
	ch, err := loadFromStorage(ts, r.ID)
	if err != nil {
		return err
	}

	updateChore(&ch, r.Content())

	if err := ts.Validator.Validate(ch); err != nil {
		return ts.logError("failed to validate updated chore", err)
	}

	if err := updateNotificationTime(&ch); err != nil {
		return ts.logError("failed to update notification time", err)
	}
	if err := ts.Storage.Update(ch); err != nil {
		return ts.logError("failed to update chore in storage", err)
	}
	ts.Logger.Infof("chore %v updated", ch.Key())

	return nil
}

func (ts *TaskService) DeleteTask(r dto.DeleteRequest) error {
	ch, err := loadFromStorage(ts, r.ID)
	if err != nil {
		return err
	}
	key := ch.Key()
	if err := ts.Storage.Delete(ch.ID); err != nil {
		return ts.logError("failed to delete chore", err)
	}
	ts.Logger.Infof("chore %v deleted", key)
	return nil
}

func updateNotificationTime(ch *domain.Chore) error {
	exp, err := cronexpr.Parse(ch.Schedule)
	if err != nil {
		return fmt.Errorf("failed to parse chore shedule: %v", err)
	}
	ch.NextNotification = exp.Next(time.Now())
	return nil
}

func updateChore(ch *domain.Chore, r dto.ChoreContent) {
	ch.Title = setUpdated(ch.Title, r.Title)
	ch.Description = setUpdated(ch.Description, r.Description)
	ch.Author = setUpdated(ch.Author, r.Author)
	ch.Schedule = setUpdated(ch.Schedule, r.Schedule)
	ch.Comment = setUpdated(ch.Comment, r.Comment)
}

func setUpdated(old string, new *string) string {
	return helpers.SetUpdatedContent(old, new)
}
