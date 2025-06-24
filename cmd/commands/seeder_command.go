package commands

import (
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"github.com/wesleysnt/finance-api/database/seeders"
)

var seederRunCmd = &cobra.Command{
	Use:   "db:seed",
	Short: "Run seeders",
	Run: func(cmd *cobra.Command, args []string) {
		err := seeders.Execute()

		if err != nil {
			color.Redf("Error when execute seeders %v \n", err)
			return
		}

		color.Greenln("Successfully to run db:seed")

	},
}

func init() {
	gobaseCommand.AddCommand(seederRunCmd)
}
