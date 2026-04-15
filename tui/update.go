package tui

import tea "github.com/charmbracelet/bubbletea"

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		return m.handleWindowSize(msg)
	case tea.KeyMsg:
		return m.handleKeyInput(msg)
	}

	return m, nil
}

func (m Model) handleWindowSize(msg tea.WindowSizeMsg) (tea.Model, tea.Cmd) {
	m.width = msg.Width
	m.height = msg.Height
	return m, nil
}

func (m Model) handleKeyInput(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	key := msg.String()

	if key == "ctrl+c" || key == "q" {
		return m, tea.Quit
	}

	if key == "up" || key == "k" {
		return m.moveCursorUp(), nil
	}

	if key == "down" || key == "j" {
		return m.moveCursorDown(), nil
	}

	if key == "enter" || key == " " {
		return m.toggleSelection(), nil
	}

	return m, nil
}

func (m Model) moveCursorUp() Model {
	if m.cursor > 0 {
		m.cursor--
	}
	return m
}

func (m Model) moveCursorDown() Model {
	if m.cursor < len(m.worktrees)-1 {
		m.cursor++
	}
	return m
}

func (m Model) toggleSelection() Model {
	if _, ok := m.selected[m.cursor]; ok {
		delete(m.selected, m.cursor)
		return m
	}
	m.selected[m.cursor] = struct{}{}
	return m
}
