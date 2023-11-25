package worker

import (
	"context"
	"encoding/json"
	"fmt"

	db "github.com/alifanza259/learn-go-library-system/db/sqlc"
	"github.com/hibiken/asynq"
)

const TaskSendBorrowProcessedEmail = "task:send_borrow_processed_email"

type PayloadSendBorrowProcessedEmail struct {
	TransactionID string    `json:"transaction_id"`
	MemberEmail   string    `json:"member_email"`
	MemberName    string    `json:"member_name"`
	BookTitle     string    `json:"book_title"`
	Status        db.Status `json:"status"`
	Note          string    `json:"note"`
}

func (distributor *RedisTaskDistributor) DistributeTaskSendBorrowProcessedEmail(
	ctx context.Context,
	payload *PayloadSendBorrowProcessedEmail,
	opts ...asynq.Option,
) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	task := asynq.NewTask(TaskSendBorrowProcessedEmail, jsonPayload, opts...)
	info, err := distributor.client.EnqueueContext(ctx, task)
	if err != nil {
		return err
	}

	fmt.Println(info.ID)

	return nil
}

func (processor *RedisTaskProcessor) ProcessTaskSendBorrowProcessedEmail(ctx context.Context, task *asynq.Task) error {
	fmt.Println("masuk")
	var payload PayloadSendBorrowProcessedEmail
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return err
	}

	subject := "Update on Your Borrow Request Status"
	var content string
	if payload.Status == db.StatusApproved {
		content = fmt.Sprintf(`Hello %s,<br/>
		Your borrow request of book %s with transaction ID: %s, has been approved<br/>
		You can pick up the book at the library by showing this email to the librarian<br/>
		`, payload.MemberName, payload.BookTitle, payload.TransactionID)
	} else {
		content = fmt.Sprintf(`Hello %s,<br/>
		Your borrow request of book %s with transaction ID: %s, has been declined<br/>
		Note: %s<br/>
		Please resubmit your application<br/>
		`, payload.MemberName, payload.BookTitle, payload.TransactionID, payload.Note)
	}

	to := []string{payload.MemberEmail}

	err := processor.mailer.SendEmail(subject, content, to, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to send verify email: %w", err)
	}

	fmt.Println("processed task")

	return nil
}
