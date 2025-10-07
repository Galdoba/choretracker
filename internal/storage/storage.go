package storage

import (
	"github.com/Galdoba/choretracker/internal/appcontext"
	"github.com/Galdoba/choretracker/internal/models/chore"
	"github.com/Galdoba/choretracker/internal/storage/ims"
	"github.com/Galdoba/choretracker/internal/storage/js"
)

type storageManager struct {
	inMemory *ims.InMemoryStorage
	json     *js.JsonStore
}

var store *storageManager

func NewJsonStorage(cfg appcontext.Config) Storage {
store.json = js.New(cfg.)
	return js.New()
}

type Storage interface {
	Create(*chore.Chore) error
	Update(*chore.Chore) error
	Read(int) (*chore.Chore, error)
	Delete(*chore.Chore) error
	GetAll() ([]*chore.Chore, error)
}
