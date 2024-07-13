package dto

type InsertProductRequest struct {
	Name  string `json:"name" validate:"required"`
	Price uint   `json:"price" validate:"required,min=1"`
}
