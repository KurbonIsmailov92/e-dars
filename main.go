package main

import (
	"e-dars/configs"
	"e-dars/internals/db"
	"e-dars/logger"
	"e-dars/pkg/repository"
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

// @title E-dars API
// @version 1.0
// @description API Server for E-dars Application

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Ошибка загрузки .env файла: %s", err)
	}

	if err := configs.ReadSettings(); err != nil {
		log.Fatalf("Ошибка чтения настроек: %s", err)
	}

	if err := logger.Init(); err != nil {
		log.Fatalf("Ошибка инициализации логгера: %s", err)
	}

	var err error
	err = db.ConnectToDb()
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %s", err)
	}

	if err = db.MigrateTables(); err != nil {
		log.Fatalf("Ошибка миграции базы данных: %s", err)
	}
	fmt.Println(repository.GetAllClasses())

	/*if err = db.InsertSeeds(); err != nil {
		log.Fatalf("Ошибка при загрузки необходимых данных в таблицы: %s", err)
	}

	mainServer := new(server.Server)
	go func() {
		if err = mainServer.Run(configs.AppSettings.AppParams.PortRun, controllers.InitRoutes()); err != nil {
			log.Fatalf("Ошибка при запуске HTTP сервера: %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if sqlDB, err := db.GetDBConnection().DB(); err == nil {
		if err := sqlDB.Close(); err != nil {
			log.Fatalf("Ошибка при закрытии соединения с БД: %s", err)
		}
	} else {
		log.Fatalf("Ошибка при получении *sql.DB из GORM: %s", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = mainServer.Shutdown(ctx); err != nil {
		log.Fatalf("Ошибка при завершении работы сервера: %s", err)
	}
	*/
}
