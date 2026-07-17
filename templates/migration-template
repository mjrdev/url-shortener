package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func init() {
	AllMigrations = append(AllMigrations, &gormigrate.Migration{
		ID: "---",
		Migrate: func(tx *gorm.DB) error {
			type T struct {
				gorm.Model
			}
			return tx.AutoMigrate(&T{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("t")
		},
	})
}
