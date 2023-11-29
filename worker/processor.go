package worker

import (
	"context"
	"fmt"

	db "github.com/alifanza259/learn-go-library-system/db/sqlc"
	"github.com/alifanza259/learn-go-library-system/mail"
	"github.com/hibiken/asynq"
)

const (
	QueueCritical = "critical"
	QueueDefault  = "default"
)

type TaskProcessor interface {
	Start() error
	ProcessTaskSendVerifyEmail(ctx context.Context, task *asynq.Task) error
	ProcessTaskSendBorrowProcessedEmail(ctx context.Context, task *asynq.Task) error
	ProcessTaskSendReturnProcessedEmail(ctx context.Context, task *asynq.Task) error
}

type RedisTaskProcessor struct {
	server  *asynq.Server
	library db.Library
	mailer  mail.EmailSender
}

func NewRedisTaskProcessor(redisOpt asynq.RedisClientOpt, library db.Library, mailer mail.EmailSender) TaskProcessor {
	server := asynq.NewServer(
		redisOpt,
		asynq.Config{
			Queues: map[string]int{
				QueueCritical: 10,
				QueueDefault:  5,
			},
			ErrorHandler: asynq.ErrorHandlerFunc(func(ctx context.Context, task *asynq.Task, err error) {
				fmt.Println("process task failed: ", task.Type(), err)
			}),
		},
	)

	return &RedisTaskProcessor{
		server:  server,
		library: library,
		mailer:  mailer,
	}
}

func (processor *RedisTaskProcessor) Start() error {
	mux := asynq.NewServeMux()

	// mux.HandleFunc(TaskSendVerifyEmail, processor.ProcessTaskSendVerifyEmail)
	mux.HandleFunc(TaskSendBorrowProcessedEmail, processor.ProcessTaskSendBorrowProcessedEmail)

	return processor.server.Start(mux)
}
