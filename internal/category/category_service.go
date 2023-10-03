package category

import (
	"context"

	dto "rest_base/internal/category/web/dto"
	response "rest_base/internal/category/web/response"
)

type CategoryService interface {
	Create(ctx context.Context, request dto.CategoryCreateRequest) response.CategoryResponse
	Update(ctx context.Context, request dto.CategoryUpdateRequest) response.CategoryResponse
	Delete(ctx context.Context, categoryID int)
	FindByID(ctx context.Context, categoryID int) response.CategoryResponse
	FindByName(ctx context.Context, categoryName string) response.CategoryResponse
	FindAll(ctx context.Context) []response.CategoryResponse
}
