package utils

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	TelegramBotToken string `mapstructure:"TELEGRAM_BOT_TOKEN"`
	DatabaseURL      string `mapstructure:"DATABASE_URL"`
	DeepSeekAPIKey   string `mapstructure:"DEEPSEEK_API_KEY"`
	Prompt           string `mapstructure:"PROMPT"`
	AI_URL           string `mapstructure:"AI_URL"`
	Model            string `mapstructure:"MODEL"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &config, nil
}
