package openai

import (
	"context"
	"encoding/json"
	"openai-go/common"
	"openai-go/openai/model"
)

func (o *openAI) Images(ctx context.Context, prompt string) ([]string, error) {
	url := o.getImageURL()
	req := CreateImageReq{
		Prompt: prompt,
		N:      o.replyMessageCount,
		Size:   "1024x1024",
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
		return nil, common.NewError("[openAI.Images] Get. url:%s, statusCode:%d, msg:%s", url, resp.StatusCode, resp.Status)
	}
	response := &model.Image{}
	err = json.NewDecoder(resp.Body).Decode(response)
	if err != nil {
		return nil, err
	}

	result := make([]string, 0, len(response.Data))
	for i := range response.Data {
		data := response.Data[i]
		result = append(result, data.Url)
	}

	return result, nil
}
