package main

import (
	_ "github.com/go-sql-driver/mysql"
	"runtime"
	"study-go/web/api"
	_ "study-go/web/common"
	"study-go/web/config"
	"study-go/web/plugin"
	_ "study-go/web/plugin/database"
	"study-go/web/server"
)

func main() {
	defer plugin.ClosePlugins()
	runtime.GOMAXPROCS(runtime.NumCPU())

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
