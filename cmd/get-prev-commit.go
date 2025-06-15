package cmd

import (
	"cz/cache"
	"fmt"

	"github.com/spf13/cobra"
)

var getPrevCommitCmd = &cobra.Command{
	Use:     "get-prev-commit",
	Short:   "Get the previous commit message from the cache",
	Long:    "Get the previous commit message that was saved in the cache. This is useful for retrying a commit with the last used message.",
	Aliases: []string{"gpc"},
	Run: func(cmd *cobra.Command, args []string) {
		prevCommitMsg := cache.GetPrevCommit()
		fmt.Println(prevCommitMsg)
	},
}

func init() {
	rootCmd.AddCommand(getPrevCommitCmd)
}
