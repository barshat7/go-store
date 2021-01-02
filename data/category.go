package data
import (
	"database/sql"
)
// Category Model representing category entity
type Category struct {
	ID int
	Name string
	Description string
}

// NewCategory Creates new instance of Category
func NewCategory(name, description string) *Category {
	return &Category{Name: name, Description: description}
}

// Save saves the category Object
func (c *Category) Save() (int64, error) {
	sql := "INSERT category SET name = ?, description = ?"
	stmt, err := Db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	res, err := stmt.Exec(c.Name, c.Description)
	lastInserted, err := res.LastInsertId()
	return lastInserted, err
}

// FindCategoryByID find by id
func FindCategoryByID(categoryID int64) *Category {
	sql := "SELECT id, name, description FROM category WHERE id = ?"
	stmt, err := Db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	res, err := stmt.Query(categoryID)
	defer closeRows(res)
	if err != nil {
		panic(err)
	}
	var category Category
	if res.Next() {
		res.Scan(&category.ID, &category.Name, &category.Description)
	}
	return &category
}

// FindAllCategories find all order by name
func FindAllCategories() (categories []Category) {
	sql := "SELECT id, name, description FROM category order by name"
	rows, err := Db.Query(sql)
	if (err != nil) {
		panic(err)
	}
	for rows.Next() {
		var category Category
		err := rows.Scan(&category.ID, &category.Name, &category.Description)
		if err != nil {
			panic(err)
		}
		categories = append(categories, category)
	}
	return categories
}

func closeRows(rows *sql.Rows) {
	if rows != nil {
		rows.Close()
	}
}