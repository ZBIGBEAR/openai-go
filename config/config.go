package config

type Config struct {
	OpenAI `json:"openai"`
}

type OpenAI struct {
	OpenAPIKey        string  `json:"open_api_key" env:"OPEN_API_KEY"`
	Host              string  `json:"host" env:"OPENAI_HOST"`
	Temperature       float64 `json:"temperature" env:"TEMPERATURE"`
	ReplyMessageCount int     `json:"reply_message_count" env:"REPLY_MESSAGE_COUNT"`
}
