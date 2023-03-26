package openai

import (
	"fmt"
	"openai-go/common"
)

func (o *openAI) getModelsURL() string {
	return fmt.Sprintf("%s%s", o.host, common.ModelListURL)
}

func (o *openAI) getModelURL(modelID string) string {
	return fmt.Sprintf("%s%s%s", o.host, common.ModelURL, modelID)
}

func (o *openAI) getCompletionsURL() string {
	return fmt.Sprintf("%s%s", o.host, common.CompletionsURL)
}

func (o *openAI) getChatCompletionsURL() string {
	return fmt.Sprintf("%s%s", o.host, common.ChatCompletionsURL)
}

func (o *openAI) getImageURL() string {
	return fmt.Sprintf("%s%s", o.host, common.ImageURL)
}

func (o *openAI) getEmbeddingsURL() string {
	return fmt.Sprintf("%s%s", o.host, common.EmbeddingsURL)
}

func (o *openAI) getModerationsURL() string {
	return fmt.Sprintf("%s%s", o.host, common.ModerationsURL)
}
