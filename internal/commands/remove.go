package commands

import (
	"fmt"
	"github.com/sokinpui/worktree/internal/git"

	"github.com/spf13/cobra"
)

var forceRemove bool

var removeCmd = &cobra.Command{
	Use:   "remove <worktree-path>",
	Aliases: []string{"rm"},
	Short: "Remove a worktree and its associated local branch",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runRemove(args[0], forceRemove)
	},
}

func init() {
	removeCmd.Flags().BoolVarP(&forceRemove, "force", "f", false, "Force removal of worktree and branch")
}

func runRemove(path string, force bool) error {
	worktrees, err := git.GetWorktrees()
	if err != nil {
		return err
	}

	var targetBranch string
	for _, wt := range worktrees {
		if wt.Path == path {
			targetBranch = wt.Branch
			break
		}
	}

	if targetBranch == "" {
		return fmt.Errorf("worktree at path '%s' not found", path)
	}

	if err := git.RemoveWorktree(path, force); err != nil {
		return fmt.Errorf("failed to remove worktree: %w", err)
	}

	if err := git.DeleteBranch(targetBranch, force); err != nil {
		return fmt.Errorf("failed to delete branch '%s': %w", targetBranch, err)
	}

	return nil
}
