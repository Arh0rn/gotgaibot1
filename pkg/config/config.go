package config

type Config struct {
	Env            string `yaml:"environment"`
	TelegramConfig `yaml:"tg"`
	LLMConfig      `yaml:"llm"`
}

type TelegramConfig struct {
	Token string `yaml:"token" env:"TELEGRAM_TOKEN"`
}

type LLMConfig struct {
	Provider    string  `yaml:"provider"`
	APIKey      string  `yaml:"api_key" env:"LLM_API_KEY"`
	Model       string  `yaml:"model"`
	Temperature float64 `yaml:"temperature"`
	MaxTokens   int     `yaml:"max_tokens"`
	BaseUrl     string  `yaml:"base_url"`
	Legend      string  `yaml:"legend" env:"LLM_LEGEND"`
}
