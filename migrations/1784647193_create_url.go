package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	models "github.com/mjrdev/internal/model"
	"gorm.io/gorm"
)

func init() {
	AllMigrations = append(AllMigrations, &gormigrate.Migration{
		ID: "1784647193",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&models.Url{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("urls")
		},
	})
}
