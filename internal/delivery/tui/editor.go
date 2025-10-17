package tui

import (
	"fmt"

	"github.com/Galdoba/choretracker/internal/constants"
	"github.com/Galdoba/choretracker/internal/core/dto"
	"github.com/Galdoba/choretracker/pkg/cronexpr"
	"github.com/charmbracelet/huh"
)

type editor struct {
}

var e = &editor{}

func (e *editor) Edit(content map[string]string) (map[string]string, error) {
	title := content[constants.Fld_Title]
	desc := content[constants.Fld_Descr]
	schedule := content[constants.Fld_Schedule]
	comment := content[constants.Fld_Comment]
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("chore title:").
				// Description(fmt.Sprintf("ID: %v", ch.ID)).
				Validate(validateName).
				Value(&title),
			huh.NewText().
				Title("description").
				Value(&desc).
				WithWidth(40).
				WithHeight(5),
			huh.NewInput().
				Title("cron shedule").
				Description("crontab expression: mm hh dom mon dow").
				Validate(validateShedule).
				Value(&schedule),
			huh.NewText().
				Title("comments").
				Value(&comment).
				WithWidth(40).
				WithHeight(5),
		),
	)
	if err := form.Run(); err != nil {
		return content, fmt.Errorf("failed to run chore editor form: %v", err)
	}
	content[constants.Fld_Title] = title
	content[constants.Fld_Descr] = desc
	content[constants.Fld_Schedule] = schedule
	content[constants.Fld_Comment] = comment
	return content, nil

}

func validateShedule(s string) error {
	_, err := cronexpr.Parse(s)
	if err != nil {
		return err
	}
	return nil

}

func validateName(s string) error {
	if s == "" {
		return fmt.Errorf("this fileld must not be empty")
	}
	return nil

}

func EditRequest(req *dto.ToServiceRequest) error {
	content, err := e.Edit(req.Fields.Content())
	if err != nil {
		return fmt.Errorf("failed to edit creation request: %v", err)
	}
	req.InjectContent(content)
	return nil
}

// func EditUpdateRequest(ur dto.UpdateRequest) (dto.UpdateRequest, error) {
// 	content := ur.Content()
// 	if err := e.Edit(&content); err != nil {
// 		return ur, fmt.Errorf("failed to edit creation request: %v", err)
// 	}
// 	return dto.UpdateRequest{
// 		ChoreIdentity: ur.ChoreIdentity,
// 		ChoreContent:  content,
// 	}, nil
// }
