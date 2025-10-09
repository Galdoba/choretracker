package domain

import "time"

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
