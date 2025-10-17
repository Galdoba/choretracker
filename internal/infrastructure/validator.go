package infrastructure

import (
	"github.com/Galdoba/choretracker/internal/core/dto"
)

type defaultValidator struct{}

func DefaultValidator() *defaultValidator {
	return &defaultValidator{}
}

func (v *defaultValidator) ValidateRequest(req dto.ToServiceRequest) error {
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

func validateContent(c dto.ChoreContent) error {

	return nil
}

func validateIdentity(i dto.ChoreIdentity) error {

	return nil
}
