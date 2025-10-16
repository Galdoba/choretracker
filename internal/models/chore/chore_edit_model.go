package chore

// func (ch *Chore) Edit() error {

// 	form := huh.NewForm(
// 		huh.NewGroup(
// 			huh.NewInput().
// 				Title("chore title:").
// 				Description(fmt.Sprintf("ID: %v", ch.ID)).
// 				Validate(validateName).
// 				Value(&ch.Title),
// 			huh.NewText().
// 				Title("description").
// 				Value(&ch.Description).
// 				WithWidth(40).
// 				WithHeight(5),
// 			huh.NewInput().
// 				Title("cron shedule").
// 				Description("crontab expression: mm hh dom mon dow").
// 				Validate(validateShedule).
// 				Value(&ch.CronSchedule),
// 			huh.NewText().
// 				Title("comments").
// 				Value(&ch.Comments).
// 				WithWidth(40).
// 				WithHeight(5),
// 		),
// 	)
// 	if err := form.Run(); err != nil {
// 		return fmt.Errorf("failed to run chore editor form: %v", err)
// 	}

// 	return ch.Update()
// }

// func validateShedule(s string) error {
// 	_, err := cronexpr.Parse(s)
// 	if err != nil {
// 		return err
// 	}
// 	return nil

// }

// func validateName(s string) error {
// 	if s == "" {
// 		return fmt.Errorf("this fileld must not be empty")
// 	}
// 	return nil

// }
