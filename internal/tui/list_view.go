package tui

import (
	"fmt"
	"strings"
	"github.com/sokinpui/worktree-cli/internal/git"
)

func RenderList(worktrees []git.Worktree, remotes []git.RemoteBranch) string {
	var b strings.Builder

	b.WriteString(HeaderStyle.Render("[Local branch]:") + "\n")
	if len(worktrees) == 0 {
		b.WriteString(MutedStyle.Render("  No local worktrees found") + "\n")
	}

	for _, wt := range worktrees {
		b.WriteString(fmt.Sprintf("  %-15s %s\n", wt.Path, MutedStyle.Render(wt.Branch)))
	}

	b.WriteString("\n")

	b.WriteString(HeaderStyle.Render("[Remote branch]:") + "\n")
	if len(remotes) == 0 {
		b.WriteString(MutedStyle.Render("  No remote branches found") + "\n")
	}

	for _, rb := range remotes {
		b.WriteString(fmt.Sprintf("  %s\n", rb.Name))
	}

	return b.String()
}
