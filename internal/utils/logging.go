package utils

import (
	"fmt"

	"github.com/Galdoba/choretracker/internal/core/ports"
)

func LogError(logger ports.Logger, msg string, err error) error {
	logger.Errorf(msg)
	if err == nil {
		return fmt.Errorf("%v", msg)
	}
	return fmt.Errorf("%v: %v", msg, err)
}
