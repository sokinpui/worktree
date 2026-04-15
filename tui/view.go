package tui

import (
	"fmt"
	"strings"
)

func (m Model) View() string {
	var b strings.Builder

	b.WriteString("Git Worktree Manager\n\n")

	if m.err != nil {
		return fmt.Sprintf("Error: %v\n\n(q to quit)", m.err)
	}

	if len(m.worktrees) == 0 {
		return "No worktrees found.\n\n(q to quit)"
	}

	for i, wt := range m.worktrees {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		b.WriteString(fmt.Sprintf("%s [%s] %s\n", cursor, checked, wt))
	}

	b.WriteString("\n(j/k: move, space: select, q: quit)\n")

	return b.String()
}
