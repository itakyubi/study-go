package api

import (
	"strconv"
	"study-go/web/common"
	"study-go/web/model"
)

func (api *API) AddUser(c *common.Context) (interface{}, error) {
	user := new(model.User)
	err := c.BindJSON(user)
	userId, err := api.User.Add(user)
	if err != nil {
		return nil, err
	}
	return userId, nil
}

func (api *API) DeleteUser(c *common.Context) (interface{}, error) {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return nil, err
	}
	err = api.User.Delete(userId)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (api *API) UpdateUser(c *common.Context) (interface{}, error) {
	user := new(model.User)
	err := c.BindJSON(user)
	userId, err := strconv.Atoi(c.Param("userId"))
	user.Id = userId
	err = api.User.Update(user)

	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (api *API) GetUser(c *common.Context) (interface{}, error) {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return nil, err
	}
	user, err := api.User.Get(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (api *API) GetUserByName(c *common.Context) (interface{}, error) {
	userName, _ := c.GetQuery("userName")
	user, err := api.User.GetByName(userName)
	if err != nil {
		return nil, err
	}
	return user, nil
}
