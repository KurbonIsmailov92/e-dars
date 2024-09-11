package main

import (
	"e-dars/configs"
	"e-dars/internals/db"
	"e-dars/logger"
	"github.com/joho/godotenv"
	"log"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Ошибка загрузки .env файла: %s", err)
	}

	if err := configs.ReadSettings(); err != nil {
		log.Fatalf("Ошибка чтения настроек: %s", err)
	}

	if err := logger.Init(); err != nil {
		log.Fatalf("Ошибка инициализации логирования: %s", err)
	}

	if err := db.ConnectToDb(); err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %s", err)
	}

	defer func() {
		err := db.CloseDbConnection()
		if err != nil {
			log.Fatalf("Ошибка при закрытии сессии с базой данных: %s", err)
		}
	}()

	if err := db.MigrateTables(); err != nil {
		log.Fatalf("Ошибка миграции базы данных: %s", err)
	}

	if err := db.InsertSeeds(); err != nil {
		log.Fatalf("Ошибка при загрузке первичных данных: %s", err)
	}

}
