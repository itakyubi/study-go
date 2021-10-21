package service

import (
	"fmt"
	"study-go/web/config"
	"study-go/web/model"
	"time"
)

type UserService interface {
	Add(user *model.User) (int, error)
	Delete(id int) error
	Update(user *model.User) error
	Get(id int) (*model.User, error)
	GetByName(name string) (*model.User, error)
}

type UserServiceImpl struct {
}

func NewUserService(cfg *config.Config) (UserService, error) {
	return &UserServiceImpl{}, nil
}

func (u UserServiceImpl) Add(user *model.User) (int, error) {
	s := fmt.Sprintf("add user, name=%s, age=%d", user.Name, user.Age)
	println(s)
	return int(time.Now().Unix()), nil
}

func (u UserServiceImpl) Delete(id int) error {
	s := fmt.Sprintf("delete user, id=%d", id)
	println(s)
	return nil
}

func (u UserServiceImpl) Update(user *model.User) error {
	s := fmt.Sprintf("update user, id=%d, name=%s, age=%d", user.Id, user.Name, user.Age)
	println(s)
	return nil
}

func (u UserServiceImpl) Get(id int) (*model.User, error) {
	s := fmt.Sprintf("get user by id, id=%d", id)
	println(s)
	return &model.User{
		Id:   id,
		Name: string(time.Now().Second()),
		Age:  time.Now().Second(),
	}, nil
}

func (u UserServiceImpl) GetByName(name string) (*model.User, error) {
	s := fmt.Sprintf("get user by name, name=%s", name)
	println(s)
	return &model.User{
		Id:   time.Now().Second(),
		Name: name,
		Age:  time.Now().Second(),
	}, nil
}
