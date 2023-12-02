package openai

import (
	"context"
	"encoding/json"
	"github.com/ZBIGBEAR/openai-go/common"
	"github.com/ZBIGBEAR/openai-go/openai/model"
)

func (o *openAI) Embeddings(ctx context.Context, input string) (*model.Embeddings, error) {
	url := o.getEmbeddingsURL()
	req := EmbeddingsReq{
		Model: "text-embedding-ada-002",
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
		return nil, common.NewError("[openAI.Embeddings] Get. url:%s, statusCode:%d, msg:%s", url, resp.StatusCode, resp.Status)
	}
	response := &model.Embeddings{}
	err = json.NewDecoder(resp.Body).Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
