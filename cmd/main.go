package main

import (
	"log"
	"os"
	"rest_base/config"
	"rest_base/internal/category/controller"
	"rest_base/internal/category/repository"
	"rest_base/internal/category/routes"
	"rest_base/internal/category/service"
	"rest_base/internal/server"
	"rest_base/pkg/db/mysql"
	"rest_base/pkg/logger"
	"rest_base/pkg/utils"

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
	categoryRouter := routes.NewCategoryRoutes(categoryControllerImpl, appLogger)

	httpServer := server.NewServer(config, categoryRouter)

	appLogger.Infof("Server running at http://localhost:%v/api/", config.Server.Port)
	err = httpServer.ListenAndServe()
	utils.PanicIfError(err)
}
