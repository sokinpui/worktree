package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"github.com/sokinpui/worktree-cli/internal/git"

	"github.com/spf13/cobra"
)

var targetDir string

var cloneCmd = &cobra.Command{
	Use:   "clone <repo-url>",
	Short: "Clone a repository using the bare worktree workflow",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runClone(args[0], targetDir)
	},
}

func init() {
	cloneCmd.Flags().StringVarP(&targetDir, "dir", "d", "", "Target directory name")
}

func runClone(repoURL, dir string) error {
	if dir == "" {
		dir = deriveDirName(repoURL)
	}

	if git.PathExists(dir) {
		return fmt.Errorf("directory '%s' already exists", dir)
	}

	return executeCloneWorkflow(repoURL, dir)
}

func deriveDirName(repoURL string) string {
	trimmed := strings.TrimSuffix(repoURL, "/")
	trimmed = strings.TrimSuffix(trimmed, ".git")
	
	if lastColon := strings.LastIndex(trimmed, ":"); lastColon != -1 {
		trimmed = trimmed[lastColon+1:]
	}
	
	parts := strings.Split(trimmed, "/")
	return parts[len(parts)-1]
}

func executeCloneWorkflow(repoURL, targetDir string) error {
	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return err
	}

	if err := git.Run(targetDir, "clone", "--bare", repoURL, ".bare"); err != nil {
		return err
	}

	dotGitPath := filepath.Join(targetDir, ".git")
	if err := os.WriteFile(dotGitPath, []byte("gitdir: ./.bare\n"), 0644); err != nil {
		return err
	}

	if err := git.Run(targetDir, "config", "remote.origin.fetch", "+refs/heads/*:refs/remotes/origin/*"); err != nil {
		return err
	}

	return git.Run(targetDir, "fetch", "--all")
}
