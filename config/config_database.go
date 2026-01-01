package config

import (
	"early-setup/helper"
	"early-setup/model/dto"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func ConnectDatabase(dbConfig dto.DatabaseConfig) (db *gorm.DB) {
	dsn := dbConfig.User + ":" + dbConfig.Password + "@tcp(" + dbConfig.Host + ":" + dbConfig.Port +")/" + dbConfig.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	helper.PanicIfError(err)

	sqlDb, err := db.DB()
	helper.PanicIfError(err)

	sqlDb.SetMaxIdleConns(5)
	sqlDb.SetMaxOpenConns(20)

	sqlDb.SetConnMaxIdleTime(5 * time.Minute)
	sqlDb.SetConnMaxLifetime(60 * time.Minute)

	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()

	if err != nil {
		panic("Failed to close database connection")
	}

	dbSQL.Close()
}