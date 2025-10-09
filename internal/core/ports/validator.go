package ports

import "github.com/Galdoba/choretracker/internal/core/domain"

type Validator interface {
	Validate(domain.Chore) error
}
