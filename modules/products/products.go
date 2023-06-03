package products

type productsErrCode string

const (
	FindProductsErr   productsErrCode = "products-001"
	FindOneProductErr productsErrCode = "products-002"
)

type Product struct {
	Id          string  `db:"id" json:"id"`
	Title       string  `db:"title" json:"title"`
	Description string  `db:"description" json:"description"`
	Price       float64 `db:"price" json:"price"`
	CreateAt    string  `db:"created_at" json:"created_at"`
	UpdatedAt   string  `db:"updated_at" json:"updated_at"`
}
