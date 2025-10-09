package infrastructure

import (
	"fmt"

	"github.com/Galdoba/choretracker/internal/core/domain"
	"github.com/Galdoba/choretracker/pkg/cronexpr"
)

type defaultValidator struct{}

func DefaultValidator() *defaultValidator {
	return &defaultValidator{}
}

func (v *defaultValidator) Validate(ch domain.Chore) error {
	if ch.ID == 0 {
		return fmt.Errorf("chore id is not set")
	}
	if ch.Title == "" {
		return fmt.Errorf("chore title is not set")
	}
	if ch.Schedule == "" {
		return fmt.Errorf("chore schedule is not set")
	}
	if _, err := cronexpr.Parse(ch.Schedule); err != nil {
		return fmt.Errorf("chore schedule is invalid: %v", err)
	}
	return nil
}
