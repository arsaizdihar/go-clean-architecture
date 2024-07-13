package irepository

import "go-ci/domain/entity"

type IProductRepository interface {
	Insert(name string, price uint, sellerID uint) (entity.Product, error)
	GetAll() ([]entity.Product, error)
	Delete(sellerID uint, id uint) error
}
