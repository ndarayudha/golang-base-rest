package routes

import (
	"rest_base/internal/category"
	"rest_base/internal/category/exception"
	"rest_base/internal/middleware"
	"rest_base/pkg/logger"

	"github.com/julienschmidt/httprouter"
)

func NewCategoryRoutes(categoryController category.CategoryController, logger logger.Logger) *httprouter.Router {

	router := httprouter.New()

	categoryBaseRoute := "/api/categories"

	categoryMiddleware := []middleware.Middlewares{
		middleware.RequestLoggerMiddleware,
	}

	routes := []struct {
		method  string
		path    string
		handler httprouter.Handle
		mw      []middleware.Middlewares
	}{
		{"GET", "", categoryController.FindAll, categoryMiddleware},
		{"POST", "", categoryController.Create, categoryMiddleware},
		{"GET", "/:id", categoryController.FindById, categoryMiddleware},
		{"PUT", "/:id", categoryController.Update, categoryMiddleware},
		{"DELETE", "/:id", categoryController.Delete, categoryMiddleware},
	}

	for _, route := range routes {
		router.Handle(route.method, categoryBaseRoute+route.path, middleware.BuildChain(route.handler, logger, route.mw...))
	}

	router.PanicHandler = exception.ErrorHandler

	return router
}
