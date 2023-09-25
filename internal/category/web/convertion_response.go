package web

import (
	"rest_base/internal/category"
	web "rest_base/internal/category/web/response"
)

func ToCategoryResponse(category category.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:        category.Id,
		Name:      category.Name,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}
}

func ToCategoryResponses(categories []category.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}

	return categoryResponses
}
