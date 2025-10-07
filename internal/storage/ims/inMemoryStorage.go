package ims

import (
	"fmt"
	"sync"

	"github.com/Galdoba/choretracker/internal/models/chore"
)

type InMemoryStorage struct {
	chores map[int]*chore.Chore
	mutext sync.Mutex
}

func NewInMemoryStorage() *InMemoryStorage {
	imStrg := InMemoryStorage{}
	imStrg.chores = make(map[int]*chore.Chore)
	return &imStrg
}

func (ims *InMemoryStorage) Create(ch *chore.Chore) error {
	ims.mutext.Lock()
	defer ims.mutext.Unlock()

	if _, ok := ims.chores[ch.ID]; ok {
		return fmt.Errorf("chore %v already created", ch.ID)
	}

	ims.chores[ch.ID] = ch

	return nil
}

// Update(*chore.Chore) error
func (ims *InMemoryStorage) Update(ch *chore.Chore) error {
	ims.mutext.Lock()
	defer ims.mutext.Unlock()

	if _, ok := ims.chores[ch.ID]; !ok {
		return fmt.Errorf("chore %v does not exist", ch.ID)
	}

	ims.chores[ch.ID] = ch

	return nil
}

// Read(int) (*chore.Chore, error)
func (ims *InMemoryStorage) Read(id int) (*chore.Chore, error) {
	ims.mutext.Lock()
	defer ims.mutext.Unlock()

	ch, ok := ims.chores[id]
	switch ok {
	default:
		return nil, fmt.Errorf("chore %v does not exist", id)
	case true:
		return ch, nil
	}
}

// Delete(*chore.Chore) error
func (ims *InMemoryStorage) Delete(id int) error {
	ims.mutext.Lock()
	defer ims.mutext.Unlock()

	if _, ok := ims.chores[id]; !ok {
		return fmt.Errorf("chore %v does not exist", id)
	}

	delete(ims.chores, id)
	return nil
}

// GetAll() ([]*chore.Chore, error)
func (ims *InMemoryStorage) GetAll() ([]*chore.Chore, error) {
	ims.mutext.Lock()
	defer ims.mutext.Unlock()

	chores := []*chore.Chore{}
	for _, v := range ims.chores {
		chores = append(chores, v)

	}
	return chores, nil
}
