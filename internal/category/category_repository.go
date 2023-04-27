package category

import (
	"context"
	"database/sql"
	"golang_restfull_api/internal/category/model"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, category model.Category) model.Category
	Update(ctx context.Context, tx *sql.Tx, category model.Category) model.Category
	Delete(ctx context.Context, tx *sql.Tx, category model.Category)
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (model.Category, error)
	FindByName(ctx context.Context, tx *sql.Tx, categoryName string) (model.Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) []model.Category
}
