package handler

import (
	"study-go/task/model"
)

type DeleteSkillHandler struct {
}

func NewDeleteSkillHandler() DeleteSkillHandler {
	return DeleteSkillHandler{}
}

func (h *DeleteSkillHandler) CreateContext(task model.Task) (interface{}, error) {
	println("DeleteSkillHandler CreateContext")
	return nil, nil
}

func (h *DeleteSkillHandler) IsTaskNeedExecute(task model.Task, ctx interface{}) (bool, error) {
	println("DeleteSkillHandler IsTaskNeedExecute")
	return false, nil
}

func (h *DeleteSkillHandler) RunTask(task model.Task, ctx interface{}) error {
	println("DeleteSkillHandler DoCheckPrecondition")
	return nil
}

func (h *DeleteSkillHandler) DoSuccess(task model.Task, ctx interface{}) error {
	println("DeleteSkillHandler DoSuccess")
	return nil
}

func (h *DeleteSkillHandler) DoFailed(task model.Task, ctx interface{}) error {
	println("DeleteSkillHandler DoFailed")
	return nil
}
