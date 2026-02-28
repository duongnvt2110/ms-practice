package usecase

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"ms-practice/retry-management-service/pkg/model"
	"ms-practice/retry-management-service/pkg/repository"

	"github.com/segmentio/kafka-go"
)

type DLQUsecase interface {
	Ingest(ctx context.Context, msg kafka.Message) error
	List(ctx context.Context, page, pageSize int) ([]model.DLQRecord, int64, error)
	GetByID(ctx context.Context, id int64) (*model.DLQRecord, error)
}

type dlqUsecase struct {
	repo repository.DLQRepository
}

func NewDLQUsecase(repo repository.DLQRepository) DLQUsecase {
	return &dlqUsecase{repo: repo}
}

type headerView struct {
	Key          string `json:"key"`
	ValueBase64  string `json:"value_base64"`
	ValuePreview string `json:"value_preview"`
}

func (u *dlqUsecase) Ingest(ctx context.Context, msg kafka.Message) error {
	headersJSON, err := encodeHeaders(msg.Headers)
	if err != nil {
		return fmt.Errorf("encode headers: %w", err)
	}

	payloadJSON := bestEffortPrettyJSON(msg.Value)

	record := &model.DLQRecord{
		Topic:       msg.Topic,
		Partition:   msg.Partition,
		Offset:      msg.Offset,
		Key:         msg.Key,
		Headers:     headersJSON,
		Payload:     msg.Value,
		PayloadJSON: payloadJSON,
	}
	return u.repo.Create(ctx, record)
}

func (u *dlqUsecase) List(ctx context.Context, page, pageSize int) ([]model.DLQRecord, int64, error) {
	return u.repo.List(ctx, page, pageSize)
}

func (u *dlqUsecase) GetByID(ctx context.Context, id int64) (*model.DLQRecord, error) {
	return u.repo.GetByID(ctx, id)
}

func encodeHeaders(headers []kafka.Header) (string, error) {
	if len(headers) == 0 {
		return "[]", nil
	}
	out := make([]headerView, 0, len(headers))
	for _, h := range headers {
		preview := string(h.Value)
		if len(preview) > 200 {
			preview = preview[:200] + "..."
		}
		out = append(out, headerView{
			Key:          h.Key,
			ValueBase64:  base64.StdEncoding.EncodeToString(h.Value),
			ValuePreview: preview,
		})
	}
	b, err := json.Marshal(out)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func bestEffortPrettyJSON(payload []byte) *string {
	if len(payload) == 0 {
		return nil
	}
	var v interface{}
	if err := json.Unmarshal(payload, &v); err != nil {
		return nil
	}
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return nil
	}
	s := string(b)
	return &s
}
