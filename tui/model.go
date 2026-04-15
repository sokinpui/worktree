package tui

import (
	"worktree-cli/internal/git"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	worktrees []string
	cursor    int
	selected  map[int]struct{}
	width     int
	height    int
	err       error
}

func New() Model {
	worktrees, err := git.GetWorktrees()
	if err != nil {
		worktrees = []string{}
	}

	return Model{
		worktrees: worktrees,
		selected:  make(map[int]struct{}),
		err:       err,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
