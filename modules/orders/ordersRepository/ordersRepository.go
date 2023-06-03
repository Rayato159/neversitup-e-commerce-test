package ordersRepository

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/Rayato159/neversuitup-e-commerce-test/modules/orders"
	"github.com/jmoiron/sqlx"
)

type IOrdersRepository interface {
	InsertOrder(req *orders.Order) (string, error)
	FindOrders(userId string) []*orders.Order
	FindOneOrder(userId, orderId string) (*orders.Order, error)
	FindOrderStatus(userId, orderId string) (string, error)
	CancelOrder(userId, orderId string) error
}

type ordersRepository struct {
	db *sqlx.DB
}

func NewOrdersRepository(db *sqlx.DB) IOrdersRepository {
	return &ordersRepository{db}
}

func (r *ordersRepository) InsertOrder(req *orders.Order) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return "", nil
	}

	queryInsertOrder := `
	INSERT INTO "orders" (
		"user_id",
		"contact",
		"address",
		"status"
	)
	VALUES ($1, $2, $3, 'waiting')
	RETURNING "id";`

	var orderId string
	if err := tx.QueryRowContext(
		ctx,
		queryInsertOrder,
		req.UserId,
		req.Contact,
		req.Address,
	).Scan(&orderId); err != nil {
		tx.Rollback()
		return "", fmt.Errorf("insert order failed: %v", err)
	}

	queryInsertProductsOrders := `
	INSERT INTO "products_orders" (
		"order_id",
		"qty",
		"product"
	)
	VALUES`

	values := make([]any, 0)
	lastIndex := 0

	for i, p := range req.Products {
		values = append(values, orderId, p.Qty, p)

		queryInsertProductsOrders += fmt.Sprintf(`
		($%d, $%d, $%d)`, lastIndex+1, lastIndex+2, lastIndex+3)

		if i != len(req.Products) {
			queryInsertProductsOrders += ","
		} else {
			queryInsertProductsOrders += ";"
		}

		lastIndex = len(values)
	}

	if _, err := tx.ExecContext(
		ctx,
		queryInsertProductsOrders,
		values...,
	); err != nil {
		tx.Rollback()
		return "", fmt.Errorf("insert products_orders failed: %v", err)
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return "", err
	}

	return orderId, nil
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

func (r *ordersRepository) FindOneOrder(userId, orderId string) (*orders.Order, error) {
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

func (r *ordersRepository) FindOrderStatus(userId, orderId string) (string, error) {
	_, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	query := `
	SELECT
		"status"
	FROM "orders"
	WHERE "id" = $1 
	AND "user_id" = $2;`

	var status string
	if err := r.db.Get(&status, query, orderId, userId); err != nil {
		return "", fmt.Errorf("get order_id: %v failed: %v", err, orderId)
	}
	return status, nil
}

func (r *ordersRepository) CancelOrder(userId, orderId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	query := `
	UPDATE "orders" SET
		"status" = 'canceled'
	WHERE "id" = $1 AND "user_id" = $2;`

	if _, err := r.db.ExecContext(ctx, query, orderId, userId); err != nil {
		return fmt.Errorf("cancel order_id: %v failed: %v", err, orderId)
	}
	return nil
}
