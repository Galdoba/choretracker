package storage

import (
	"fmt"

	"github.com/Galdoba/choretracker/internal/appcontext"
	"github.com/Galdoba/choretracker/internal/core/ports"
	"github.com/Galdoba/choretracker/internal/infrastructure/storage/ims"
	"github.com/Galdoba/choretracker/internal/infrastructure/storage/js"
)

// type storageManager struct {
// 	inMemory *ims.InMemoryStorage
// 	json     *js.JsonStore
// }

type StorageType string

const (
	InMemoryStorage StorageType = "ims"
	JsonStorage     StorageType = "js"
)

// var store *storageManager

// func NewJsonStorage(cfg appcontext.Config) Storage {
// 	store, json = js.New(cfg.StoragePath)
// 	return js.New()
// }
//

func NewStorage(st StorageType, cfg *appcontext.Config) (ports.Storage, error) {
	switch st {
	case InMemoryStorage:
		return ims.NewInMemoryStorage(), nil
	case JsonStorage:
		store, err := js.New(cfg.StoragePath)
		return store, err
	default:
		return nil, fmt.Errorf("unknown storage type: %v", st)

	}
}

// type Storage interface {
// 	Create(*domain.Chore) error
// 	Update(*domain.Chore) error
// 	Read(int) (*domain.Chore, error)
// 	Delete(*domain.Chore) error
// 	GetAll() ([]*domain.Chore, error)
// }
