package chore

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

func (ch *Chore) Edit() error {
	ce := newEditModel(ch)
	editor, err := tea.NewProgram(ce).Run()
	if err != nil {
		return err
	}
	updatedEditor := editor.(choreEditor)
	ch = updatedEditor.chore
	return nil
}

func summary(ch *Chore) func() string {
	return func() string {
		s := ""
		s += fmt.Sprintf("chore      : %s", ch.Name) + "\n"
		s += fmt.Sprintf("by: %s", ch.Author) + "\n"
		s += fmt.Sprintf("description: %s", ch.Description) + "\n"
		s += fmt.Sprintf("shedule: %s", strings.Join(ch.CronSchedule, " "))
		return s
	}

}

type choreEditor struct {
	chore   *Chore
	summary string
	done    bool
	lg      *lipgloss.Renderer
	form    *huh.Form
}

func newEditModel(ch *Chore) choreEditor {
	ce := choreEditor{
		chore: ch,
		done:  false,
		lg:    lipgloss.DefaultRenderer(),
	}
	// id := ch.ID
	// name := ch.Name
	// descr := ch.Description
	// shedule := strings.Join(ch.CronSchedule, " ")

	ce.form = huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Title("name").Value(&ce.chore.Name)),
		huh.NewGroup(
			huh.NewInput().Title("description").Value(&ce.chore.Description)),
		huh.NewGroup(
			huh.NewInput().Title("shedule").Value(&ce.summary)),
	).WithLayout(huh.LayoutColumns(1))
	return ce
}

func (ce choreEditor) Init() tea.Cmd {
	return ce.form.GetFocusedField().Focus()
}

func (ce choreEditor) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return ce, tea.Quit
		}
	}
	var cmds []tea.Cmd

	// Process the form
	form, cmd := ce.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		ce.form = f
		cmds = append(cmds, cmd)
	}

	if ce.form.State == huh.StateCompleted {
		// Quit when the form is done.
		cmds = append(cmds, tea.Quit)
	}
	return ce, tea.Batch(cmds...)
}

func (ce choreEditor) View() string {

	if ce.done {
		return ""
	}
	out := summary(ce.chore)() + "\n"
	out += ce.form.View()

	return out
}
