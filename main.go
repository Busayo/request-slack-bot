package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/request-slack-bot/configs"  //change back to right links
	"github.com/request-slack-bot/pkg/generalrequests"
	"github.com/request-slack-bot/pkg/utils" //change back to right links
  "github.com/request-slack-bot/pkg/vms"  //change back to right links
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
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

	// Parse Slash Command
	slashCMD, err := buildSlashCommand(event)
	if err != nil {
		log.Error(err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
		}, err
	}

	// Handle command
	if slashCMD.Command == "/request-vm" { // Check if is request command
		token, err := configs.GetBotToken()
		if err != nil {
			log.Error(err)
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
			}, err
		}
		api := slack.New(token)

		// Build modal
		modalRequest := vms.BuildVMRequestModal()


		if _, err := api.OpenView(slashCMD.TriggerID, modalRequest); err != nil { // Show modal
			log.Error(err)
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
			}, err
		}

	
	} else if slashCMD.Command == "/request"{
		token, err := configs.GetBotToken()
		if err != nil {
			log.Error(err)
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
			}, err
		}
		api := slack.New(token)


		
		// Build modal
		modalRequest2 := generalrequests.BuildVMRequestModal2()


		if _, err := api.OpenView(slashCMD.TriggerID, modalRequest2); err != nil { // Show modal
			log.Error(err)
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
			}, err
		}

	
	}else { // Wrong command
		err := fmt.Errorf("invalid command executed. expected \"/request-vm\" or \"/request\" but got %s", slashCMD.Command)
		log.Error(err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
	}, nil
}

func buildSlashCommand(event events.APIGatewayProxyRequest) (slack.SlashCommand, error) {
	r, err := utils.BuildHTTPRequestFrom(event)
	if err != nil {
		return slack.SlashCommand{}, err
	}
	return slack.SlashCommandParse(r)
}




func main() {
	lambda.Start(Handler)
}