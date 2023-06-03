package whydoweneedtest

import (
	"testing"

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
