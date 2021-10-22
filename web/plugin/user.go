package plugin

import "study-go/web/model"

type User interface {
	AddUser(user *model.User) (int, error)
	DeleteUser(id int) error
	UpdateUser(user *model.User) error
	GetUser(id int) (*model.User, error)
	GetUserByName(name string) (*model.User, error)
}
