package shedule

import "fmt"

type CronEntry struct {
	Shedule []string `json:"shedule"`
	Comment string   `json:"comment,omitempty"`
}

func NewEntry(shedule []string) CronEntry {
	return CronEntry{Shedule: shedule}
}

func (ce CronEntry) WithComment(comment string) CronEntry {
	ce.Comment = comment
	return ce
}

func validate(ce CronEntry) error {
	if len(ce.Shedule) != 5 {
		return fmt.Errorf("cron expect 5 values in shedule: %v", ce.Shedule)
	}
	return nil
}

//отдельная функция без интерфейса
// Next(CronEntry, ...CronEntry) (CronEntry, error)
