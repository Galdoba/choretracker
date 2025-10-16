package infrastructure

import (
	"fmt"

	"github.com/Galdoba/choretracker/internal/core/dto"
	"github.com/Galdoba/choretracker/pkg/cronexpr"
)

type defaultValidator struct{}

func DefaultValidator() *defaultValidator {
	return &defaultValidator{}
}

func (v *defaultValidator) ValidateRequest(req *dto.ToServiceRequest) error {
	switch req.Action {
	case dto.Create:
		if err := validateContent(req.Fields); err != nil {
			return err
		}
	case dto.Read, dto.Delete:
		if err := validateIdentity(req.Identity); err != nil {
			return err
		}
	case dto.Update:
		if err := validateIdentity(req.Identity); err != nil {
			return err
		}
		if err := validateContent(req.Fields); err != nil {
			return err
		}

	}
	return nil
}

func (v *defaultValidator) ValidateResponce(res dto.FromServiceResponce) error {
	if err := validateIdentity(res.Id()); err != nil {
		return fmt.Errorf("failed to validate response: %v", err)
	}
	if err := validateContent(res.Content()); err != nil {
		return fmt.Errorf("failed to validate response: %v", err)
	}
	return nil
}

func validateContent(c dto.ChoreContent) error {
	if c.Title != nil && *c.Title == "" {
		return fmt.Errorf("title must not be empty")
	}
	if c.Author != nil && *c.Author == "" {
		return fmt.Errorf("author must not be empty")
	}
	if c.Schedule != nil {
		_, err := cronexpr.Parse(*c.Schedule)
		if err != nil {
			return fmt.Errorf("schedule is invalid: %v", err)
		}
	}
	return nil
}

func validateIdentity(i dto.ChoreIdentity) error {
	if i.ID == nil || *i.ID == 0 {
		return fmt.Errorf("invalid identity: %v", i.ID)
	}
	return nil
}
