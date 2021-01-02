package main

import (
	_ "encoding/json"
	d "github.com/barshat7/go-store/data"
)

type CategoryDTO struct {

	ID int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
}

func CategoryDTOOfCategory (c d.Category) *CategoryDTO {

	return & CategoryDTO {
		ID: c.ID,
		Name: c.Name,
		Description: c.Description,
	}
}

