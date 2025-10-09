package actions

import (
	"context"
	"fmt"
	"os"

	"github.com/Galdoba/choretracker/cmd/choretracker/app/flags"
	"github.com/Galdoba/choretracker/internal/appcontext"
	"github.com/Galdoba/choretracker/internal/models/chore"
	"github.com/Galdoba/choretracker/internal/storage/js"
	"github.com/urfave/cli/v3"
)

func Add(actx *appcontext.AppContext) cli.ActionFunc {
	return func(ctx context.Context, c *cli.Command) error {
		title := c.String(flags.CHORE_TITLE)
		description := c.String(flags.CHORE_DESCRIPTION)
		shedule := c.String(flags.CHORE_SCHEDULE)
		comment := c.String(flags.CHORE_COMMENT)
		author := c.String(flags.CHORE_AUTHOR)
		opts := []chore.ChoreOption{}
		for i, value := range []string{title, description, shedule, comment, author} {
			if value == "" {
				continue
			}
			switch i {
			case 0:
				opts = append(opts, chore.WithTitle(title))
			case 1:
				opts = append(opts, chore.WithDescription(description))
			case 2:
				opts = append(opts, chore.WithShedule(shedule))
			case 3:
				opts = append(opts, chore.WithComment(comment))
			case 4:
				opts = append(opts, chore.WithAuthor(author))
			}
		}
		ch, err := chore.NewChore(opts...)
		if err != nil {
			return fmt.Errorf("failed to create chore")
		}

		switch c.Bool(flags.GLOBAL_CLI) {
		case false:
			if err := ch.Edit(); err != nil {
				return fmt.Errorf("failed to edit chore: %v", err)
			}
			fmt.Fprintf(os.Stderr, "%v\n", ch.String())
		case true:
		}

		if err := ch.Validate(); err != nil {
			return fmt.Errorf("new chore validation failed: %v", err)
		}
		store, err := js.New(actx.Config().StoragePath)
		if err != nil {
			return fmt.Errorf("failed to open persistent storage")
		}
		if err := store.Create(ch); err != nil {
			return fmt.Errorf("failed to create new chore entry in storage: %v", err)
		}
		return nil
	}
}
