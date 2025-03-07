package cmd

import (
	"bytes"
	"cz/internal"
	"errors"
	"fmt"
	"os"
	"text/template"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cz",
	Short: "cz - A simple and customizable commit message tool",
	Long: `cz helps developers write structured commit messages.
It follows conventional commit guidelines, ensuring consistency and clarity in commit history.`,
	Run: func(cmd *cobra.Command, args []string) {
		typeTemplate := &promptui.SelectTemplates{
			Label:    "{{ . }}?",
			Active:   "{{ .Label | cyan }}",
			Inactive: "{{ .Label }}",
		}

		typePrompt := promptui.Select{
			Label:     "Something",
			Items:     internal.COMMIT_TYPES,
			Templates: typeTemplate,
			Size:      3,
		}

		index, _, err := typePrompt.Run()
		internal.AbortOnError(err, "Failed to select commit type. Please try again.")

		commitType := internal.COMMIT_TYPES[index].Type

		scopePrompt := promptui.Prompt{
			Label: "Scope (e.g., auth, db, api, ui, cli) - leave empty for none",
		}

		scope, err := scopePrompt.Run()
		internal.AbortOnError(err, "Failed to get a commit message. Please try again.")

		messagePrompt := promptui.Prompt{
			Label: "Enter short commit description",
			Validate: func(s string) error {
				if len(s) < 10 {
					return errors.New("please enter at least 10 characters")
				}

				return nil
			},
		}

		message, err := messagePrompt.Run()
		internal.AbortOnError(err, "Failed to get the commit message. Please try again.")

		bodyPrompt := promptui.Prompt{
			Label:     "Enter detailed commit message (optional)",
			IsVimMode: true,
		}

		body, err := bodyPrompt.Run()
		internal.AbortOnError(err, "Failed to get the commit body. Please try again.")

		commitTemplate, err := template.New("commit-message").Parse(internal.DEFAULT_COMMIT_FORMAT)
		internal.AbortOnError(err, "Failed to compile commit message. Please try again.")

		var commitMessageBuf bytes.Buffer
		data := internal.CommitMessageData{
			Type:    commitType,
			Scope:   scope,
			Message: message,
			Body:    body,
		}

		err = commitTemplate.Execute(&commitMessageBuf, data)
		internal.AbortOnError(err, "Failed to execute commit message template. Please try again.")

		commitMessage := commitMessageBuf.String()
		internal.GitCommit(commitMessage)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
