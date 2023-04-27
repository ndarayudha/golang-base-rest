package service

import (
	"context"
	"database/sql"
	"golang_restfull_api/internal/category"
	"golang_restfull_api/internal/category/exception"
	"golang_restfull_api/internal/category/model"
	"golang_restfull_api/internal/category/web"
	"golang_restfull_api/pkg/logger"
	"golang_restfull_api/pkg/utils"
	"time"

	"github.com/go-playground/validator"
)

type CategoryServiceImpl struct {
	CategoryRepository category.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
	Logger             logger.Logger
}

func NewCategoryService(categoryRepository category.CategoryRepository, DB *sql.DB, validate *validator.Validate, logger logger.Logger) *CategoryServiceImpl {
	return &CategoryServiceImpl{CategoryRepository: categoryRepository, DB: DB, Validate: validate, Logger: logger}
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	utils.PanicIfError(err)

	tx, err := service.DB.Begin()
	utils.PanicIfError(err)
	defer utils.CommitOrRollback(tx)

	category := model.Category{
		Name:      request.Name,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	category = service.CategoryRepository.Save(ctx, tx, category)

	return model.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	utils.PanicIfError(err)

	tx, err := service.DB.Begin()
	utils.PanicIfError(err)
	defer utils.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	category.Name = request.Name

	category = service.CategoryRepository.Update(ctx, tx, category)

	return model.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	utils.PanicIfError(err)
	defer utils.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.CategoryRepository.Delete(ctx, tx, category)
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	utils.PanicIfError(err)
	defer utils.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return model.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindByName(ctx context.Context, categoryName string) web.CategoryResponse {
	tx, err := service.DB.Begin()
	utils.PanicIfError(err)
	defer utils.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindByName(ctx, tx, categoryName)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return model.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	utils.PanicIfError(err)
	defer utils.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)

	return model.ToCategoryResponses(categories)
}
