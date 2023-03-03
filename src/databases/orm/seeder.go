package orm

import (
	"fmt"
	"log"

	"lectronic/src/databases/orm/models"
	seeder "lectronic/src/databases/orm/seeders"

	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

type seederData struct {
	name  string
	model interface{}
	size  int
}

var SeedCmd = &cobra.Command{
	Use:   "seed",
	Short: "db seeder",
	RunE:  Seed,
}

var sedUp, sedown bool

func init() {
	SeedCmd.Flags().BoolVarP(&sedUp, "up", "u", true, "run seed up")
	SeedCmd.Flags().BoolVarP(&sedown, "down", "d", false, "run seed down")
}

func Seed(cmd *cobra.Command, args []string) error {
	var err error
	db, err := ConnectDB()
	if err != nil {
		return err
	}

	if sedown {
		err = seedDown(db)
		return err
	}

	if sedUp {
		err = seedUp(db)
	}

	return err
}

func seedUp(db *gorm.DB) error {
	var err error

	var seedModel = []seederData{
		{
			name:  "user",
			model: seeder.UserSeed,
			size:  cap(seeder.UserSeed),
		},
	}

	for _, data := range seedModel {
		log.Println("create seeding data for", data.name)
		err = db.CreateInBatches(data.model, data.size).Error
	}

	return err

}

func seedDown(db *gorm.DB) error {
	var err error

	var seedModel = []seederData{
		{
			name:  models.User{}.TableName(),
			model: models.User{},
		},
	}

	for _, data := range seedModel {
		log.Println("delete seeding data for", data.name)
		sql := fmt.Sprintf("DELETE FROM %v", data.name)
		err = db.Exec(sql).Error
	}

	return err
}
