package whydoweneedtest

import (
	"testing"

	"github.com/Rayato159/neversuitup-e-commerce-test/modules/users"
	"github.com/stretchr/testify/assert"
)

type testInsertUser struct {
	label   string
	req     *users.UserCredential
	expect  *users.UserPassport
	isError bool
}

func TestInsertUser(t *testing.T) {
	tests := []testInsertUser{
		{
			label: "failed -> user must be unique",
			req: &users.UserCredential{
				Username: "prayus",
				Password: "123456",
			},
			expect: &users.UserPassport{
				User: &users.User{
					Id:       "U000003",
					Username: "prayus",
				},
			},
			isError: true,
		},
		{
			label: "success",
			req: &users.UserCredential{
				Username: "test001",
				Password: "123456",
			},
			expect: &users.UserPassport{
				User: &users.User{
					Id:       "U000004",
					Username: "test001",
				},
			},
			isError: false,
		},
	}

	usersModule := SetupUsersTest().NewUsersModule()

	for _, test := range tests {
		result, err := usersModule.Usecase().InsertUser(test.req)
		if test.isError {
			assert.NotEmpty(t, err)
		} else {
			assert.Empty(t, err)
			assert.Equal(t, test.expect, result)
		}
	}
}

type testFindOneUser struct {
	label   string
	req     string
	expect  *users.User
	isError bool
}

func TestFindOneUser(t *testing.T) {
	tests := []testFindOneUser{
		{
			label:   "failed -> not found",
			req:     "U000099",
			expect:  &users.User{},
			isError: true,
		},
		{
			label: "success",
			req:   "U000001",
			expect: &users.User{
				Id:       "U000001",
				Username: "ruangyot",
			},
			isError: false,
		},
	}

	usersModule := SetupUsersTest().NewUsersModule()

	for _, test := range tests {
		result, err := usersModule.Usecase().FindOneUser(test.req)
		if test.isError {
			assert.NotEmpty(t, err)
		} else {
			assert.Empty(t, err)
			assert.Equal(t, test.expect, result)
		}
	}
}
