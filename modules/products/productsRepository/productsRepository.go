package productsRepository

import "github.com/jmoiron/sqlx"

type IProductsRepository interface{}

type productsRepository struct {
	db *sqlx.DB
}

func NewProductsRepository(db *sqlx.DB) IProductsRepository {
	return &productsRepository{db: db}
}
