package commands

import (
	"fmt"
	"github.com/sokinpui/worktree-cli/internal/git"
	"github.com/sokinpui/worktree-cli/internal/tui"

	"github.com/spf13/cobra"
)

var (
	listLocal  bool
	listRemote bool
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List local worktrees and remote branches",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runList(listLocal, listRemote)
	},
}

func init() {
	listCmd.Flags().BoolVarP(&listLocal, "local", "l", false, "List only local worktrees")
	listCmd.Flags().BoolVarP(&listRemote, "remote", "r", false, "List only remote branches")
}

func runList(showLocal, showRemote bool) error {
	if !showLocal && !showRemote {
		showLocal = true
		showRemote = true
	}

	var wt []git.Worktree
	var rb []git.RemoteBranch
	var err error

	if showLocal {
		wt, err = git.GetWorktrees()
		if err != nil {
			return err
		}
	}

	if showRemote {
		rb, err = git.GetRemoteBranches()
		if err != nil {
			return err
		}
	}

	fmt.Print(tui.RenderList(wt, rb, showLocal, showRemote))
	return nil
}
