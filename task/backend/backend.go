package backend

import "study-go/task/model"

type TaskBackend interface {
	SaveResult(finished *model.TaskFinished)
	GetResult(taskType, resourceId, resourceType string) *model.TaskFinished
}
