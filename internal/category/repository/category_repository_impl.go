package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang_restfull_api/internal/category/model"
	"golang_restfull_api/pkg/utils"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category model.Category) model.Category {
	result, err := tx.ExecContext(ctx, createCategory, category.Name, category.CreatedAt, category.UpdatedAt)
	utils.PanicIfError(err)

	id, err := result.LastInsertId()
	utils.PanicIfError(err)

	category.Id = int(id)
	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category model.Category) model.Category {
	_, err := tx.ExecContext(ctx, updateCategory, category.Name, category.Id)
	utils.PanicIfError(err)

	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category model.Category) {
	_, err := tx.ExecContext(ctx, deleteCategoryById, category.Id)
	utils.PanicIfError(err)
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (model.Category, error) {
	rows, err := tx.QueryContext(ctx, findCategoryById, categoryId)
	utils.PanicIfError(err)

	defer rows.Close()

	category := model.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		utils.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("category tidak ditemukan")
	}
}

func (repository *CategoryRepositoryImpl) FindByName(ctx context.Context, tx *sql.Tx, categoryName string) (model.Category, error) {
	rows, err := tx.QueryContext(ctx, findCategoryByName, categoryName)
	utils.PanicIfError(err)

	defer rows.Close()

	category := model.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		utils.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("category tidak ditemukan")
	}
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []model.Category {
	rows, err := tx.QueryContext(ctx, findCategories)
	utils.PanicIfError(err)

	defer rows.Close()
	var categories []model.Category
	for rows.Next() {
		category := model.Category{}
		err := rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt)
		utils.PanicIfError(err)
		categories = append(categories, category)
	}

	return categories
}
