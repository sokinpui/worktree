package commands

import (
	"fmt"
	"path/filepath"
	"worktree-cli/internal/git"

	"github.com/spf13/cobra"
)

var baseBranch string

var addCmd = &cobra.Command{
	Use:   "add <branch> [path]",
	Short: "Add a new worktree",
	Args:  cobra.RangeArgs(1, 2),
	RunE: func(cmd *cobra.Command, args []string) error {
		branch := args[0]
		path := branch
		if len(args) > 1 {
			path = args[1]
		}
		return runAdd(branch, path, baseBranch)
	},
}

func init() {
	addCmd.Flags().StringVarP(&baseBranch, "base", "b", "", "Base branch/commit for new branch (defaults to current branch)")
}

func runAdd(branch, path, base string) error {
	targetPath := path
	if !filepath.IsAbs(path) {
		root, err := git.GetProjectRoot()
		if err != nil {
			return fmt.Errorf("failed to find project root: %w", err)
		}
		targetPath = filepath.Join(root, path)
	}

	if git.PathExists(targetPath) {
		return fmt.Errorf("directory '%s' already exists", targetPath)
	}

	if git.BranchExists(branch) {
		return git.AddWorktree(targetPath, branch, "", false)
	}

	if git.RemoteBranchExists("origin/" + branch) {
		return git.AddWorktree(targetPath, branch, "", false)
	}

	if base == "" {
		var err error
		base, err = git.GetCurrentBranch()
		if err != nil || base == "" || base == "HEAD" {
			base = "main"
		}
	}

	return git.AddWorktree(targetPath, branch, base, true)
}
