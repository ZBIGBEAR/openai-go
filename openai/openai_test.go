package openai

import (
	"context"
	"github.com/stretchr/testify/assert"
	"openai-go/common"
	"openai-go/config"
	"openai-go/openai/model"
	"os"
	"testing"
)

func getOpenAI() OpenAI {
	cfg := &config.Config{
		OpenAI: config.OpenAI{
			OpenAPIKey: os.Getenv(common.OpenAPIKey),
		},
	}

	return New(cfg)
}

func TestModel(t *testing.T) {
	ai := getOpenAI()
	ctx := context.Background()
	models, err := ai.Models(ctx)
	assert.Nil(t, err)
	// 目前官方会返回62个model
	assert.NotEqual(t, 0, len(models))
	for i := range models {
		m, err := ai.Model(ctx, string(models[i].ID))
		assert.Nil(t, err)
		assert.Equal(t, model.ObjectTypeOfModel, m.Object)
		assert.NotEqual(t, 0, len(m.Permissions))
	}
}

func TestCompletions(t *testing.T) {
	ai := getOpenAI()
	ctx := context.Background()
	prompt := "Say this is a test"
	completions, err := ai.Completions(ctx, prompt)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(completions.Choices))
}

func TestChatCompletions(t *testing.T) {
	ai := getOpenAI()
	ctx := context.Background()
	prompt := "你是谁"
	result, err := ai.ChatCompletions(ctx, prompt)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))
}

func TestImages(t *testing.T) {
	ai := getOpenAI()
	ctx := context.Background()
	prompt := "韩国美女"
	result, err := ai.Images(ctx, prompt)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))
}

func TestEmbeddings(t *testing.T) {
	ai := getOpenAI()
	ctx := context.Background()
	input := "你是谁"
	result, err := ai.Embeddings(ctx, input)
	assert.Nil(t, err)
	assert.NotEqual(t, 0, len(result.Data))
	assert.Equal(t, model.ObjectTypeOfEmbedding, result.Data[0].Object)
}

func TestModerations(t *testing.T) {
	ai := getOpenAI()
	ctx := context.Background()
	input := "性爱"
	result, err := ai.Moderations(ctx, input)
	assert.Nil(t, err)
	assert.Equal(t, true, result.Results[0].Categories.Sexual)
}
