package cmd

import (
	"cz/internal"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cz",
	Short: "cz - A simple and customizable commit message tool",
	Long: `cz helps developers write structured commit messages.
It follows conventional commit guidelines, ensuring consistency and clarity in commit history.`,
	Run: func(cmd *cobra.Command, args []string) {
		commitType := internal.InputCommitType()
		scope := internal.InputScope()
		message := internal.InputMessage()
		body := internal.InputBody()

		data := internal.CommitMessageData{
			Type:    commitType,
			Scope:   scope,
			Message: message,
			Body:    body,
		}

		commitMessage := internal.CompileCommitMessage(data)
		internal.GitCommit(commitMessage)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
