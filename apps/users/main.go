package main

import (
	"os"

	"github.com/Rayato159/neversuitup-e-commerce-test/config"
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/servers"
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/servers/usersServer"
	"github.com/Rayato159/neversuitup-e-commerce-test/pkg/database"
)

func main() {
	var cfgPath string
	if len(os.Args) == 1 {
		cfgPath = ".env.users.dev"
	} else {
		cfgPath = os.Args[1]
	}

	cfg := config.LoadConfig(cfgPath)

	db := database.DbConnect(cfg.Db())
	defer db.Close()

	server := servers.NewServer(cfg, db)
	usersServer.NewUsersServer(server).Start()
}
