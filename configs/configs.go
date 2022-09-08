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

//func GetTopicArn() (string, error) {
//	arn := os.Getenv("NOTIFICATION_TOPIC_ARN")
//	if arn == "" {
//		return arn, fmt.Errorf("notifications Topic ARN is empty")
//	}
//	return arn, nil
//}

//func GetSecurityKey() (string, error) {
//	key := os.Getenv("SECURITY_KEY")
//	if key == "" {
//		return key, fmt.Errorf("security key is empty")
//	}
//	if len(key) != 32 {
//		return key, fmt.Errorf("invalid security key")
//	}
//	return key, nil
//}
// minus the extra lines that are unneeded here but needed in the main program, this checks out.

