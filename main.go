package main

import (
	"golang_restfull_api/config"
	"golang_restfull_api/internal/category/controller"
	"golang_restfull_api/internal/category/delivery"
	"golang_restfull_api/internal/category/repository"
	"golang_restfull_api/internal/category/service"
	"golang_restfull_api/internal/server"
	"golang_restfull_api/pkg/db/mysql"
	"golang_restfull_api/pkg/logger"
	"golang_restfull_api/pkg/utils"
	"log"
	"os"

	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	log.Println("Starting api server")

	configPath := utils.GetConfigPath(os.Getenv("config"))
	configFile, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}
	config, err := config.ParseConfig(configFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}

	appLogger := logger.NewApiLogger(config)
	appLogger.InitLogger()
	appLogger.Infof("AppVersion: %s, LogLevel: %s, Mode: %s", config.Server.AppVersion, config.Logger.Level, config.Server.Mode)

	db := mysql.NewDB(config)
	validate := validator.New()
	categoryRepositoryImpl := repository.NewCategoryRepository()
	categoryServiceImpl := service.NewCategoryService(categoryRepositoryImpl, db, validate, appLogger)
	categoryControllerImpl := controller.NewCategoryController(categoryServiceImpl)
	categoryRouter := delivery.NewCategoryRoutes(categoryControllerImpl, appLogger)

	httpServer := server.NewServer(config, categoryRouter)

	appLogger.Infof("Server running at http://localhost:%v/api/", config.Server.Port)
	err = httpServer.ListenAndServe()
	utils.PanicIfError(err)
}
