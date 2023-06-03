package ordersRepository

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Rayato159/neversuitup-e-commerce-test/modules/orders"
	"github.com/jmoiron/sqlx"
)

type IOrdersRepository interface {
	InsertOrder(req *orders.Order) error
	FindOrders(userId string) []*orders.Order
	FindOneOrder(userId string, orderId string) (*orders.Order, error)
	CancelOrder(userId string) (*orders.Order, error)
}

type ordersRepository struct {
	db *sqlx.DB
}

func NewOrdersRepository(db *sqlx.DB) IOrdersRepository {
	return &ordersRepository{db}
}

func (r *ordersRepository) InsertOrder(req *orders.Order) error {
	return nil
}

func (r *ordersRepository) FindOrders(userId string) []*orders.Order {
	query := `
	SELECT
		array_to_json(array_agg("t"))
	FROM (
		SELECT
			"o"."id",
			"o"."user_id",
			"o"."contact",
			"o"."address",
			"o"."status",
			"o"."created_at",
			"o"."updated_at",
			(
				SELECT
					ROUND(SUM(("po"."product"->>'price')::NUMERIC*"po"."qty"::NUMERIC), 2)
				FROM "products_orders" "po"
				WHERE "po"."order_id" = "o"."id"
			) AS "total",
			(
				SELECT
					array_to_json(array_agg("at"))
				FROM (
					SELECT
						"po"."id",
						"po"."qty",
						"po"."product"
					FROM "products_orders" "po"
					WHERE "po"."order_id" = "o"."id"
				) AS "at"
			) AS "products"
		FROM "orders" "o"
		WHERE "o"."user_id" = $1
	) AS "t"`

	raw := make([]byte, 0)
	if err := r.db.Get(&raw, query, userId); err != nil {
		log.Printf("get orders failed: %v", err)
		return make([]*orders.Order, 0)
	}

	ordersData := make([]*orders.Order, 0)
	if err := json.Unmarshal(raw, &ordersData); err != nil {
		log.Printf("unmarshal orders failed: %v", err)
		return make([]*orders.Order, 0)
	}
	return ordersData
}

func (r *ordersRepository) FindOneOrder(userId string, orderId string) (*orders.Order, error) {
	query := `
	SELECT
		to_jsonb("t")
	FROM (
		SELECT
			"o"."id",
			"o"."user_id",
			"o"."contact",
			"o"."address",
			"o"."status",
			"o"."created_at",
			"o"."updated_at",
			(
				SELECT
					ROUND(SUM(("po"."product"->>'price')::NUMERIC*"po"."qty"::NUMERIC), 2)
				FROM "products_orders" "po"
				WHERE "po"."order_id" = "o"."id"
			) AS "total",
			(
				SELECT
					array_to_json(array_agg("at"))
				FROM (
					SELECT
						"po"."id",
						"po"."qty",
						"po"."product"
					FROM "products_orders" "po"
					WHERE "po"."order_id" = "o"."id"
				) AS "at"
			) AS "products"
		FROM "orders" "o"
		WHERE "o"."user_id" = $1
		AND "o"."id" = $2
	) AS "t"`

	raw := make([]byte, 0)
	if err := r.db.Get(&raw, query, userId, orderId); err != nil {
		return nil, fmt.Errorf("get order_id: %v failed: %v", orderId, err)
	}

	orderData := &orders.Order{
		Products: make([]*orders.OrderProduct, 0),
	}
	if err := json.Unmarshal(raw, &orderData); err != nil {
		return nil, fmt.Errorf("unmarshal order_id: %v failed: %v", orderId, err)
	}
	return orderData, nil
}

func (r *ordersRepository) CancelOrder(userId string) (*orders.Order, error) {
	return nil, nil
}
