package git

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Worktree struct {
	Path   string
	Branch string
}

type RemoteBranch struct {
	Name       string
	IsSelected bool
}

func Run(dir string, args ...string) error {
	cmd := exec.Command("git", args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func GetWorktrees() ([]Worktree, error) {
	cmd := exec.Command("git", "worktree", "list", "--porcelain")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var worktrees []Worktree
	lines := strings.Split(string(output), "\n")
	var current Worktree

	for _, line := range lines {
		if strings.HasPrefix(line, "worktree ") {
			current.Path = filepath.Base(strings.TrimPrefix(line, "worktree "))
			continue
		}
		if strings.HasPrefix(line, "branch ") {
			current.Branch = strings.TrimPrefix(line, "branch refs/heads/")
			worktrees = append(worktrees, current)
			current = Worktree{}
		}
	}
	return worktrees, nil
}

func GetRemoteBranches() ([]RemoteBranch, error) {
	cmd := exec.Command("git", "branch", "-r")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var branches []RemoteBranch
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			continue
		}
		if strings.Contains(trimmed, "->") {
			continue
		}
		branches = append(branches, RemoteBranch{
			Name: trimmed,
		})
	}

	return branches, nil
}

func RemoveWorktree(path string, force bool) error {
	args := []string{"worktree", "remove", path}
	if force {
		args = append(args, "--force")
	}
	return Run(".", args...)
}

func DeleteBranch(branch string, force bool) error {
	flag := "-d"
	if force {
		flag = "-D"
	}
	return Run(".", "branch", flag, branch)
}
