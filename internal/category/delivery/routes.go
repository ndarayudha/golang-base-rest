package delivery

import (
	"golang_restfull_api/internal/category"
	"golang_restfull_api/internal/category/exception"
	"golang_restfull_api/internal/middleware"
	"golang_restfull_api/pkg/logger"

	"github.com/julienschmidt/httprouter"
)

func NewCategoryRoutes(categoryController category.CategoryController, logger logger.Logger) *httprouter.Router {

	router := httprouter.New()

	var categoryMiddleware = []middleware.Middlewares{
		middleware.RequestLoggerMiddleware,
		middleware.AuthMiddleware,
	}

	router.GET("/api/categories", middleware.BuildChain(categoryController.FindAll, logger, categoryMiddleware...))
	router.POST("/api/categories", middleware.BuildChain(categoryController.Create, logger, categoryMiddleware...))
	router.GET("/api/categories/", middleware.BuildChain(categoryController.FindByName, logger, categoryMiddleware...))
	router.GET("/api/categories/:id", middleware.BuildChain(categoryController.FindById, logger, categoryMiddleware...))
	router.PUT("/api/categories/:id", middleware.BuildChain(categoryController.Update, logger, categoryMiddleware...))
	router.DELETE("/api/categories/:id", middleware.BuildChain(categoryController.Delete, logger, categoryMiddleware...))

	router.PanicHandler = exception.ErrorHandler

	return router
}
