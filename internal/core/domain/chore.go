package domain

import (
	"fmt"
	"time"

	"github.com/Galdoba/choretracker/internal/constants"
)

func fields() []string {
	return []string{
		constants.Fld_Title, constants.Fld_Descr, constants.Fld_Author, constants.Fld_Schedule, constants.Fld_Comment,
	}
}

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

func (ch Chore) GetID() (int64, bool) {
	return ch.ID, true
}

func (ch Chore) Content() map[string]string {
	c := make(map[string]string)
	for i, key := range fields() {
		switch i {
		case 0:
			c[key] = ch.Title
		case 1:
			c[key] = ch.Description
		case 2:
			c[key] = ch.Author
		case 3:
			c[key] = ch.Schedule
		case 4:
			c[key] = ch.Comment
		}
	}
	return c
}
func (ch *Chore) GetTitle() (string, bool) {
	return ch.Title, true
}
func (ch *Chore) GetDescription() (string, bool) {
	return ch.Description, true
}
func (ch *Chore) GetAuthor() (string, bool) {
	return ch.Author, true
}
func (ch *Chore) GetOpened() time.Time {
	return ch.Opened
}
func (ch *Chore) GetNextNotification() time.Time {
	return ch.NextNotification
}
func (ch *Chore) GetSchedule() (string, bool) {
	return ch.Schedule, true
}
func (ch *Chore) GetComment() (string, bool) {
	return ch.Comment, true
}

func (ch *Chore) Validate() error {
	if ch.ID == 0 {
		return fmt.Errorf("id must not be 0")
	}
	if ch.Title == "" {
		return fmt.Errorf("title must be set")
	}
	if ch.Schedule == "" {
		return fmt.Errorf("schedule must be set")
	}
	return nil
}
