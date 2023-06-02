package whydoweneedtest

import (
	"encoding/json"

	"github.com/Rayato159/neversuitup-e-commerce-test/config"
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/servers"
	"github.com/Rayato159/neversuitup-e-commerce-test/pkg/database"
)

func SetupUsersTest() servers.IModule {
	cfg := config.LoadConfig("./.env.users.test")

	db := database.DbConnect(cfg.Db())

	s := servers.NewServer(cfg, db)
	return servers.NewModule(s.GetServer(), nil)
}

func SetupProductsTest() servers.IModule {
	cfg := config.LoadConfig("./.env.products.test")

	db := database.DbConnect(cfg.Db())

	s := servers.NewServer(cfg, db)
	return servers.NewModule(s.GetServer(), nil)
}

func CompressToJSON(obj any) string {
	result, _ := json.Marshal(&obj)
	return string(result)
}
