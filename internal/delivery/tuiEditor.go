package delivery

import (
	"fmt"

	"github.com/Galdoba/choretracker/internal/core/dto"
	"github.com/Galdoba/choretracker/pkg/cronexpr"
	"github.com/charmbracelet/huh"
)

type editor struct {
}

var e = &editor{}

func (e *editor) Edit(content *dto.ChoreContent) error {
	title := ""
	if content.Title != nil {
		title = *content.Title
	}
	desc := ""
	if content.Description != nil {
		desc = *content.Description
	}
	sched := ""
	if content.Schedule != nil {
		sched = *content.Schedule
	}
	comment := ""
	if content.Comment != nil {
		comment = *content.Comment
	}
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
				Value(&sched),
			huh.NewText().
				Title("comments").
				Value(&comment).
				WithWidth(40).
				WithHeight(5),
		),
	)
	if err := form.Run(); err != nil {
		return fmt.Errorf("failed to run chore editor form: %v", err)
	}
	content.Title = &title
	content.Description = &desc
	content.Schedule = &sched
	content.Comment = &comment
	return nil

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

func EditCreateRequest(cr dto.CreateRequest) (dto.CreateRequest, error) {
	content := cr.ChoreContent
	if err := e.Edit(&content); err != nil {
		return cr, fmt.Errorf("failed to edit creation request: %v", err)
	}
	return dto.CreateRequest{
		ChoreContent: content,
	}, nil
}

func EditUpdateRequest(ur dto.UpdateRequest) (dto.UpdateRequest, error) {
	content := ur.ChoreContent
	if err := e.Edit(&content); err != nil {
		return ur, fmt.Errorf("failed to edit creation request: %v", err)
	}
	return dto.UpdateRequest{
		ChoreIdentity: ur.ChoreIdentity,
		ChoreContent:  content,
	}, nil
}
