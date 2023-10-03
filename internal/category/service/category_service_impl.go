package service

import (
	"context"
	"database/sql"
	"rest_base/internal/category"
	"rest_base/internal/category/exception"
	"rest_base/internal/category/web"
	dto "rest_base/internal/category/web/dto"
	response "rest_base/internal/category/web/response"
	"rest_base/pkg/logger"
	"rest_base/pkg/utils"
	"time"

	"github.com/go-playground/validator"
)

type CategoryServiceImpl struct {
	CategoryRepository category.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
	Logger             logger.Logger
}

func NewCategoryService(categoryRepository category.CategoryRepository, db *sql.DB, validate *validator.Validate, logger logger.Logger) *CategoryServiceImpl {
	return &CategoryServiceImpl{CategoryRepository: categoryRepository, DB: db, Validate: validate, Logger: logger}
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request dto.CategoryCreateRequest) response.CategoryResponse {
	err := service.Validate.Struct(request)
	utils.PanicIfError(err)

	tx, err := service.DB.Begin()
	utils.PanicIfError(err)
	defer utils.CommitOrRollback(tx)

	category := category.Category{
		Name:      request.Name,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	category = service.CategoryRepository.Save(ctx, tx, category)

	return web.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request dto.CategoryUpdateRequest) response.CategoryResponse {
	err := service.Validate.Struct(request)
	utils.PanicIfError(err)

	tx, err := service.DB.Begin()
	utils.PanicIfError(err)
	defer utils.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindByID(ctx, tx, request.ID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	category.Name = request.Name

	category = service.CategoryRepository.Update(ctx, tx, category)

	return web.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryID int) {
	tx, err := service.DB.Begin()
	utils.PanicIfError(err)
	defer utils.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindByID(ctx, tx, categoryID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.CategoryRepository.Delete(ctx, tx, category)
}

func (service *CategoryServiceImpl) FindByID(ctx context.Context, categoryID int) response.CategoryResponse {
	tx, err := service.DB.Begin()
	utils.PanicIfError(err)
	defer utils.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindByID(ctx, tx, categoryID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return web.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindByName(ctx context.Context, categoryName string) response.CategoryResponse {
	tx, err := service.DB.Begin()
	utils.PanicIfError(err)
	defer utils.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindByName(ctx, tx, categoryName)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return web.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []response.CategoryResponse {
	tx, err := service.DB.Begin()
	utils.PanicIfError(err)
	defer utils.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)

	return web.ToCategoryResponses(categories)
}
