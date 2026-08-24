package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

// Populated at build time via -ldflags by GoReleaser. The defaults are what
// you get from a plain `go build`.
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Print the version of cz.",
	Long:    "Prints the version, commit and build date of this cz binary, along with the Go runtime it was built for.",
	Aliases: []string{"v"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("cz %s\n", version)
		fmt.Printf("  commit:  %s\n", commit)
		fmt.Printf("  built:   %s\n", date)
		fmt.Printf("  runtime: %s %s/%s\n", runtime.Version(), runtime.GOOS, runtime.GOARCH)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.Version = version
}
