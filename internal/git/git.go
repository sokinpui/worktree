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

func PathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
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

func GetProjectRoot() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--git-common-dir")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	gitDir := strings.TrimSpace(string(out))
	absGitDir, err := filepath.Abs(gitDir)
	if err != nil {
		return "", err
	}
	return filepath.Dir(absGitDir), nil
}

func GetCurrentBranch() (string, error) {
	cmd := exec.Command("git", "branch", "--show-current")
	out, _ := cmd.Output()
	branch := strings.TrimSpace(string(out))
	if branch != "" {
		return branch, nil
	}

	cmd = exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

func BranchExists(name string) bool {
	cmd := exec.Command("git", "show-ref", "--verify", "--quiet", "refs/heads/"+name)
	return cmd.Run() == nil
}

func RemoteBranchExists(name string) bool {
	cmd := exec.Command("git", "show-ref", "--verify", "--quiet", "refs/remotes/"+name)
	return cmd.Run() == nil
}

func AddWorktree(path, branch, base string, isNew bool) error {
	args := []string{"worktree", "add"}
	if isNew {
		args = append(args, "-b", branch, path, base)
	} else {
		args = append(args, path, branch)
	}
	return Run(".", args...)
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
