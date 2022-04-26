package handler

import "study-go/task/model"

type TaskHandler interface {
	CreateContext(task model.Task) (interface{}, error)
	IsTaskNeedExecute(task model.Task, ctx interface{}) (bool, error)
	RunTask(task model.Task, ctx interface{}) error
	DoSuccess(task model.Task, ctx interface{}) error
	DoFailed(task model.Task, ctx interface{}) error
}
