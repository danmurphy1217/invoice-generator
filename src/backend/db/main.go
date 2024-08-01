package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/danmurphy1217/invoice-generator/db/collections"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
    host     = "db"
    port     = 5432
    user     = "postgres"
    password = "password"
    dbname   = "testdb"
)

// use global DB connection to avoid
// many unclosed connections to DB
var db *gorm.DB
var sqlDB *sql.DB

func Connect() (*sql.DB, *gorm.DB) {
    var err error
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

    // if pointer non-nil, we can return connection
    if sqlDB != nil && db != nil {
        return sqlDB, db
    }
    
    db, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
    if err != nil {
        log.Fatalf("Error opening database: %q", err)
    }

    // use the UUID extension to have UUIDs as primary keys
    db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

    if err := db.AutoMigrate(&collections.Invoice{}); err != nil {
        log.Fatalf("failed to migrate database: %v", err)
    }

    sqlDB, err = db.DB()
    if err != nil {
        log.Fatalf("failed to get sql DB from Gorm: %v", err)
    }

    return sqlDB, db
}