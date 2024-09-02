package db

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var session *gorm.DB

func ConnectToDb() error {
	connectionString := "host=localhost port=5432 user=kurbon dbname=e-dars_db password=ismoil"
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return err
	}
	session = db
	return nil
}

func CloseDbConnection() error {
	return nil
}

func GetDbConnection() *gorm.DB {
	return session
}
