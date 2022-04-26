package main

import (
	"study-go/task/backend"
	"study-go/task/broker"
	"study-go/task/handler"
	"study-go/web/log"
	"time"
)

type TaskWorker struct {
	taskBroker   broker.TaskBroker
	taskBackend  backend.TaskBackend
	taskHandlers map[string]handler.TaskHandler
	log          *log.Logger
}

func NewWorker(taskBroker broker.TaskBroker, taskBackend backend.TaskBackend) TaskWorker {
	return TaskWorker{
		taskBroker:   taskBroker,
		taskBackend:  taskBackend,
		taskHandlers: make(map[string]handler.TaskHandler),
		log:          log.L().With(log.Any("task", "worker")),
	}
}

func (t *TaskWorker) Run() {
	for {
		tasks := t.taskBroker.GetTasks()
		for _, task := range tasks {
			h := t.taskHandlers[task.TaskType]
			ctx, _ := h.CreateContext(task)
			if ok, _ := h.IsTaskNeedExecute(task, ctx); ok {
				err := h.RunTask(task, ctx)
				if err != nil {
					h.DoFailed(task, ctx)
				} else {
					h.DoSuccess(task, ctx)
				}
			}
		}
		time.Sleep(time.Second)
	}
}

func (t *TaskWorker) RegisterTaskHandler(name string, handler handler.TaskHandler) {
	t.taskHandlers[name] = handler
}
