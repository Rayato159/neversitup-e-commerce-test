package whydoweneedtest

import (
	"os"
	"testing"

	"github.com/Rayato159/neversuitup-e-commerce-test/modules/users"
	"github.com/Rayato159/neversuitup-e-commerce-test/pkg/lockkunchae"
	"github.com/stretchr/testify/assert"
)

func TestSignApiKey(t *testing.T) {
	auth, err := lockkunchae.Newlockkunchae(lockkunchae.ApiKey, GetConfig().Jwt(), nil)
	assert.Empty(t, err)

	err = os.WriteFile("apikey.txt", []byte(auth.SignToken()), 0777)
	assert.Empty(t, err)
}

func TestSignAccessToken(t *testing.T) {
	auth, err := lockkunchae.Newlockkunchae(lockkunchae.ApiKey, GetConfig().Jwt(), &users.UserClaims{
		Id:       "U000001",
		Username: "ruangyot",
	})
	assert.Empty(t, err)

	err = os.WriteFile("accesstoken.txt", []byte(auth.SignToken()), 0777)
	assert.Empty(t, err)
}

func TestSignRefreshToken(t *testing.T) {
	auth, err := lockkunchae.Newlockkunchae(lockkunchae.ApiKey, GetConfig().Jwt(), &users.UserClaims{
		Id:       "U000001",
		Username: "ruangyot",
	})
	assert.Empty(t, err)

	err = os.WriteFile("refreshtoken.txt", []byte(auth.SignToken()), 0777)
	assert.Empty(t, err)
}
