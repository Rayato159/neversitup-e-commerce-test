package whydoweneedtest

import (
	"testing"

	"github.com/Rayato159/neversuitup-e-commerce-test/modules/orders"
	"github.com/stretchr/testify/assert"
)

type testFindOrders struct {
	userId string
}

func TestFindOrders(t *testing.T) {
	tests := []testFindOrders{
		{
			userId: "U000001",
		},
	}

	ordersModule := SetupUsersTest().NewOrdersModule()

	for _, test := range tests {
		result := ordersModule.Usecase().FindOrders(test.userId)
		assert.Equal(t, 2, len(result))
	}
}

type testFindOneOrder struct {
	userId  string
	orderId string
	isError bool
}

func TestFindOneOrder(t *testing.T) {
	tests := []testFindOneOrder{
		{
			userId:  "U000099",
			orderId: "O000001",
			isError: true,
		},
		{
			userId:  "U000001",
			orderId: "O000099",
			isError: true,
		},
		{
			userId:  "U000001",
			orderId: "O000001",
			isError: false,
		},
	}

	ordersModule := SetupUsersTest().NewOrdersModule()

	for _, test := range tests {
		result, err := ordersModule.Usecase().FindOneOrder(test.userId, test.orderId)
		if test.isError {
			assert.NotEmpty(t, err)
		} else {
			assert.Empty(t, err)
			assert.NotEmpty(t, result)
		}
	}
}

type testCancelOrder struct {
	userId  string
	orderId string
	isError bool
}

func TestCancelOrder(t *testing.T) {
	tests := []testFindOneOrder{
		{
			userId:  "U000099",
			orderId: "O000002",
			isError: true,
		},
		{
			userId:  "U000002",
			orderId: "O000099",
			isError: true,
		},
		{
			userId:  "U000001",
			orderId: "O000001",
			isError: true,
		},
		{
			userId:  "U000001",
			orderId: "O000002",
			isError: false,
		},
	}

	ordersModule := SetupUsersTest().NewOrdersModule()

	for _, test := range tests {
		result, err := ordersModule.Usecase().CancelOrder(test.userId, test.orderId)
		if test.isError {
			assert.NotEmpty(t, err)
		} else {
			assert.Empty(t, err)
			assert.NotEmpty(t, result)
			assert.Equal(t, "canceled", result.Status)
		}
	}
}

type testInsertOrder struct {
	req     *orders.Order
	isError bool
}

func TestInsertOrder(t *testing.T) {
	tests := []testInsertOrder{
		{
			req: &orders.Order{
				UserId:  "U000001",
				Address: "Test Address",
				Contact: "Pita Pita",
				Total:   104,
				Products: []*orders.OrderProduct{
					{
						Qty: 2,
						Product: &orders.OrderProductDatum{
							Id:          "P000001",
							Title:       "Beer",
							Description: "Better than milk",
							Price:       52,
						},
					},
				},
			},
			isError: false,
		},
	}

	ordersModule := SetupUsersTest().NewOrdersModule()

	for _, test := range tests {
		result, err := ordersModule.Usecase().InsertOrder(test.req)
		if test.isError {
			assert.NotEmpty(t, err)
		} else {
			assert.Empty(t, err)
			assert.NotEmpty(t, result)
		}
	}
}
