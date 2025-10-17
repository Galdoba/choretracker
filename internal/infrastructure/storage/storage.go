package storage

import (
	"fmt"

	"github.com/Galdoba/choretracker/internal/core/ports"
	"github.com/Galdoba/choretracker/internal/infrastructure/storage/ims"
	"github.com/Galdoba/choretracker/internal/infrastructure/storage/js"
)

type StorageType string

const (
	InMemoryStorage StorageType = "ims"
	JsonStorage     StorageType = "js"
)

func NewStorage(st StorageType, path string) (ports.Storage, error) {
	switch st {
	case InMemoryStorage:
		return ims.NewInMemoryStorage(), nil
	case JsonStorage:
		store, err := js.New(path)
		return store, err
	default:
		return nil, fmt.Errorf("unknown storage type: %v", st)

	}
}
