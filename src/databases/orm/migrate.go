package orm

import (
	"lectronic/src/databases/orm/models"
	"log"

	"github.com/spf13/cobra"
)

var MigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "db migration",
	RunE:  dbMigrate,
}

var migUp, migDown bool

func init() {
	MigrateCmd.Flags().BoolVarP(&migUp, "up", "u", true, "run migration up")
	MigrateCmd.Flags().BoolVarP(&migDown, "down", "d", false, "run migration down")
}

func dbMigrate(cmd *cobra.Command, args []string) error {

	db, err := ConnectDB()
	if err != nil {
		return err
	}

	if migDown {
		log.Println("Migration down done")
		return db.Migrator().DropTable(&models.User{}, &models.Product{}, &models.Review{}, &models.Cart{})
	}

	if migUp {
		log.Println("Migration up done")
		return db.AutoMigrate(&models.User{}, &models.Product{}, &models.Review{}, &models.Cart{})
	}

	return nil
}
