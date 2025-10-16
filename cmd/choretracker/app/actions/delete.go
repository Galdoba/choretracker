package actions

// func Delete(actx *appcontext.AppContext) cli.ActionFunc {
// 	return func(ctx context.Context, c *cli.Command) error {
// 		store, err := js.New(actx.Config().StoragePath)
// 		if err != nil {
// 			return fmt.Errorf("failed to open persistent storage")
// 		}

// 		ch, err := getChore(c, store)
// 		if err != nil {
// 			return fmt.Errorf("failed to get chore data: %v", err)
// 		}

// 		if err = store.Delete(ch.ID); err != nil {
// 			return fmt.Errorf("failed to delete chore: %v", err)
// 		}
// 		return nil
// 	}
// }
