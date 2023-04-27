//go:build wireinject
// +build wireinject

package main

import (
	"golang_restfull_api/app"
	"golang_restfull_api/config"
	"golang_restfull_api/internal/category"
	"golang_restfull_api/internal/category/controller"
	"golang_restfull_api/internal/category/repository"
	"golang_restfull_api/internal/category/service"
	"golang_restfull_api/internal/middleware"
	"golang_restfull_api/internal/server"
	"golang_restfull_api/pkg/db/mysql"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepository,
	wire.Bind(new(category.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
	service.NewCategoryService,
	wire.Bind(new(category.CategoryService), new(*service.CategoryServiceImpl)),
	controller.NewCategoryController,
	wire.Bind(new(category.CategoryController), new(*controller.CategoryControllerImpl)),
)

func IntializedServer(c *config.Config) *http.Server {
	wire.Build(
		mysql.NewDB,
		validator.New,
		categorySet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		server.NewServer,
	)

	return nil
}
