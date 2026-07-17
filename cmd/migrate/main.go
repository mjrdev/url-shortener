package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/mjrdev/migrations"

	"github.com/mjrdev/internal/config"
)

func main() {
	rollback := flag.Bool("rollback", false, "faz rollback da última migration")
	flag.Parse()

	db := config.Db()

	m := gormigrate.New(db, gormigrate.DefaultOptions, migrations.AllMigrations)

	if *rollback {
		if err := m.RollbackLast(); err != nil {
			log.Fatalf("erro no rollback: %v", err)
		}
		fmt.Println("Rollback aplicado com sucesso!")
		return
	}

	if err := m.Migrate(); err != nil {
		log.Fatalf("erro ao migrar: %v", err)
	}
	fmt.Println("Migrations aplicadas com sucesso!")
}
