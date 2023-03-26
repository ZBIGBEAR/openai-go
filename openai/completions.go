package openai

import (
	"context"
	"encoding/json"
	"openai-go/common"
	"openai-go/openai/model"
)

func (o *openAI) Completions(ctx context.Context, prompt string) (*model.Completion, error) {
	url := o.getCompletionsURL()
	req := CompletionsReq{
		Model:       "text-davinci-003",
		Prompt:      "Say this is a test",
		MaxTokens:   7,
		Temperature: 0,
		TopP:        1,
		N:           1,
		Stream:      false,
		Logprobs:    nil,
		Stop:        "\n",
	}

	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	resp, err := o.httpClient.Post(ctx, nil, url, reqBytes)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, common.NewError("[openAI.Completions] Get. url:%s, statusCode:%d, msg:%s", url, resp.StatusCode, resp.Status)
	}
	response := &model.Completion{}
	err = json.NewDecoder(resp.Body).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
