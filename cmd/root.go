package cmd

import (
	"cz/cache"
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
		hasStagedChanges := internal.HasStagedChanges()
		if !hasStagedChanges {
			internal.Error("Whoa there! No staged changes found. Are you trying to commit air? Stage your changes first!")
			return
		}

		retry, err := cmd.Flags().GetBool("retry")
		if err != nil {
			internal.Warn("failed to get value for retry flag")
		}

		// Retry commit with the last commit message
		if retry {
			prevCommitMsg := cache.GetPrevCommit()
			if prevCommitMsg != "" {
				internal.GitCommit(prevCommitMsg)
				internal.Success("Your changes have been committed successfully.")
				return
			} else {
				internal.Warn("No previous commit found, continuing with new commit.")
			}
		}

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

		cache.SetPrevCommit(commitMessage)

		internal.Success("Your changes has been committed successfully. ")
	},
}

func init() {
	rootCmd.Flags().Bool("retry", false, "Reuse the previous commit message and retry the commit process.")
	cache.Init()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
