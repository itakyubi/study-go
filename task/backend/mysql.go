package backend

import (
	"study-go/task/model"
	"time"
)

type MysqlBackend struct {
}

func NewMysqlBackend() MysqlBackend {
	return MysqlBackend{}
}

func (m *MysqlBackend) SaveResult(finished *model.TaskFinished) {
	println("MysqlBackend SaveResult")
}

func (m *MysqlBackend) GetResult(taskType, resourceId, resourceType string) *model.TaskFinished {
	println("MysqlBackend GetResult")
	return &model.TaskFinished{
		Id:                  0,
		ResourceId:          "",
		ResourceType:        "",
		TaskType:            "",
		Status:              "",
		TaskInterruptSignal: "",
		RequestId:           "",
		NextExecutableTime:  time.Time{},
		RetryCount:          0,
		Message:             "",
		Content:             "",
		LastUpdateTime:      time.Time{},
		CreateTime:          time.Time{},
		FinishTime:          time.Time{},
	}
}
