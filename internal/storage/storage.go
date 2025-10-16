package storage

import (
	"github.com/Galdoba/choretracker/internal/core/domain"
	"github.com/Galdoba/choretracker/internal/storage/ims"
	"github.com/Galdoba/choretracker/internal/storage/js"
)

type storageManager struct {
	inMemory *ims.InMemoryStorage
	json     *js.JsonStore
}

var store *storageManager

// func NewJsonStorage(cfg appcontext.Config) Storage {
// 	store, json = js.New(cfg.StoragePath)
// 	return js.New()
// }

type Storage interface {
	Create(domain.Chore) error
	Update(domain.Chore) error
	Read(int) (domain.Chore, error)
	Delete(domain.Chore) error
	GetAll() ([]domain.Chore, error)
}
