package openai

import (
	"context"
	"encoding/json"
	"openai-go/common"
	"openai-go/openai/model"
)

func (o *openAI) ChatCompletions(ctx context.Context, content string) ([]string, error) {
	url := o.getChatCompletionsURL()
	req := ChatCompletionsReq{
		Model:       "gpt-3.5-turbo",
		Temperature: o.temperature,
		Messages: []*model.Message{
			{
				Role:    "user",
				Content: content,
			},
		},
		N: o.replyMessageCount,
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
		return nil, common.NewError("[openAI.ChatCompletions] Get. url:%s, statusCode:%d, msg:%s", url, resp.StatusCode, resp.Status)
	}
	response := &model.ChatCompletions{}
	err = json.NewDecoder(resp.Body).Decode(response)
	if err != nil {
		return nil, err
	}

	result := make([]string, 0, len(response.Choices))
	for i := range response.Choices {
		choice := response.Choices[i]
		result = append(result, choice.Message.Content)
	}

	return result, nil
}
