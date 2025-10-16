package domain

import (
	"fmt"
	"time"
)

type Chore struct {
	ID               int64     `json:"id"`
	Title            string    `json:"title"`
	Description      string    `json:"description,omitempty"`
	Author           string    `json:"author"`
	Opened           time.Time `json:"opened"`
	NextNotification time.Time `json:"next notification"`
	Schedule         string    `json:"schedule"`
	Comment          string    `json:"comment,omitempty"`
}

func (ch *Chore) Key() string {
	return fmt.Sprintf("%v (id=%v)", ch.Title, ch.ID)
}

func (ch *Chore) String() string {
	s := ""
	s += fmt.Sprintf("chore  : %v (id=%v)\n", ch.Title, ch.ID)
	s += fmt.Sprintf("shedule: %v", ch.Schedule)
	if ch.Author != "" {
		s += fmt.Sprintf("\n author: %v", ch.Author)
	}
	if ch.Description != "" {
		s += fmt.Sprintf("\ndescription:\n  %v", ch.Description)
	}
	if ch.Comment != "" {
		s += fmt.Sprintf("\ncomment:\n  %v", ch.Comment)
	}
	return s
}
