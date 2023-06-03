package orders

type orderErrCode string

const (
	FindOneOrderErr orderErrCode = "orders-001"
	CancelOrderErr  orderErrCode = "orders-002"
	FindOrdersErr   orderErrCode = "orders-003"
	InsertOrderErr  orderErrCode = "orders-004"
)

type Order struct {
	Id        string          `db:"id" json:"id"`
	UserId    string          `db:"user_id" json:"user_id"`
	Contact   string          `db:"contact" json:"contact"`
	Address   string          `db:"address" json:"address"`
	Status    string          `db:"status" json:"status"`
	CreatedAt string          `db:"created_at" json:"created_at"`
	UpdatedAt string          `db:"updated_at" json:"updated_at"`
	Total     float64         `json:"total"`
	Products  []*OrderProduct `json:"products"`
}

type OrderProduct struct {
	Id      string             `json:"id"`
	Qty     int                `json:"qty"`
	Product *OrderProductDatum `json:"product"`
}

type OrderProductDatum struct {
	Id          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
