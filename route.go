package main


import (
	"net/http"
	_"fmt"
	"encoding/json"
	"github.com/gorilla/mux"
	d "github.com/barshat7/go-store/data"
	"strconv"
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

func getCategoryByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	categoryID, _ := strconv.ParseInt(params["id"], 10, 64)
	category := d.FindCategoryByID(categoryID)
	dto := CategoryDTOOfCategory(*category)
	json.NewEncoder(w).Encode(dto)
}

func saveCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var categoryDTO CategoryDTO
	_ = json.NewDecoder(r.Body).Decode(&categoryDTO)
	newCategory := d.NewCategory(categoryDTO.Name, categoryDTO.Description)
	newCategory.Save()
	json.NewEncoder(w).Encode(categoryDTO)
}