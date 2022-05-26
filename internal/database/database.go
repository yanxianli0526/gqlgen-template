// Package orm provides `GORM` helpers for the creation, migration and access
// on the project's database
package orm

import (
	"fmt"
	"graphql-go-template/config"
	"graphql-go-template/internal/database/migration"
	log "log"

	//Imports the database dialect of choice
	"gorm.io/driver/postgres"
	gormLogger "gorm.io/gorm/logger"

	"gorm.io/gorm"
)

// ORM struct to holds the gorm pointer to db
type ORM struct {
	DB *gorm.DB
}

var Gorm = &gorm.Config{
	Logger: gormLogger.Default.LogMode(gormLogger.Error),
	// Logger: gormLogger.Default.LogMode(gormLogger.Info),
}

// Factory creates a db connection with the selected dialect and connection string
func Factory(env config.Database) (*ORM, error) {

	databaseConnect := fmt.Sprintf("sslmode=%s host=%s port=%v dbname=%s password=%s user=%s", env.DBSSLMode, env.DBHost, env.DBPort, env.DBname, env.DBPassword, env.DBUser)
	fmt.Println("databaseConnect", databaseConnect)
	db, err := gorm.Open(postgres.Open(databaseConnect), Gorm)
	if err != nil {
		log.Panic("[ORM] err: ", err)
	}
	orm := &ORM{
		DB: db,
	}
	// Log every SQL command on dev, @prod: this should be disabled?
	// Automigrate tables
	err = migration.UpdateMigration(orm.DB)

	log.Println("[ORM] Database connection initialized.")
	return orm, err
}
