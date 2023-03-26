package model

/*
model示例如下：
{
	"id": "text-davinci-003",
	"object": "model",
	"created": 1669599635,
	"owned_by": "openai-internal",
	"permission": [
		{
			"id": "modelperm-BptJFQovWB2rXq80ox1pVvza",
			"object": "model_permission",
			"created": 1679527838,
			"allow_create_engine": false,
			"allow_sampling": true,
			"allow_logprobs": true,
			"allow_search_indices": false,
			"allow_view": true,
			"allow_fine_tuning": false,
			"organization": "*",
			"group": null,
			"is_blocking": false
		}
	],
	"root": "text-davinci-003",
	"parent": null
}
*/

type ObjectType string
type ModelID string
type PermissonID string

const (
	ObjectTypeOfModel           ObjectType = "model"
	ObjectTypeOfModelPermission ObjectType = "model_permission"
	ObjectTypeOfObject          ObjectType = "object"
	ObjectTypeOfList            ObjectType = "list"
	ObjectTypeOfTextCompletion  ObjectType = "text_completion"
	ObjectTypeOfEmbedding       ObjectType = "embedding"
)

type Model struct {
	ID          ModelID      `json:"id"`
	Object      ObjectType   `json:"object"`
	Created     int64        `json:"created"`
	OwnedBy     string       `json:"owned_by"`
	Permissions []Permission `json:"permission"`
	Root        string       `json:"root"`
	Parent      string       `json:"parent"`
}

type Permission struct {
	ID                 PermissonID `json:"id"`
	Object             ObjectType  `json:"object"`
	Created            int64       `json:"created"`
	AllowCreateEngine  bool        `json:"allow_create_engine"`
	AllowSampling      bool        `json:"allow_sampling"`
	AllowLogprobs      bool        `json:"allow_logprobs"`
	AllowSearchIndices bool        `json:"allow_search_indices"`
	AllowView          bool        `json:"allow_view"`
	AllowFineTuning    bool        `json:"allow_fine_tuning"`
	Organization       string      `json:"organization"`
	Group              interface{} `json:"group"`
	IsBlocking         bool        `json:"is_blocking"`
}
