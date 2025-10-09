package validator

import (
	"fmt"
	"time"

	"github.com/Galdoba/choretracker/internal/core"
	"github.com/Galdoba/choretracker/pkg/cronexpr"
)

type choreValidator struct {
}

func (cv *choreValidator) Validate(ch core.Chore) error {
	if ch.Schedule == "" {
		return fmt.Errorf("shedule is not set")
	}
	exp, err := cronexpr.Parse(ch.Schedule)
	if err != nil {
		return fmt.Errorf("failed to parse cron shedule: %v", err)
	}
	ch.NextNotification = exp.Next(time.Now())

	return nil
}
