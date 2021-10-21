package api

import "study-go/web/service"

type API struct {
	User service.UserService
}

func NewAPI() (*API, error) {
	userService, err := service.NewUserService()
	if err != nil {
		return nil, err
	}
	return &API{User: userService}, nil
}
