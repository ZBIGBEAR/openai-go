package openai

import (
	"context"
	"fmt"
	"openai-go/common"
	"openai-go/config"
	"openai-go/openai/model"
	"openai-go/pkg/http"
	"os"
	"strings"
	"sync"
)

type OpenAI interface {
	Models(ctx context.Context) ([]*model.Model, error)
	Model(ctx context.Context, modelID string) (*model.Model, error)
	Completions(ctx context.Context, prompt string) (*model.Completion, error)
	ChatCompletions(ctx context.Context, content string, count int) ([]string, error)
	Images(ctx context.Context, prompt string, count int) ([]string, error)
	Embeddings(ctx context.Context, input string) (*model.Embeddings, error)
	Moderations(ctx context.Context, input string) (*model.Moderations, error)
}

type openAI struct {
	host          string
	authorization string
	httpClient    http.Http
}

const (
	defaultOpenAIHost = "https://api.openai.com/v1/"
)

var _ OpenAI = &openAI{}

func New(cfg *config.Config) OpenAI {
	if cfg == nil {
		cfg = &config.Config{}
	}

	host := cfg.OpenAI.Host
	if host == "" {
		host = defaultOpenAIHost
	} else {
		if !strings.HasSuffix(cfg.OpenAI.Host, common.UrlSeperator) {
			host = cfg.OpenAI.Host + common.UrlSeperator
		} else {
			host = cfg.OpenAI.Host
		}
	}

	openApiKey := cfg.OpenAI.OpenAPIKey
	if openApiKey == "" {
		openApiKey := os.Getenv(common.OpenAPIKey)
		if openApiKey == "" {
			panic("open api key is empty")
		}

		cfg.OpenAI.OpenAPIKey = openApiKey
	}

	httpConfig := &http.Config{
		Authorization: fmt.Sprintf("%s %s", common.Bearer, cfg.OpenAI.OpenAPIKey),
	}
	httpClient := http.New(httpConfig, http.TimeOutOption(10))

	return &openAI{
		host:          host,
		authorization: cfg.OpenAI.OpenAPIKey,
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

func ChatCompletions(ctx context.Context, content string, count int) ([]string, error) {
	return defaultOpenAI.ChatCompletions(ctx, content, count)
}

func Images(ctx context.Context, prompt string, count int) ([]string, error) {
	return defaultOpenAI.Images(ctx, prompt, count)
}

func Embeddings(ctx context.Context, input string) (*model.Embeddings, error) {
	return defaultOpenAI.Embeddings(ctx, input)
}

func Moderations(ctx context.Context, input string) (*model.Moderations, error) {
	return defaultOpenAI.Moderations(ctx, input)
}
