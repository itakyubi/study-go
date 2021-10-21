package api

import (
	"study-go/web/config"
	"study-go/web/service"
)

type API struct {
	User service.UserService
}

func NewAPI(cfg *config.Config) (*API, error) {
	userService, err := service.NewUserService(cfg)
	if err != nil {
		return nil, err
	}
	return &API{User: userService}, nil
}
