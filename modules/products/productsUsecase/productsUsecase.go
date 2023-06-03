package productsUsecase

import (
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/products"
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/products/productsRepository"
)

type IProductsUsecase interface {
	FindProducts() []*products.Product
	FindOneProduct(productId string) (*products.Product, error)
}

type productsUsecase struct {
	productsRepository productsRepository.IProductsRepository
}

func NewProductsUsecase(productsRepository productsRepository.IProductsRepository) IProductsUsecase {
	return &productsUsecase{
		productsRepository: productsRepository,
	}
}

func (u *productsUsecase) FindProducts() []*products.Product {
	return u.productsRepository.FindProducts()
}

func (u *productsUsecase) FindOneProduct(productId string) (*products.Product, error) {
	product, err := u.productsRepository.FindOneProduct(productId)
	if err != nil {
		return nil, err
	}
	return product, nil
}
