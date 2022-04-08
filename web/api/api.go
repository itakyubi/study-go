package api

import (
	"study-go/web/config"
	"study-go/web/log"
	"study-go/web/service"
)

type API struct {
	User service.UserService
	log  *log.Logger
}

func NewAPI(cfg *config.Config) (*API, error) {
	userService, err := service.NewUserService(cfg)
	if err != nil {
		return nil, err
	}
	return &API{
		User: userService,
		log:  log.L().With(log.Any("api", "admin")),
	}, nil
}
