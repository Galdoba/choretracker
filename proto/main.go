package main

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Модель для структуры Chore
type Chore struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	Author           string    `json:"author"`
	Opened           time.Time `json:"opened"`
	NextNotification time.Time `json:"next_notification"`
	CronSchedule     string    `json:"schedule"`
}

// Модель приложения Bubble Tea
type model struct {
	inputs    []textinput.Model // Поля ввода
	chore     Chore             // Данные формы
	cursor    int               // Текущее активное поле
	submitted bool              // Флаг отправки формы
}

// Стили для улучшенного внешнего вида
var (
	titleStyle   = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("62")).MarginBottom(1)
	cursorStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))
	fieldStyle   = lipgloss.NewStyle().Width(60).MarginBottom(1)
	helpStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).MarginTop(2)
	successStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("10")).MarginTop(1)
)

// Инициализация модели
func initialModel() model {
	m := model{
		inputs: make([]textinput.Model, 4),
	}

	// Поле "Название задачи"
	m.inputs[0] = textinput.New()
	m.inputs[0].Placeholder = "Введите название задачи"
	m.inputs[0].Focus()
	m.inputs[0].CharLimit = 100
	m.inputs[0].Width = 50
	m.inputs[0].Prompt = "Название задачи: "
	m.inputs[0].PromptStyle = cursorStyle

	// Поле "Описание"
	m.inputs[1] = textinput.New()
	m.inputs[1].Placeholder = "Подробное описание задачи"
	m.inputs[1].CharLimit = 500
	m.inputs[1].Width = 50
	m.inputs[1].Prompt = "Описание: "
	m.inputs[1].PromptStyle = fieldStyle

	// Поле "Автор"
	m.inputs[2] = textinput.New()
	m.inputs[2].Placeholder = "Ваше имя"
	m.inputs[2].CharLimit = 50
	m.inputs[2].Width = 50
	m.inputs[2].Prompt = "Автор: "
	m.inputs[2].PromptStyle = fieldStyle

	// Поле "Cron расписание"
	m.inputs[3] = textinput.New()
	m.inputs[3].Placeholder = "Например: 0 9 * * * (каждый день в 9 утра)"
	m.inputs[3].CharLimit = 50
	m.inputs[3].Width = 50
	m.inputs[3].Prompt = "Cron расписание: "
	m.inputs[3].PromptStyle = fieldStyle

	return m
}

// Инициализация приложения :cite[2]:cite[7]
func (m model) Init() tea.Cmd {
	return textinput.Blink
}

// Обработка сообщений и обновление состояния :cite[2]:cite[9]
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit

		case "tab", "enter":
			// Переход к следующему полю или отправка формы
			if m.cursor == len(m.inputs)-1 {
				// Последнее поле - сохраняем данные
				return m.submitForm()
			}
			m.cursor++
			return m.updateInputsFocus()

		case "shift+tab":
			// Переход к предыдущему полю
			if m.cursor > 0 {
				m.cursor--
			}
			return m.updateInputsFocus()

		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
			return m.updateInputsFocus()

		case "down":
			if m.cursor < len(m.inputs)-1 {
				m.cursor++
			}
			return m.updateInputsFocus()
		}
	}

	// Обновление активного поля ввода
	cmd := m.updateActiveInput(msg)
	return m, cmd
}

// Обновление фокуса полей ввода
func (m *model) updateInputsFocus() (tea.Model, tea.Cmd) {
	for i := range m.inputs {
		m.inputs[i].Blur()
	}
	m.inputs[m.cursor].Focus()

	var cmds []tea.Cmd
	for i := range m.inputs {
		var cmd tea.Cmd
		m.inputs[i], cmd = m.inputs[i].Update(nil)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

// Обновление активного поля ввода
func (m *model) updateActiveInput(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	m.inputs[m.cursor], cmd = m.inputs[m.cursor].Update(msg)
	return cmd
}

// Отправка формы и сохранение данных
func (m *model) submitForm() (tea.Model, tea.Cmd) {
	// Сохранение данных в структуру Chore
	m.chore.Name = m.inputs[0].Value()
	m.chore.Description = m.inputs[1].Value()
	m.chore.Author = m.inputs[2].Value()
	m.chore.CronSchedule = m.inputs[3].Value()

	// Установка временных меток
	m.chore.Opened = time.Now()
	m.chore.NextNotification = calculateNextNotification(m.chore.CronSchedule)

	m.submitted = true
	return m, tea.Quit
}

// Отображение интерфейса :cite[2]:cite[7]
func (m model) View() string {
	if m.submitted {
		return m.successView()
	}
	return m.formView()
}

// Отображение формы
func (m model) formView() string {
	s := titleStyle.Render("📝 Создание новой задачи") + "\n\n"

	for i := range m.inputs {
		s += fieldStyle.Render(m.inputs[i].View())
		if i < len(m.inputs)-1 {
			s += "\n"
		}
	}

	s += helpStyle.Render("\n\nНавигация: ↑/↓ или Tab/Shift+Tab • Ввод: Enter • Выход: Esc/Ctrl+C")

	return s
}

// Отображение успешного заполнения
func (m model) successView() string {
	s := successStyle.Render("✅ Задача успешно создана!") + "\n\n"
	s += fmt.Sprintf("Название: %s\n", m.chore.Name)
	s += fmt.Sprintf("Описание: %s\n", m.chore.Description)
	s += fmt.Sprintf("Автор: %s\n", m.chore.Author)
	s += fmt.Sprintf("Cron расписание: %s\n", m.chore.CronSchedule)
	s += fmt.Sprintf("Время создания: %s\n", m.chore.Opened.Format("2006-01-02 15:04:05"))
	s += helpStyle.Render("\nНажмите любую клавишу для выхода...")
	return s
}

// Расчет следующего уведомления (заглушка)
func calculateNextNotification(cron string) time.Time {
	// Здесь должна быть реальная логика парсинга cron выражения
	// Пока возвращаем время через 24 часа
	return time.Now().Add(24 * time.Hour)
}

// Основная функция :cite[2]:cite[10]
func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Ошибка запуска приложения: %v", err)
		os.Exit(1)
	}
}
