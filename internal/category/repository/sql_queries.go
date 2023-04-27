package repository

const (
	createCategory     = "insert into category(name, created_at, updated_at) values (?, ?, ?)"
	updateCategory     = "update category set name = ? where id = ?"
	deleteCategoryById = "delete from category where id = ?"
	findCategoryById   = "select id, name from category where id = ?"
	findCategoryByName = "select id, name from category where name = ?"
	findCategories     = "select id, name, created_at, updated_at from category"
)
