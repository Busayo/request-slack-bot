package interactionshandlers

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/request-slack-bot/configs"
	"github.com/request-slack-bot/pkg/generalrequests"
	"github.com/request-slack-bot/pkg/utils"
	log "github.com/sirupsen/logrus"
	"github.com/slack-go/slack"
)

func sendChannelNotification(api *slack.Client, i slack.InteractionCallback, modalValues generalrequests.GeneralRequest) error {
	// Get request channel ids
	requestsChannelId, err := configs.GetRequestsChannel()
	if err != nil {
		return err
	}

	// Build message blocks
	msgBlocks := generalrequests.BuildRequestNotificationMessageBlocks(modalValues)

	_, _, err = api.PostMessage(
		requestsChannelId,
		slack.MsgOptionBlocks(msgBlocks...),
	)
	return err
}

func sendUserNotification(api *slack.Client, i slack.InteractionCallback, modalValues generalrequests.GeneralRequest) error {
	msgText := 
	fmt.Sprintf( "Hi %s, \n\nYour request has been submitted successfully :+1: \n\nIf your request is extremely urgent, please drop a message on the angkor channel :slightly_smiling_face:",
		i.User.Name)

	_, _, err := api.PostMessage(
		i.User.ID,
		slack.MsgOptionBlocks(
			slack.NewSectionBlock(
				slack.NewTextBlockObject(
					slack.MarkdownType,
					msgText,
					false,
					false,
				),
				nil,
				nil,
			),
		),
	)
	return err
}

func HandleRequestSubmission(
	event events.APIGatewayProxyRequest,
	i slack.InteractionCallback,
) (events.APIGatewayProxyResponse, error) {
	// Build request data
	modalValues, errs := generalrequests.NewGeneralRequestFromModal(i)

	// Get slack api
	token, err := configs.GetBotToken()
	if err != nil {
		log.Error(err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
		}, err
	}
	api := slack.New(token)

	// Show errors in modal
	if len(errs) != 0 {
		log.Infof("Detected errors in modal.")
		body, err := utils.BuildGeneralResponseErrorsBody(errs)
		if err != nil {
			log.Error(err)
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
			}, err
		}
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers: map[string]string{
				"content-type": "application/json",
			},
			Body: body,
		}, nil
	}

	//Send request to requests channel
	if err := sendChannelNotification(api, i, modalValues); err != nil {
		log.Error(err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	//Notify user that the request is created
	if err := sendUserNotification(api, i, modalValues); err != nil {
		log.Error(err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
	}, err
}

