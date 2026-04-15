package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"worktree-cli/internal/git"
)

func Clone(repoURL, targetDir string) error {
	if targetDir == "" {
		targetDir = deriveDirName(repoURL)
	}

	if isDirectoryExists(targetDir) {
		return fmt.Errorf("directory '%s' already exists", targetDir)
	}

	return executeCloneWorkflow(repoURL, targetDir)
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

func isDirectoryExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
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
