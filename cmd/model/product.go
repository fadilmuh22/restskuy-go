package model


type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description *string  `json:"description"`
	Stock       int     `json:"stock"`
}