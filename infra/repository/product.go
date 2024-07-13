package repository

import (
	"go-ci/domain/entity"
	derror "go-ci/domain/error"
	irepository "go-ci/domain/repository"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) irepository.IProductRepository {
	return &ProductRepository{db}
}

func (pr *ProductRepository) Insert(name string, price uint, sellerID uint) (entity.Product, error) {
	product := entity.Product{
		Name:     name,
		Price:    price,
		SellerID: sellerID,
	}

	if err := pr.db.Create(&product).Error; err != nil {
		if err.Error() == "FOREIGN KEY constraint failed" {
			println("seller not found")
			return entity.Product{}, derror.ErrUserNotFound
		}

		return entity.Product{}, err
	}

	return product, nil
}

func (pr *ProductRepository) GetAll() ([]entity.Product, error) {
	var products []entity.Product

	if err := pr.db.Preload("Seller").Find(&products).Error; err != nil {
		return products, err
	}

	return products, nil
}

func (pr *ProductRepository) Delete(sellerID uint, id uint) error {
	return pr.db.Where("seller_id = ?", sellerID).Delete(&entity.Product{}, id).Error
}
