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
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		commitType := internal.COMMIT_TYPES[index].Type

		scopePrompt := promptui.Prompt{
			Label: "Scope (e.g., auth, db, api, ui, cli) - leave empty for none",
		}

		scope, err := scopePrompt.Run()
		if err != nil {
			fmt.Printf("error: %v\n", err)
			return
		}

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
		if err != nil {
			fmt.Printf("error: %v\n", err)
			return
		}

		bodyPrompt := promptui.Prompt{
			Label:     "Enter detailed commit message (optional)",
			IsVimMode: true,
		}

		body, err := bodyPrompt.Run()
		if err != nil {
			fmt.Printf("error: %v\n", err)
			return
		}

		commitTemplate, err := template.New("commit-message").Parse(internal.DEFAULT_COMMIT_FORMAT)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}

		var commitMessageBuf bytes.Buffer
		data := internal.CommitMessageData{
			Type:    commitType,
			Scope:   scope,
			Message: message,
			Body:    body,
		}

		err = commitTemplate.Execute(&commitMessageBuf, data)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}

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
