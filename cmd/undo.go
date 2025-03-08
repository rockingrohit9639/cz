package cmd

import (
	"cz/internal"
	"fmt"

	"github.com/spf13/cobra"
)

var undoCmd = &cobra.Command{
	Use:   "undo",
	Short: "Undo the last commit made with cz",
	Long:  "Revert the last commit created using cz. This will reset the commit while keeping the changes staged, allowing you to make modifications and commit again if needed.",
	Run: func(cmd *cobra.Command, args []string) {
		confirm := internal.InputConfirm("Are you sure you want to undo last commit?")
		if !confirm {
			fmt.Print("Okay, as you with :)")
			return
		}

		internal.UndoLastCommit()
		internal.Success("last commit is reverted")
	},
}

func init() {
	rootCmd.AddCommand(undoCmd)
}
