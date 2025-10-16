package ports

import (
	"github.com/Galdoba/choretracker/internal/core/dto"
)

type Validator interface {
	ValidateRequest(dto.ToServiceRequest) error
	ValidateResponse(dto.FromServiceResponce) error
}
