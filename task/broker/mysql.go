package broker

import (
	"study-go/task/model"
	"time"
)

type MysqlBroker struct {
}

func NewMysqlBroker() MysqlBroker {
	return MysqlBroker{}
}

func (m *MysqlBroker) AddTask(taskType, resourceId, resourceType, requestId, content string) {
	println("MysqlBroker AddTask")
}

func (m *MysqlBroker) AddTaskWithExecTime(taskType, resourceId, resourceType, requestId, content string, execTime time.Time) {
	println("MysqlBroker AddTaskWithExecTime")
}

func (m *MysqlBroker) GetTasks() []model.Task {
	println("MysqlBroker GetTasks")
	return []model.Task{}
}

func (m *MysqlBroker) LockTask(task model.Task) bool {
	println("MysqlBroker LockTask")
	return true
}
