package tui

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
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("chore title:").
				// Description(fmt.Sprintf("ID: %v", ch.ID)).
				Validate(validateName).
				Value(content.Title),
			huh.NewText().
				Title("description").
				Value(content.Description).
				WithWidth(40).
				WithHeight(5),
			huh.NewInput().
				Title("cron shedule").
				Description("crontab expression: mm hh dom mon dow").
				Validate(validateShedule).
				Value(content.Schedule),
			huh.NewText().
				Title("comments").
				Value(content.Comment).
				WithWidth(40).
				WithHeight(5),
		),
	)
	if err := form.Run(); err != nil {
		return fmt.Errorf("failed to run chore editor form: %v", err)
	}
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
	content := cr.Content()
	if err := e.Edit(&content); err != nil {
		return cr, fmt.Errorf("failed to edit creation request: %v", err)
	}
	return dto.CreateRequest{
		ChoreContent: content,
	}, nil
}

func EditUpdateRequest(ur dto.UpdateRequest) (dto.UpdateRequest, error) {
	content := ur.Content()
	if err := e.Edit(&content); err != nil {
		return ur, fmt.Errorf("failed to edit creation request: %v", err)
	}
	return dto.UpdateRequest{
		ChoreIdentity: ur.ChoreIdentity,
		ChoreContent:  content,
	}, nil
}
