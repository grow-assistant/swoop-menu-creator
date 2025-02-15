package db

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// Migrator provides ability to run migrations
type Migrator interface {
	Run() error
}

// PGMigrator is used to encapsulate the database instance
type PGMigrator struct {
	db *Database
}

// NewPostgresMigrator creates a new migrator instance
func NewPostgresMigrator(db *Database) Migrator {
	return PGMigrator{
		db,
	}
}

// Run invokes the migration
func (p PGMigrator) Run() error {

	// Gets to the *sql.DB instance
	db := p.db.DB.DB()

	driver, err := postgres.WithInstance(db, &postgres.Config{})

	m, err := migrate.NewWithDatabaseInstance(
		"file://pkg/db/migrations",
		"postgres", driver)
	if err != nil {
		log.Println(err)
		return err
	}

	err = m.Up()
	if err != nil {
		panic(err)
	}

	return nil
}
