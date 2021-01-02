package main


import (
	"net/http"
	_"fmt"
	"encoding/json"
	_"github.com/gorilla/mux"
	d "github.com/barshat7/go-store/data"
)

func getCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	categories := d.FindAllCategories()
	var categoryDtoList [] *CategoryDTO
	for _, category := range categories {
		dto := CategoryDTOOfCategory(category)
		categoryDtoList = append(categoryDtoList, dto)
	} 
	json.NewEncoder(w).Encode(categoryDtoList)
}