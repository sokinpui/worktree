package main

import (
	"flag"
	"fmt"
	"os"
	"worktree-cli/internal/commands"
	"worktree-cli/tui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	if len(os.Args) < 2 {
		runTUI()
		return
	}

	switch os.Args[1] {
	case "clone":
		handleClone()
	default:
		runTUI()
	}
}

func handleClone() {
	fs := flag.NewFlagSet("clone", flag.ExitOnError)
	dirPtr := fs.String("d", "", "target directory name")

	if err := fs.Parse(os.Args[2:]); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing flags: %v\n", err)
		os.Exit(1)
	}

	args := fs.Args()
	if len(args) < 1 {
		fmt.Println("Usage: worktree clone <repo-url> [-d <directory>]")
		os.Exit(1)
	}

	repoURL := args[0]
	if err := commands.Clone(repoURL, *dirPtr); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
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
