package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cz",
	Short: "cz - A simple and customizable commit message tool",
	Long: `cz helps developers write structured commit messages.
It follows conventional commit guidelines, ensuring consistency and clarity in commit history.`,
	Run: func(cmd *cobra.Command, args []string) {
		templates := &promptui.SelectTemplates{
			Label:    "{{ . }}?",
			Active:   "{{ .Label | cyan }}",
			Inactive: "{{ .Label }}",
		}

		typePrompt := promptui.Select{
			Label:     "Something",
			Items:     COMMIT_TYPES,
			Templates: templates,
			Size:      3,
		}

		index, _, err := typePrompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		commitType := COMMIT_TYPES[index].Type

		scopePrompt := promptui.Prompt{
			Label: "Scope (e.g., auth, db, api, ui, cli) - leave empty for none",
		}

		scope, err := scopePrompt.Run()
		if err != nil {
			fmt.Printf("error: %v\n", err)
			return
		}

		shortDescPrompt := promptui.Prompt{
			Label: "Enter short commit description",
			Validate: func(s string) error {
				if len(s) < 10 {
					return errors.New("please enter at least 50 characters")
				}

				return nil
			},
		}

		shortDesc, err := shortDescPrompt.Run()
		if err != nil {
			fmt.Printf("error: %v\n", err)
			return
		}

		longDescPrompt := promptui.Prompt{
			Label:     "Enter detailed commit message (optional)",
			IsVimMode: true,
		}

		longDesc, err := longDescPrompt.Run()
		if err != nil {
			fmt.Printf("error: %v\n", err)
			return
		}

		var commitMessage string

		if scope != "" {
			commitMessage = fmt.Sprintf("%s(%s): %s\n\n%s", commitType, scope, shortDesc, longDesc)
		} else {
			commitMessage = fmt.Sprintf("%s: %s\n\n%s", commitType, shortDesc, longDesc)
		}

		commitCommand := exec.Command("git", "commit", "-m", commitMessage)
		err = commitCommand.Run()
		if err != nil {
			fmt.Printf("error: %v", err)
			return
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
