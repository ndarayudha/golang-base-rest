package model

import (
	"golang_restfull_api/internal/category/web"
)

type Category struct {
	Id        int
	Name      string
	CreatedAt int64
	UpdatedAt int64
}

func ToCategoryResponse(category Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:        category.Id,
		Name:      category.Name,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}
}

func ToCategoryResponses(categories []Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}

	return categoryResponses
}
