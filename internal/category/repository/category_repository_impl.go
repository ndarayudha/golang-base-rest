package repository

import (
	"context"
	"database/sql"
	"errors"
	"rest_base/internal/category"
	"rest_base/pkg/utils"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category category.Category) category.Category {
	result, err := tx.ExecContext(ctx, createCategory, category.Name, category.CreatedAt, category.UpdatedAt)
	utils.PanicIfError(err)

	id, err := result.LastInsertId()
	utils.PanicIfError(err)

	category.ID = int(id)
	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category category.Category) category.Category {
	_, err := tx.ExecContext(ctx, updateCategory, category.Name, category.ID)
	utils.PanicIfError(err)

	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category category.Category) {
	_, err := tx.ExecContext(ctx, deleteCategoryByID, category.ID)
	utils.PanicIfError(err)
}

func (repository *CategoryRepositoryImpl) FindByID(ctx context.Context, tx *sql.Tx, categoryID int) (category.Category, error) {
	rows, err := tx.QueryContext(ctx, findCategoryByID, categoryID)
	utils.PanicIfError(err)

	defer rows.Close()

	category := category.Category{}
	if rows.Next() {
		err := rows.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.UpdatedAt)
		utils.PanicIfError(err)
		return category, nil
	}

	return category, errors.New("category not found")
}

func (repository *CategoryRepositoryImpl) FindByName(ctx context.Context, tx *sql.Tx, categoryName string) (category.Category, error) {
	rows, err := tx.QueryContext(ctx, findCategoryByName, categoryName)
	utils.PanicIfError(err)

	defer rows.Close()

	category := category.Category{}
	if rows.Next() {
		err := rows.Scan(&category.ID, &category.Name)
		utils.PanicIfError(err)
		return category, nil
	}

	return category, errors.New("category not found")
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []category.Category {
	rows, err := tx.QueryContext(ctx, findCategories)
	utils.PanicIfError(err)

	defer rows.Close()
	var categories []category.Category
	for rows.Next() {
		category := category.Category{}
		err := rows.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.UpdatedAt)
		utils.PanicIfError(err)
		categories = append(categories, category)
	}

	return categories
}
