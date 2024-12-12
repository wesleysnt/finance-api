package commands

import (
	"github.com/spf13/cobra"
)

var gobaseCommand = &cobra.Command{
	Use:   "gobase",
	Short: "gobase related commands",
}

func init() {
	rootCmd.AddCommand(gobaseCommand)
}
