git worktree workflow

```
my-project/
├── .bare/        (Hidden: The actual Git history/metadata)
├── .git          (File: Points to .bare)
├── main/         (Folder: Your "main" branch code)
└── feat-logic/   (Folder: Your "feature" branch code)
```

We create the project like this:

```bash
mkdir my-project && cd my-project

git clone --bare <url> .bare

echo "gitdir: ./.bare" > .git

git config remote.origin.fetch "+refs/heads/*:refs/remotes/origin/*"

git fetch --all
```

The `main/master`

```bash
git worktree add main main
```

We create a new branch `feat-logic` and add it as a worktree:

```bash
git worktree add feat-logic -b feat-logic origin/feat-logic
```

or

```
git worktree add -b feat-logic feat-logic main
```

The template is:

```bash
# Usage: git worktree add -b <new-branch-name> <path> <base-point>
```
