package actions

import (
	"context"
	"fmt"

	"github.com/Galdoba/choretracker/cmd/choretracker/app/flags"
	"github.com/Galdoba/choretracker/internal/appcontext"
	"github.com/Galdoba/choretracker/internal/models/chore"
	"github.com/Galdoba/choretracker/internal/storage/js"
	"github.com/Galdoba/consolio/prompt"
	"github.com/urfave/cli/v3"
)

func Get(actx *appcontext.AppContext) cli.ActionFunc {
	return func(ctx context.Context, c *cli.Command) error {
		id := c.Int64(flags.CHORE_ID)
		store, err := js.New(actx.Config().StoragePath)
		if err != nil {
			return fmt.Errorf("failed to open persistent storage")
		}
		ch := &chore.Chore{}
		switch c.Bool(flags.GLOBAL_CLI) {
		case false:
			ch, err = getChoreTUI(store)
		case true:
			ch, err = getChoreCLI(store, id)
		}
		if err != nil {
			return fmt.Errorf("failed to get chore data: %v", err)
		}
		fmt.Println(ch.String())
		return nil
	}
}

// getChoreTUI return chore from storage in interactive mode
func getChoreTUI(store *js.JsonStore) (*chore.Chore, error) {
	chores, err := store.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get chore list from storage: %v", err)
	}
	ids := []int64{}
	choreItems := []*prompt.Item{}
	for _, ch := range chores {
		ids = append(ids, ch.ID)
		choreItems = append(choreItems, prompt.NewItem(fmt.Sprintf("%v (id: %v) %v", ch.Title, ch.ID, ch.Description), ch))
	}
	found, err := prompt.SearchItem(prompt.FromItems(choreItems))
	if err != nil {
		return nil, fmt.Errorf("chore search failed: %v", err)
	}
	ch := found.Payload().(*chore.Chore)
	return ch, nil
}

// getChoreCLI return chore from storage in command line mode
func getChoreCLI(store *js.JsonStore, id int64) (*chore.Chore, error) {
	if id == 0 {
		return nil, fmt.Errorf("chore id is not set\nuse 'id' option to set chore id")
	}
	ch, err := store.Read(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read chore: %v", err)
	}
	return ch, nil
}
