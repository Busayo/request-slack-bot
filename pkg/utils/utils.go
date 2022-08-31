package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/request-slack-bot/configs"
	"github.com/aws/aws-lambda-go/events"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	log "github.com/sirupsen/logrus"
	"github.com/slack-go/slack"
)

func VerifySigningSecret(r events.APIGatewayProxyRequest) error {
	signingSecret, err := configs.GetSigningSecret()
	if err != nil {
		log.Error(err)
		return err
	}

	verifier, err := slack.NewSecretsVerifier(r.MultiValueHeaders, signingSecret)
	if err != nil {
		log.Error(err)
		return err
	}

	verifier.Write([]byte(r.Body))
	if err = verifier.Ensure(); err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func EnsureDecoded(event *events.APIGatewayProxyRequest) error {
	if event.IsBase64Encoded {
		body, err := base64.StdEncoding.DecodeString(event.Body)
		if err != nil {
			return err
		}
		event.Body = string(body)
		event.IsBase64Encoded = false
	}
	return nil
}

func BuildHTTPRequestFrom(event events.APIGatewayProxyRequest) (*http.Request, error) {
	adapter := httpadapter.New(http.DefaultServeMux)
	return adapter.EventToRequest(event)
}

func GetMessageFrom(api *slack.Client, channelID, messageTs string) (slack.Message, error) {
	conversation, err := api.GetConversationHistory(&slack.GetConversationHistoryParameters{
		ChannelID: channelID,
		Latest:    messageTs,
		Limit:     1,
		Inclusive: true,
	})
	if err != nil {
		return slack.Message{}, err
	}

	if len(conversation.Messages) < 1 {
		return slack.Message{}, fmt.Errorf("no message found on channel %s with Ts %s", channelID, messageTs)
	}

	msg := conversation.Messages[0]
	return msg, nil
}

func BuildInteractiveCallback(event events.APIGatewayProxyRequest) (slack.InteractionCallback, error) {
	var i slack.InteractionCallback
	r, err := BuildHTTPRequestFrom(event)
	if err != nil {
		return i, err
	}
	err = json.Unmarshal([]byte(r.FormValue("payload")), &i)
	return i, err
}
