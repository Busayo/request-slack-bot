package configs

import (
	"fmt"
	"os"
)

func GetBotToken() (string, error) {
	botToken := os.Getenv("SLACK_BOT_TOKEN")
	if botToken == "" {
		return botToken, fmt.Errorf("slack bot token is empty")
	}
	return botToken, nil
}

func GetSigningSecret() (string, error) {
	signingSecret := os.Getenv("SLACK_SIGNING_SECRET")
	if signingSecret == "" {
		return signingSecret, fmt.Errorf("slack signing secret is empty")
	}
	return signingSecret, nil
}



