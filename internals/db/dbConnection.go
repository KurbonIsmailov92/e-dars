package db

import (
	"e-dars/configs"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var dbSession *gorm.DB

func ConnectToDb() error {
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		configs.AppSettings.PostgresParams.Host,
		configs.AppSettings.PostgresParams.User,
		os.Getenv("DB_PASSWORD"),
		configs.AppSettings.PostgresParams.Database,
		configs.AppSettings.PostgresParams.Port)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return err
	}
	dbSession = db
	return nil
}

/*
	func CloseDbConnection() error {
		return nil
	}
*/
func GetDBConnection() *gorm.DB {
	return dbSession
}
