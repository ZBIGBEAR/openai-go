package openai

import (
	"context"
	"encoding/json"
	"openai-go/common"
	"openai-go/openai/model"
)

func (o *openAI) Models(ctx context.Context) ([]*model.Model, error) {
	url := o.getModelsURL()
	resp, err := o.httpClient.Get(ctx, nil, url, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, common.NewError("[openAI.Models] Get. url:%s, statusCode:%d, msg:%s", url, resp.StatusCode, resp.Status)
	}
	response := &ModelsResponse{}
	err = json.NewDecoder(resp.Body).Decode(response)
	if err != nil {
		return nil, err
	}

	if response.Object != model.ObjectTypeOfList {
		return nil, common.NewError("[openAI.Models] response. url:%s,response.Object:%s", url, response.Object)
	}

	return response.Data, nil
}

func (o *openAI) Model(ctx context.Context, modelID string) (*model.Model, error) {
	url := o.getModelURL(modelID)
	resp, err := o.httpClient.Get(ctx, nil, url, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, common.NewError("[openAI.Model] Get. url:%s, statusCode:%d, msg:%s", url, resp.StatusCode, resp.Status)
	}
	response := &model.Model{}
	err = json.NewDecoder(resp.Body).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
