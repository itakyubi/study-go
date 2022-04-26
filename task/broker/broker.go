package broker

import (
	"study-go/task/model"
	"time"
)

type TaskBroker interface {
	AddTask(taskType, resourceId, resourceType, requestId, content string)
	AddTaskWithExecTime(taskType, resourceId, resourceType, requestId, content string, execTime time.Time)
	GetTasks() []model.Task
	LockTask(task model.Task) bool
}
