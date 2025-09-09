package dto

import "github.com/leomoritz/fullcycle-arq-hexagonal-app/application"

type ProductDto struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Status string  `json:"status"`
}

func (dto *ProductDto) Bind(product *application.Product) (*application.Product, error) {
	if dto.ID != "" {
		product.ID = dto.ID
	}

	product.Name = dto.Name
	product.Price = dto.Price
	product.Status = dto.Status

	_, err := product.IsValid()
	if err != nil {
		return &application.Product{}, err
	}

	return product, nil
}
