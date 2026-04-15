package commands

import (
	"fmt"
	"github.com/sokinpui/worktree-cli/internal/git"
	"github.com/sokinpui/worktree-cli/internal/tui"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List local worktrees and remote branches",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runList()
	},
}

func runList() error {
	wt, err := git.GetWorktrees()
	if err != nil {
		return err
	}

	rb, err := git.GetRemoteBranches()
	if err != nil {
		return err
	}

	fmt.Print(tui.RenderList(wt, rb))
	return nil
}
