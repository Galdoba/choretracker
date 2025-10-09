package ims

import (
	"fmt"
	"sync"

	"github.com/Galdoba/choretracker/internal/core/domain"
)

type InMemoryStorage struct {
	chores map[int64]domain.Chore
	mutext sync.Mutex
}

func NewInMemoryStorage() *InMemoryStorage {
	imStrg := InMemoryStorage{}
	imStrg.chores = make(map[int64]domain.Chore)
	return &imStrg
}

func (ims *InMemoryStorage) Create(ch domain.Chore) error {
	ims.mutext.Lock()
	defer ims.mutext.Unlock()

	if _, ok := ims.chores[ch.ID]; ok {
		return fmt.Errorf("chore %v already created", ch.ID)
	}

	ims.chores[ch.ID] = ch

	return nil
}

// Update(*chore.Chore) error
func (ims *InMemoryStorage) Update(ch domain.Chore) error {
	ims.mutext.Lock()
	defer ims.mutext.Unlock()

	if _, ok := ims.chores[ch.ID]; !ok {
		return fmt.Errorf("chore %v does not exist", ch.ID)
	}

	ims.chores[ch.ID] = ch

	return nil
}

// Read(int) (*chore.Chore, error)
func (ims *InMemoryStorage) Read(id int64) (domain.Chore, error) {
	ims.mutext.Lock()
	defer ims.mutext.Unlock()

	ch, ok := ims.chores[id]
	switch ok {
	default:
		return ch, fmt.Errorf("chore %v does not exist", id)
	case true:
		return ch, nil
	}
}

// Delete(*chore.Chore) error
func (ims *InMemoryStorage) Delete(id int64) error {
	ims.mutext.Lock()
	defer ims.mutext.Unlock()

	if _, ok := ims.chores[id]; !ok {
		return fmt.Errorf("chore %v does not exist", id)
	}

	delete(ims.chores, id)
	return nil
}

// GetAll() ([]*chore.Chore, error)
func (ims *InMemoryStorage) GetAll() ([]domain.Chore, error) {
	ims.mutext.Lock()
	defer ims.mutext.Unlock()

	chores := []domain.Chore{}
	for _, v := range ims.chores {
		chores = append(chores, v)

	}
	return chores, nil
}
