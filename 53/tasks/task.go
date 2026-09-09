package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/hibiken/asynq"
)

// Task types
const (
	TypeEmailDelivery = "email:deliver"
	TypeImageResize   = "image:resize"
)

// Task payloads
type EmailDeliveryPayload struct {
	UserID     int    `json:"user_id"`
	TemplateID string `json:"template_id"`
}

type ImageResizePayload struct {
	SourceURL string `json:"source_url"`
}

// ----------------------------------------------------
// Task Constructors
// ----------------------------------------------------

func NewEmailDeliveryTask(userID int, tmplID string) (*asynq.Task, error) {
	payload, err := json.Marshal(EmailDeliveryPayload{UserID: userID, TemplateID: tmplID})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeEmailDelivery, payload), nil
}

func NewImageResizeTask(src string) (*asynq.Task, error) {
	payload, err := json.Marshal(ImageResizePayload{SourceURL: src})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(
		TypeImageResize,
		payload,
		asynq.MaxRetry(5),
		asynq.Timeout(20*time.Minute),
	), nil
}

// ----------------------------------------------------
// Task Handlers
// ----------------------------------------------------

// HandlerFunc style (Stateless)
func HandleEmailDeliveryTask(ctx context.Context, t *asynq.Task) error {
	var p EmailDeliveryPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	log.Printf("Sending Email to User: user_id=%d, template_id=%s", p.UserID, p.TemplateID)
	return nil
}

// Struct style Handler (Stateful / Dependency Injection)
type ImageProcessor struct {
	// e.g. S3 Client, Database connection
}

func NewImageProcessor() *ImageProcessor {
	return &ImageProcessor{}
}

func (processor *ImageProcessor) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var p ImageResizePayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	log.Printf("Resizing image: src=%s", p.SourceURL)
	return nil
}
