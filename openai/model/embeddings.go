package model

type Embeddings struct {
	Object ObjectType       `json:"object"`
	Data   []EmbeddingsData `json:"data"`
	Model  string           `json:"model"`
	Usage  Usage            `json:"usage"`
}

type EmbeddingsData struct {
	Object    ObjectType `json:"object"`
	Index     int        `json:"index"`
	Embedding []float64  `json:"embedding"`
}
