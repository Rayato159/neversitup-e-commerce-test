package productsUsecase

import "github.com/Rayato159/neversuitup-e-commerce-test/modules/products/productsRepository"

type IProductsUsecase interface{}

type productsUsecase struct {
	productsRepository productsRepository.IProductsRepository
}

func NewProductsUsecase(productsRepository productsRepository.IProductsRepository) IProductsUsecase {
	return &productsUsecase{
		productsRepository: productsRepository,
	}
}
