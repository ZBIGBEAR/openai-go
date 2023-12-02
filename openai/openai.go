package openai

import (
	"context"
	"fmt"
	"github.com/ZBIGBEAR/openai-go/common"
	"github.com/ZBIGBEAR/openai-go/config"
	"github.com/ZBIGBEAR/openai-go/openai/model"
	"github.com/ZBIGBEAR/openai-go/pkg/http"
	"sync"
)

type OpenAI interface {
	Models(ctx context.Context) ([]*model.Model, error)
	Model(ctx context.Context, modelID string) (*model.Model, error)
	Completions(ctx context.Context, prompt string) (*model.Completion, error)
	ChatCompletions(ctx context.Context, content string) ([]string, error)
	Images(ctx context.Context, prompt string) ([]string, error)
	Embeddings(ctx context.Context, input string) (*model.Embeddings, error)
	Moderations(ctx context.Context, input string) (*model.Moderations, error)
}

type openAI struct {
	host              string
	temperature       float64
	replyMessageCount int
	authorization     string
	httpClient        http.Http
}

const (
	defaultOpenAIHost = "https://api.openai.com/v1/"
)

var _ OpenAI = &openAI{}

func New(cfg *config.OpenAI, options ...Option) OpenAI {
	if cfg == nil {
		cfg = &config.OpenAI{}
	}

	opts := []Option{withDefaultHost(), withEnvOpenApiKey()}
	opts = append(opts, options...)
	for i := range opts {
		opts[i](cfg)
	}

	httpConfig := &http.Config{
		Authorization: fmt.Sprintf("%s %s", common.Bearer, cfg.OpenAPIKey),
	}
	httpClient := http.New(httpConfig, http.TimeOutOption(10))

	return &openAI{
		host:          cfg.Host,
		authorization: cfg.OpenAPIKey,
		httpClient:    httpClient,
	}
}

var (
	// 定义内部默认openai实体
	defaultOpenAI     OpenAI
	defaultOpenAIOnce sync.Once
)

func init() {
	// 单例模式实例化内部默认openai实体
	defaultOpenAIOnce.Do(func() {
		defaultOpenAI = New(nil)
	})
}

func Models(ctx context.Context) ([]*model.Model, error) {
	return defaultOpenAI.Models(ctx)
}

func Model(ctx context.Context, modelID string) (*model.Model, error) {
	return defaultOpenAI.Model(ctx, modelID)
}

func Completions(ctx context.Context, prompt string) (*model.Completion, error) {
	return defaultOpenAI.Completions(ctx, prompt)
}

func ChatCompletions(ctx context.Context, content string) ([]string, error) {
	return defaultOpenAI.ChatCompletions(ctx, content)
}

func Images(ctx context.Context, prompt string) ([]string, error) {
	return defaultOpenAI.Images(ctx, prompt)
}

func Embeddings(ctx context.Context, input string) (*model.Embeddings, error) {
	return defaultOpenAI.Embeddings(ctx, input)
}

func Moderations(ctx context.Context, input string) (*model.Moderations, error) {
	return defaultOpenAI.Moderations(ctx, input)
}
