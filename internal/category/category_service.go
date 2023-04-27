package category

import (
	"context"
	"golang_restfull_api/internal/category/web"
)

type CategoryService interface {
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) web.CategoryResponse
	FindByName(ctx context.Context, categoryName string) web.CategoryResponse
	FindAll(ctx context.Context) []web.CategoryResponse
}
