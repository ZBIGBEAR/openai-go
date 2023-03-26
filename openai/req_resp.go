package openai

import "openai-go/openai/model"

type ModelsResponse struct {
	Object model.ObjectType `json:"object"`
	Data   []*model.Model   `json:"data"`
}

type CompletionsReq struct {
	Model       string      `json:"model"`
	Prompt      string      `json:"prompt"`
	MaxTokens   int         `json:"max_tokens"`
	Temperature int         `json:"temperature"`
	TopP        int         `json:"top_p"`
	N           int         `json:"n"`
	Stream      bool        `json:"stream"`
	Logprobs    interface{} `json:"logprobs"`
	Stop        string      `json:"stop"`
}

type ChatCompletionsReq struct {
	Model       string           `json:"model"`
	Messages    []*model.Message `json:"messages"`
	Temperature float64          `json:"temperature"`
	N           int              `json:"n"`
}

type CreateImageReq struct {
	Prompt string `json:"prompt"`
	N      int    `json:"n"`
	Size   string `json:"size"`
}

type EmbeddingsReq struct {
	Model string `json:"model"`
	Input string `json:"input"`
}

type ModerationsReq struct {
	Input string `json:"input"`
}
