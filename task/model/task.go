package model

import "time"

type Task struct {
	Id                  int       `db:"id" json:"id"`
	ResourceId          string    `db:"resource_id" json:"resourceId"`
	ResourceType        string    `db:"resource_type" json:"resourceType"`
	TaskType            string    `db:"task_type" json:"taskType"`
	Status              string    `db:"status" json:"status"`
	TaskInterruptSignal string    `db:"task_interrupt_signal" json:"taskInterruptSignal"`
	RequestId           string    `db:"request_id" json:"requestId"`
	NextExecutableTime  time.Time `db:"next_executable_time" json:"nextExecutableTime"`
	RetryCount          int       `db:"retry_count" json:"retryCount"`
	Message             string    `db:"message" json:"message"`
	Content             string    `db:"content" json:"content"`
	CreateTime          time.Time `db:"create_time" json:"createTime"`
	UpdateTime          time.Time `db:"update_time" json:"updateTime"`
}
