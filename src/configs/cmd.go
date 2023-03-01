package configs

import (
	"lectronic/src/databases/orm"

	"github.com/spf13/cobra"
)

var initCommand = cobra.Command{
	Short: "letronic backend golang",
	Long:  "letronic backend golang with gorilla/mux & gorm",
}

func init() {
	initCommand.AddCommand(ServeCmd)
	initCommand.AddCommand(orm.MigrateCmd)
	initCommand.AddCommand(orm.SeedCmd)
}

func Run(args []string) error {

	initCommand.SetArgs(args)

	return initCommand.Execute()
}
