package db

import (
	"time"

	"swoop/pkg/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Database is a facade over gorm.DB
type Database struct {
	*gorm.DB
}

// Connect returns an instance of Database
func Connect(c config.DBConfig) *Database {
	db, err := gorm.Open("postgres", c.ToConnectionString())

	// db.LogMode(true)

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	db.DB().SetMaxIdleConns(5)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	db.DB().SetMaxOpenConns(50)

	// SetConnMaxLifetime sets the maximum amount of time an idle connection may be reused.
	db.DB().SetConnMaxIdleTime(time.Minute * 5)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	db.DB().SetConnMaxLifetime(time.Hour)

	if err != nil {
		panic(err)
	}

	return &Database{
		db,
	}
}
