package internal

import (
	"bytes"
	"text/template"
)

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

type CommitMessageData struct {
	Type    string
	Scope   string
	Message string
	Body    string
}

const DEFAULT_COMMIT_FORMAT = `{{ .Type }}{{ if .Scope}}({{ .Scope }}){{end}}: {{ .Message }}{{ if .Body}}

{{ .Body}}{{end}}`

// This function generates a commit message by parsing and executing a predefined template with the provided data.
// It uses the CommitMessageData struct to fill in the template. If parsing or execution fails, the function aborts
// the program and shows an error message. It returns the generated commit message as a string.
func CompileCommitMessage(data CommitMessageData, format string) string {
	commitTemplate, err := template.New("commit-message").Parse(format)
	AbortOnError(err, "failed to compile commit message. Please try again.")

	var commitMessageBuf bytes.Buffer
	err = commitTemplate.Execute(&commitMessageBuf, data)
	AbortOnError(err, "failed to execute commit message template. Please try again.")

	return commitMessageBuf.String()
}
