package category

import (
	"context"
	"database/sql"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, category Category) Category
	Update(ctx context.Context, tx *sql.Tx, category Category) Category
	Delete(ctx context.Context, tx *sql.Tx, category Category)
	FindByID(ctx context.Context, tx *sql.Tx, categoryID int) (Category, error)
	FindByName(ctx context.Context, tx *sql.Tx, categoryName string) (Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) []Category
}
