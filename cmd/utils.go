package cmd

type CommitType struct {
	Type  string
	Label string
}

var COMMIT_TYPES = []CommitType{
	{Type: "feat", Label: "feat: A new feature"},
	{Type: "fix", Label: "fix: A bug fix"},
	{Type: "chore", Label: "chore: Routine tasks, build process, or maintenance changes"},
	{Type: "docs", Label: "docs: Documentation updates (e.g., README, API docs)"},
	{Type: "style", Label: "style: Code style changes (formatting, missing semi-colons, etc.)"},
	{Type: "refactor", Label: "refactor: Code changes that improve structure without adding features or fixing bugs"},
	{Type: "perf", Label: "perf: Performance improvements"},
	{Type: "test", Label: "test: Adding or updating tests"},
	{Type: "build", Label: "build: Changes affecting the build system or dependencies"},
	{Type: "ci", Label: "ci: Changes related to CI/CD configuration or scripts"},
	{Type: "revert", Label: "revert: Reverting a previous commit"},
	{Type: "wip", Label: "wip: Work in progress, not a final commit"},
	{Type: "security", Label: "security: Fixing security vulnerabilities"},
}
