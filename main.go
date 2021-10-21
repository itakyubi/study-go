package main

import (
	"study-go/web/api"
	"study-go/web/config"
	"study-go/web/server"
)

func main() {
	var cfg *config.Config
	err := config.LoadConfig(&cfg)
	if err != nil {
		return
	}

	a, err := api.NewAPI(cfg)
	if err != nil {
		return
	}

	s, err := server.NewAdminServer(cfg)
	if err != nil {
		return
	}

	s.SetAPI(a)
	s.InitRoute()
	s.Run()

}
