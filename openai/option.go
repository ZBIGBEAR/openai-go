package openai

import (
	"openai-go/common"
	"openai-go/config"
	"os"
	"strings"
)

type Option func(cfg *config.OpenAI)

func WithHost(host string) Option {
	if !strings.HasSuffix(host, common.UrlSeperator) {
		host = host + common.UrlSeperator
	}
	return func(cfg *config.OpenAI) {
		cfg.Host = host
	}
}

func withDefaultHost() Option {
	return WithHost(defaultOpenAIHost)
}

func WithOpenApiKey(key string) Option {
	return func(cfg *config.OpenAI) {
		cfg.OpenAPIKey = key
	}
}

func withEnvOpenApiKey() Option {
	return WithOpenApiKey(os.Getenv(common.OpenAPIKey))
}

func WithTemperature(temperature float64) Option {
	return func(cfg *config.OpenAI) {
		cfg.Temperature = temperature
	}
}

func WithReplyMessageCount(replyMessageCount int) Option {
	return func(cfg *config.OpenAI) {
		cfg.ReplyMessageCount = replyMessageCount
	}
}
