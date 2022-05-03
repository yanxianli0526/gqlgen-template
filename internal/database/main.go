// Package orm provides `GORM` helpers for the creation, migration and access
// on the project's database
package orm

import (
	"fmt"
	log "log"

	//Imports the database dialect of choice
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/jinzhu/gorm"
)

var autoMigrate, logMode, seedDB bool
var dsn, dialect string

// ORM struct to holds the gorm pointer to db
type ORM struct {
	DB *gorm.DB
}

// Factory creates a db connection with the selected dialect and connection string
func Factory() (*ORM, error) {

	databaseConnect := fmt.Sprintf("sslmode=%s host=%s port=%v dbname=%s password=%s user=%s", "disable", "localhost", "5432", "inventory-toll", "", "postgres")
	fmt.Println("databaseConnect", databaseConnect)
	db, err := gorm.Open("postgres", databaseConnect)
	if err != nil {
		log.Panic("[ORM] err: ", err)
	}
	orm := &ORM{
		DB: db,
	}
	// Log every SQL command on dev, @prod: this should be disabled?
	db.LogMode(true)
	// Automigrate tables
	// err = migration.ServiceAutoMigration(orm.DB)

	log.Println("[ORM] Database connection initialized.")
	return orm, err
}
