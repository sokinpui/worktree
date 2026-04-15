package tui

import "github.com/charmbracelet/lipgloss"

var (
	ColorPrimary   = lipgloss.AdaptiveColor{Light: "#7D56F4", Dark: "#BD93F9"}
	ColorSecondary = lipgloss.AdaptiveColor{Light: "#04B575", Dark: "#50FA7B"}
	ColorMuted     = lipgloss.AdaptiveColor{Light: "#9B9B9B", Dark: "#6272A4"}

	TitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorPrimary).
			MarginBottom(1)

	HeaderStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorSecondary).
			Border(lipgloss.NormalBorder(), false, false, true, false).
			Padding(0, 1)

	TableStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(ColorMuted).
			Padding(0, 1)

	MutedStyle = lipgloss.NewStyle().Foreground(ColorMuted)

	SelectedStyle = lipgloss.NewStyle().
			Foreground(ColorPrimary).
			Bold(true)

	ColumnStyle = lipgloss.NewStyle().
			Padding(1, 2)
)
