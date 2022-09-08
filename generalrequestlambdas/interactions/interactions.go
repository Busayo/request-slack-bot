package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	generalhandlers "github.com/request-slack-bot/generalrequestlambdas/interactions/interactions_handlers"
	"github.com/request-slack-bot/pkg/generalrequests"
	"github.com/request-slack-bot/pkg/utils"
//	"github.com/request-slack-bot/pkg/vms"
	log "github.com/sirupsen/logrus"
	"github.com/slack-go/slack"
)

func Handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Ensure event is decoded
	if err := utils.EnsureDecoded(&event); err != nil {
		log.Error(err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
		}, err
	}

	// Verify signing secret
	if err := utils.VerifySigningSecret(event); err != nil {
		log.Error(err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusUnauthorized,
		}, err
	}

	// Build interaction callback
	i, err := utils.BuildInteractiveCallback(event)
	if err != nil {
		log.Error(err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusUnauthorized,
		}, err
	}

	if i.Type == slack.InteractionTypeViewSubmission { // Is a modal submission
		switch i.View.CallbackID {
//		case vms.RequestModalCallbackId: // Handle request modal
//			return handlers.HandleRequestSubmission(event, i)
		case generalrequests.GeneralRequestModalCallbackId:
			return generalhandlers.HandleRequestSubmission(event, i)
		}
	}
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
