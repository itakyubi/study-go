package main

import (
	"study-go/task/backend"
	"study-go/task/broker"
	"study-go/task/handler"
)

func main() {
	// 初始化broker、backend、worker
	taskBroker := broker.NewMysqlBroker()
	taskBackend := backend.NewMysqlBackend()
	taskWorker := NewWorker(&taskBroker, &taskBackend)

	// 注册handler
	deleteSkillHandler := handler.NewDeleteSkillHandler()
	taskWorker.RegisterTaskHandler("deleteSkill", &deleteSkillHandler)

	// 启动worker
	go taskWorker.Run()

	// 添加任务
	taskBroker.AddTask("deleteSkill", "1", "skill", "1231212", "{}")

	<-make(chan bool)
}
