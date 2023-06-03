package whydoweneedtest

import (
	"fmt"
	"testing"

	"github.com/Rayato159/neversuitup-e-commerce-test/modules/products"
	"github.com/Rayato159/neversuitup-e-commerce-test/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestFindProducts(t *testing.T) {
	productsModule := SetupProductsTest().NewProductsModule()

	products := productsModule.Usecase().FindProducts()
	assert.Equal(t, 5, len(products))
}

type testFindOneProduct struct {
	label     string
	productId string
	expect    *products.Product
	isError   bool
}

func TestFindOneProduct(t *testing.T) {
	tests := []testFindOneProduct{
		{
			label:     "not found",
			productId: "P000099",
			expect:    nil,
			isError:   true,
		},
		{
			label:     "success",
			productId: "P000001",
			expect: &products.Product{
				Id:          "P000001",
				Title:       "Beer",
				Description: "Better than milk",
				Price:       52,
				CreateAt:    "2023-06-03T17:47:10.291603Z",
				UpdatedAt:   "2023-06-03T17:47:10.291603Z",
			},
			isError: false,
		},
	}

	productsModule := SetupProductsTest().NewProductsModule()

	for _, test := range tests {
		fmt.Println(test.label)
		product, err := productsModule.Usecase().FindOneProduct(test.productId)
		if test.isError {
			assert.NotEmpty(t, err)
		} else {
			utils.Debug(product)
			assert.Empty(t, err)
			assert.Equal(t, test.expect, product)
		}
	}
}
