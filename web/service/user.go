package service

import (
	"fmt"
	"study-go/web/config"
	"study-go/web/model"
	"study-go/web/plugin"
)

type UserService interface {
	Add(user *model.User) (int, error)
	Delete(id int) error
	Update(user *model.User) error
	Get(id int) (*model.User, error)
	GetByName(name string) (*model.User, error)
}

type UserServiceImpl struct {
	user plugin.User
}

func NewUserService(cfg *config.Config) (UserService, error) {
	user, err := plugin.GetPlugin(cfg.Plugin.User)
	if err != nil {
		return nil, err
	}

	return &UserServiceImpl{
		user: user.(plugin.User),
	}, nil
}

func (u UserServiceImpl) Add(user *model.User) (int, error) {
	s := fmt.Sprintf("add user, name=%s, age=%d", user.Name, user.Age)
	println(s)
	return u.user.AddUser(user)
}

func (u UserServiceImpl) Delete(id int) error {
	s := fmt.Sprintf("delete user, id=%d", id)
	println(s)
	return u.user.DeleteUser(id)
}

func (u UserServiceImpl) Update(user *model.User) error {
	s := fmt.Sprintf("update user, id=%d, name=%s, age=%d", user.Id, user.Name, user.Age)
	println(s)
	return u.user.UpdateUser(user)
}

func (u UserServiceImpl) Get(id int) (*model.User, error) {
	s := fmt.Sprintf("get user by id, id=%d", id)
	println(s)
	return u.user.GetUser(id)
}

func (u UserServiceImpl) GetByName(name string) (*model.User, error) {
	s := fmt.Sprintf("get user by name, name=%s", name)
	println(s)
	return u.user.GetUserByName(name)
}
