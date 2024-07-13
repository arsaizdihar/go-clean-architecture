package usecase

import (
	"go-ci/domain/entity"
	irepository "go-ci/domain/repository"
)

type ProductUseCase interface {
	Insert(name string, price uint, sellerID uint) (entity.Product, error)
	GetAll() ([]entity.Product, error)
	Delete(sellerID uint, id uint) error
}

type ProductUseCaseImpl struct {
	productRepository irepository.IProductRepository
}

func NewProductUseCase(productRepository irepository.IProductRepository) ProductUseCase {
	return &ProductUseCaseImpl{productRepository}
}

func (puc *ProductUseCaseImpl) Insert(name string, price uint, sellerID uint) (entity.Product, error) {
	return puc.productRepository.Insert(name, price, sellerID)
}

func (puc *ProductUseCaseImpl) GetAll() ([]entity.Product, error) {
	return puc.productRepository.GetAll()
}

func (puc *ProductUseCaseImpl) Delete(sellerID uint, id uint) error {
	return puc.productRepository.Delete(sellerID, id)
}
