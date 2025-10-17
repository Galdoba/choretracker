package tui

import (
	"fmt"

	"github.com/Galdoba/choretracker/internal/core/domain"
	"github.com/Galdoba/choretracker/internal/core/services"
	"github.com/Galdoba/consolio/prompt"
)

type ChoreSelector struct {
	ts *services.TaskService
}

// SelectChore return chore selected by user from storage
func SelectChore(ts *services.TaskService) (*domain.Chore, error) {
	cs := ChoreSelector{ts}
	return cs.selectChore()
}

func (cs *ChoreSelector) selectChore() (*domain.Chore, error) {
	chr := domain.Chore{}
	chores, err := cs.ts.Storage.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get chore list from storage: %v", err)
	}
	// ids := []int64{}
	selectionPool := []*prompt.Item{}
	for _, ch := range chores {
		selectionPool = append(selectionPool, prompt.NewItem(fmt.Sprintf("id:%v (%v)", ch.ID, ch.Title), ch))
	}
	ch, err := prompt.SearchItem(prompt.WithTitle("search chores:"), prompt.FromItems(selectionPool))
	if err != nil {
		return nil, err
	}
	chr = ch.Payload().(domain.Chore)
	return &chr, nil
}
