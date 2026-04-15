package commands

import (
	"fmt"
	"os"
	"worktree-cli/internal/tui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "worktree",
	Short: "A CLI tool for managing Git worktrees with a bare repository workflow",
	Run: func(cmd *cobra.Command, args []string) {
		runTUI()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func runTUI() {
	program := tea.NewProgram(tui.New())
	if _, err := program.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(cloneCmd)
	rootCmd.AddCommand(listCmd)
}
