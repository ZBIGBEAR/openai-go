package config

type Config struct {
	OpenAI `json:"openai"`
}

type OpenAI struct {
	OpenAPIKey string `json:"open_api_key" env:"OPEN_API_KEY"`
	Host       string `json:"host" env:"OPENAI_HOST"`
}
