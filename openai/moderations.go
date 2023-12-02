package openai

import (
	"context"
	"encoding/json"
	"github.com/ZBIGBEAR/openai-go/common"
	"github.com/ZBIGBEAR/openai-go/openai/model"
)

func (o *openAI) Moderations(ctx context.Context, input string) (*model.Moderations, error) {
	url := o.getModerationsURL()
	req := ModerationsReq{
		Input: input,
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
		return nil, common.NewError("[openAI.Moderations] Get. url:%s, statusCode:%d, msg:%s", url, resp.StatusCode, resp.Status)
	}
	response := &model.Moderations{}
	err = json.NewDecoder(resp.Body).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
