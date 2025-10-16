package actions

// func Update(actx *appcontext.AppContext) cli.ActionFunc {
// 	return func(ctx context.Context, c *cli.Command) error {
// 		store, err := js.New(actx.Config().StoragePath)
// 		if err != nil {
// 			return fmt.Errorf("failed to open persistent storage")
// 		}

// 		ch, err := getChore(c, store)
// 		if err != nil {
// 			return fmt.Errorf("failed to get chore data: %v", err)
// 		}

// 		if err = updateChore(c, ch); err != nil {
// 			return fmt.Errorf("failed to update chore data: %v", err)
// 		}

// 		if err = ch.Validate(); err != nil {
// 			return fmt.Errorf("failed to validate updated chore: %v", err)
// 		}
// 		if err = store.Update(ch); err != nil {
// 			return fmt.Errorf("failed to update chore: %v", err)
// 		}
// 		return nil
// 	}
// }

// func updateChore(c *cli.Command, ch *chore.Chore) error {
// 	injectFromFlags(c, ch)
// 	if !c.Bool(flags.GLOBAL_CLI) {
// 		if err := ch.Edit(); err != nil {
// 			return fmt.Errorf("failed to edit chore: %v", err)
// 		}
// 	}
// 	return nil
// }

// func injectFromFlags(c *cli.Command, ch *chore.Chore) {
// 	for i, val := range []string{
// 		c.String(flags.CHORE_TITLE),
// 		c.String(flags.CHORE_DESCRIPTION),
// 		c.String(flags.CHORE_SCHEDULE),
// 		c.String(flags.CHORE_COMMENT),
// 		c.String(flags.CHORE_AUTHOR),
// 	} {
// 		switch i {
// 		case 0:
// 			ch.Title = updateValue(ch.Title, val)
// 		case 1:
// 			ch.Description = updateValue(ch.Description, val)
// 		case 2:
// 			ch.CronSchedule = updateValue(ch.CronSchedule, val)
// 		case 3:
// 			ch.Comments = updateValue(ch.Comments, val)
// 		case 4:
// 			ch.Author = updateValue(ch.Author, val)
// 		}
// 	}
// }

// func updateValue(old, new string) string {
// 	if new == "" {
// 		return old
// 	}
// 	return new
// }
