package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"

	db "github.com/alifanza259/learn-go-library-system/db/sqlc"
	"github.com/google/uuid"
	"github.com/hibiken/asynq"
)

const TaskSendVerifyEmail = "task:send_verify_email"

type PayloadSendVerifyEmail struct {
	Username string `json:"username"`
	UUID     string `json:"uuid"`
}

func (distributor *RedisTaskDistributor) DistributeTaskSendVerifyEmail(
	ctx context.Context,
	payload *PayloadSendVerifyEmail,
	opts ...asynq.Option,
) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal task payload: %w", err)
	}

	task := asynq.NewTask(TaskSendVerifyEmail, jsonPayload, opts...)
	info, err := distributor.client.EnqueueContext(ctx, task)
	if err != nil {
		return fmt.Errorf("failed to enqueue task: %w", err)
	}

	fmt.Println(info.ID)

	return nil
}

func (processor *RedisTaskProcessor) ProcessTaskSendVerifyEmail(ctx context.Context, task *asynq.Task) error {
	var payload PayloadSendVerifyEmail

	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("failed to unmarshal payload: %w", asynq.SkipRetry)
	}

	uuid, err := uuid.Parse(payload.UUID)
	if err != nil {
		return fmt.Errorf("failed to parse UUID: %w", err)
	}

	member, err := processor.library.GetMember(ctx, uuid)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	secretCode := generateSecretCode()
	arg := db.CreateEmailVerificationParams{
		Token:    secretCode,
		MemberID: member.ID,
	}
	_, err = processor.library.CreateEmailVerification(ctx, arg)
	if err != nil {
		return fmt.Errorf("failed to create verify email: %w", err)
	}

	subject := "Complete Your Registration"
	verifyUrl := fmt.Sprintf("http://localhost:8080/v1/member/verify?token=%s", secretCode)
	content := fmt.Sprintf(`Hello %s,<br/>
	Thank you for registering with us!<br/>
	Please <a href="%s">click here</a> to verify your email address.<br/>
	`, member.FirstName, verifyUrl)
	to := []string{member.Email}

	err = processor.mailer.SendEmail(subject, content, to, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to send verify email: %w", err)
	}

	fmt.Println("processed task")
	return nil

}

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func generateSecretCode() string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < 32; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}
