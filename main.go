package main

import (
	"study-go/web/api"
	"study-go/web/server"
)

func main() {
	a, err := api.NewAPI()
	if err != nil {
		return
	}

	s, err := server.NewAdminServer()
	if err != nil {
		return
	}

	s.SetAPI(a)
	s.InitRoute()
	s.Run()

}
