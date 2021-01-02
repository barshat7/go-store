package main

import (
	_ "github.com/go-sql-driver/mysql"
	d "github.com/barshat7/go-store/data"
	"log"
)

func main() {

	product := d.NewProduct("IPhone X", "SmartPhone", 100.00, 1)
	id, _ := product.SaveProduct()
	log.Println(id)
}