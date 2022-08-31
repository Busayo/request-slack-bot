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

func GetRequestsChannel() (string, error) {
	requestChannel := os.Getenv("SLACK_REQUESTS_CHANNEL")
	if requestChannel == "" {
		return requestChannel, fmt.Errorf("slack signing secret is empty")
	}
	return requestChannel, nil
}

func GetLinodeAccessToken() (string, error) {
	token := os.Getenv("LINODE_ACCESS_TOKEN")
	if token == "" {
		return token, fmt.Errorf("linode access token is empty")
	}
	return token, nil
}

func GetPulumiServiceURL() (string, error) {
	url := os.Getenv("PULUMI_SERVICE_URL")
	if url == "" {
		return url, fmt.Errorf("pulumi service url is empty")
	}
	return url, nil
}

func GetTopicArn() (string, error) {
	arn := os.Getenv("NOTIFICATION_TOPIC_ARN")
	if arn == "" {
		return arn, fmt.Errorf("notifications Topic ARN is empty")
	}
	return arn, nil
}

func GetNotionAccessToken() (string, error) {
	token := os.Getenv("NOTION_ACCESS_TOKEN")
	if token == "" {
		return token, fmt.Errorf("notion access token is empty")
	}
	return token, nil
}

func GetNotionDBID() (string, error) {
	id := os.Getenv("NOTION_DB_ID")
	if id == "" {
		return id, fmt.Errorf("notion db id is empty")
	}
	return id, nil
}

func GetSecurityKey() (string, error) {
	key := os.Getenv("SECURITY_KEY")
	if key == "" {
		return key, fmt.Errorf("security key is empty")
	}
	if len(key) != 32 {
		return key, fmt.Errorf("invalid security key")
	}
	return key, nil
}
