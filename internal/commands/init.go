package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"github.com/sokinpui/worktree/internal/git"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init [directory]",
	Short: "Initialize a new empty repository with the bare worktree workflow",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		path := "."
		if len(args) > 0 {
			path = args[0]
		}
		return runInit(path)
	},
}

func runInit(path string) error {
	if path != "." {
		if err := os.MkdirAll(path, 0755); err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}
	}

	if git.PathExists(filepath.Join(path, ".bare")) {
		return fmt.Errorf("repository already initialized in '%s'", path)
	}

	if err := git.InitBare(path); err != nil {
		return fmt.Errorf("failed to init bare repo: %w", err)
	}

	if err := git.CreateDotGit(path); err != nil {
		return fmt.Errorf("failed to create .git file: %w", err)
	}

	if err := git.Run(path, "config", "remote.origin.fetch", "+refs/heads/*:refs/remotes/origin/*"); err != nil {
		fmt.Printf("Warning: failed to set fetch config: %v\n", err)
	}

	fmt.Println("Creating 'main' worktree...")
	if err := git.AddOrphanWorktree(path, "main", "main"); err != nil {
		return fmt.Errorf("failed to create main worktree: %w", err)
	}

	fmt.Printf("Successfully initialized project in '%s'\n", path)
	fmt.Println("\nNext steps:")
	fmt.Printf("  1. cd %s\n", filepath.Join(path, "main"))
	fmt.Println("  2. git remote add origin <url>")
	fmt.Println("  3. git add . && git commit -m \"initial commit\"")
	fmt.Println("  4. git push -u origin main")

	return nil
}
