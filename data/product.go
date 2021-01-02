package data

import (
	"time"
	_ "database/sql"
)


type Product struct {
	ID int64
	Name string
	Description string
	Available bool
	UnitPrice float64
	Added_date time.Time
	CategoryID int
}

func NewProduct(name, description string, unitPrice float64, categoryID int) *Product {

	return & Product{
		Name: name,
		Description: description,
		Available: true,
		Added_date: time.Now(),
		UnitPrice: unitPrice,
		CategoryID: categoryID,
	}
}

func (p *Product) SaveProduct() (int64, error) {
	sql := "INSERT product SET name = ?, description = ?, available = ?, added_date = ?, unit_price = ?, category_id = ?"
	stmt, err := Db.Prepare(sql)
	if (err != nil) {
		panic(err)
	}
	defer stmt.Close()
	res, err := stmt.Exec(p.Name, p.Description, p.Available, p.Added_date, p.UnitPrice, p.CategoryID)
	lastInserted, err := res.LastInsertId()
	return lastInserted, err
}