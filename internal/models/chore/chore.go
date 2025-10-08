package chore

import (
	"fmt"
	"os/user"
	"path/filepath"
	"time"

	"github.com/Galdoba/choretracker/pkg/cronexpr"
)

type Chore struct {
	ID               int64     `json:"id"`
	Title            string    `json:"title"`
	Description      string    `json:"description,omitempty"`
	Author           string    `json:"author"`
	Opened           time.Time `json:"opened"`
	NextNotification time.Time `json:"next notification"`
	CronSchedule     string    `json:"schedule"`
	Comments         string    `json:"comments,omitempty"`
}

func NewChore(knownFields ...ChoreOption) (*Chore, error) {
	ch := &Chore{}
	openTime := time.Now()
	ch.ID = openTime.Unix()
	ch.Opened = openTime
	for _, addFieldTo := range knownFields {
		addFieldTo(ch)
	}

	if ch.Author == "" {
		currentUser, err := user.Current()
		if err != nil {
			return nil, fmt.Errorf("failed to get current username: %v", err)
		}
		ch.Author = filepath.Base(currentUser.Username)
	}

	ch.Update()

	return ch, nil
}

func (ch *Chore) Update() error {
	if err := validateCronExpression(ch.CronSchedule); err != nil {
		return fmt.Errorf("chore shedule validation failed: %v", err)
	}
	exp := cronexpr.MustParse(ch.CronSchedule)
	ch.NextNotification = exp.Next(time.Now())
	return nil
}

func validateCronExpression(expression string) error {
	if expression == "" {
		return fmt.Errorf("shedule is not set")
	}
	_, err := cronexpr.Parse(expression)
	if err != nil {
		return fmt.Errorf("failed to parse cron shedule: %v", err)
	}
	return nil
}

func (ch *Chore) String() string {
	s := fmt.Sprintf("chore: %v", ch.Title) + "\n"
	s += fmt.Sprintf("ID: %v", ch.ID) + "\n"
	s += fmt.Sprintf("started: %v", ch.Opened.Format(time.DateTime)) + "\n"
	s += fmt.Sprintf("description: %v", ch.Description) + "\n"
	s += fmt.Sprintf("shedule: %v", ch.CronSchedule) + "\n"
	s += fmt.Sprintf("next trigger time: %v", ch.NextNotification.Format(time.DateTime))

	return s
}

func (ch *Chore) Validate() error {
	if ch.Title == "" {
		return fmt.Errorf("chore name is not set")
	}
	if ch.ID == 0 {
		return fmt.Errorf("chore ID is not set")
	}
	if err := validateCronExpression(ch.CronSchedule); err != nil {
		return err
	}
	return nil
}

type ChoreOption func(*Chore)

func WithTitle(title string) ChoreOption {
	return func(c *Chore) {
		c.Title = title
	}
}

func WithDescription(descr string) ChoreOption {
	return func(c *Chore) {
		c.Description = descr
	}
}

func WithAuthor(author string) ChoreOption {
	return func(c *Chore) {
		c.Author = author
	}
}

func WithShedule(shedule string) ChoreOption {
	return func(c *Chore) {
		c.CronSchedule = shedule
	}
}

func WithComment(comment string) ChoreOption {
	return func(c *Chore) {
		c.Comments = comment
	}
}
