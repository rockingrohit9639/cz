package cmd

import (
	"cz/config"
	"cz/internal"

	"github.com/spf13/cobra"
)

var addFormatCmd = &cobra.Command{
	Use:     "add-format",
	Aliases: []string{"af"},
	Short:   "Add a new commit format",
	Long:    "Add a new commit format to the config. This format will be used to generate the commit message.",
	Run: func(cmd *cobra.Command, args []string) {
		name := internal.InputString("Name of the format")
		pattern := internal.InputString("Pattern of the format This should follow go template syntax.")

		err := config.AddFormat(name, pattern)
		if err != nil {
			internal.Error(err.Error())
			return
		}

		internal.Success("Format added successfully")
	},
}

func init() {
	rootCmd.AddCommand(addFormatCmd)
}
