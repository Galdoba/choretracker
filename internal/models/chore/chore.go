package chore

import (
	"time"
)

type Chore struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	Author           string    `json:"author"`
	Opened           time.Time `json:"opened"`
	NextNotification time.Time `json:"next notification"`
	CronSchedule     []string  `json:"schedule"`
}

func NewChore(knownFields ...ChoreOption) (*Chore, error) {
	ch := Chore{}
	for _, addFieldTo := range knownFields {
		addFieldTo(&ch)
	}

	return &ch, nil
}

type ChoreOption func(*Chore)

func WithName(name string) ChoreOption {
	return func(c *Chore) {
		c.Name = name
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

func WithShedule(shedule []string) ChoreOption {
	return func(c *Chore) {
		c.CronSchedule = shedule
	}
}
