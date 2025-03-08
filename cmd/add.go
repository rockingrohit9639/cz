package cmd

import (
	"cz/internal"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add all changed files to the staging area.",
	Long:    "Adds all modified, deleted, and new files to the staging area, preparing them for commit (equivalent to 'git add .')",
	Aliases: []string{"a"},
	Run: func(cmd *cobra.Command, args []string) {
		internal.StageAll()
		internal.Success("staged all changes")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
