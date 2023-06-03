package productsRepository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Rayato159/neversuitup-e-commerce-test/modules/products"
	"github.com/jmoiron/sqlx"
)

type IProductsRepository interface {
	FindProducts() []*products.Product
	FindOneProduct(productId string) (*products.Product, error)
}

type productsRepository struct {
	db *sqlx.DB
}

func NewProductsRepository(db *sqlx.DB) IProductsRepository {
	return &productsRepository{db: db}
}

func (r *productsRepository) FindProducts() []*products.Product {
	_, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	query := `
	SELECT
		"p"."id",
		"p"."title",
		"p"."description",
		"p"."price",
		"p"."created_at",
		"p"."updated_at"
	FROM "products" "p";`

	productsData := make([]*products.Product, 0)
	if err := r.db.Select(&productsData, query); err != nil {
		log.Printf("select products failed: %v", err)
		return make([]*products.Product, 0)
	}
	return productsData
}

func (r *productsRepository) FindOneProduct(productId string) (*products.Product, error) {
	_, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	query := `
	SELECT
		"p"."id",
		"p"."title",
		"p"."description",
		"p"."price",
		"p"."created_at",
		"p"."updated_at"
	FROM "products" "p"
	WHERE "p"."id" = $1;`

	product := new(products.Product)
	if err := r.db.Get(product, query, productId); err != nil {
		return nil, fmt.Errorf("get product failed: %v", err)
	}
	return product, nil
}
