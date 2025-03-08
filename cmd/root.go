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

		isGitInit := internal.IsGitInitialized()
		if !isGitInit {
			internal.Error("This is not a Git repository. Run this command inside a Git project.")
			return
		}

		retry, _ := cmd.Flags().GetBool("retry")

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

		commitType, _ := cmd.Flags().GetString("type")
		if commitType == "" {
			commitType = internal.InputCommitType()
		}

		scope, _ := cmd.Flags().GetString("scope")
		if scope == "" {
			scope = internal.InputScope()
		}

		message, _ := cmd.Flags().GetString("message")
		if message == "" {
			message = internal.InputMessage()
		}

		body, _ := cmd.Flags().GetString("body")
		if body == "" {
			body = internal.InputBody()
		}

		data := internal.CommitMessageData{
			Type:    commitType,
			Scope:   scope,
			Message: message,
			Body:    body,
		}

		commitMessage := internal.CompileCommitMessage(data)

		confirmCommit := internal.InputConfirm(fmt.Sprintf("Commit message -> %s?", commitMessage))
		if !confirmCommit {
			internal.Error("commit cancelled")
			return
		}

		internal.GitCommit(commitMessage)
		cache.SetPrevCommit(commitMessage)
		internal.Success("Your changes has been committed successfully. ")
	},
}

func init() {
	rootCmd.Flags().Bool("retry", false, "Reuse the previous commit message and retry the commit process.")
	rootCmd.Flags().StringP("type", "t", "", "Specify the commit type directly and skip the prompt (e.g., cz --type feat)")
	rootCmd.Flags().StringP("scope", "s", "", "Specify the commit scope directly and skip the prompt (e.g., cz --scope auth)")
	rootCmd.Flags().StringP("message", "m", "", "Specify the commit message directly and skip the prompt (e.g., cz --message 'Fix login bug')")
	rootCmd.Flags().StringP("body", "b", "", "Specify the commit body directly and skip the prompt (e.g., cz --body 'Fixed the issue causing session timeout')")

	cache.Init()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
