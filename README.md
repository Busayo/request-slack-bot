## Request Slack Bot

This is a Slack bot that creates separate modals when slash commands "/request" or "/request-vm" are used. To use the bot, you will need to create a Slack app and assign the appropriate OAUTH permissions to it. The two slash commands can then be added to the app.

### Prerequisites

To run the bot, you will need to create an AWS Lambda function with a Golang runtime and an APIGateway (REST) trigger. You will also need to set the following environment variables in your Lambda function:

- SLACK_BOT_TOKEN
- SLACK_SIGNING_SECRET

### Building the Golang Program

To build the Golang program, run the following command on your local machine:

```GOOS=linux go build main.go```

(where "main.go" is the name of the main Golang file in the program)

### Creating the Executable File

To create the executable file, zip the resulting executable file by running the following command in your terminal:

```zip function.zip main```

### Uploading the Executable File to AWS Lambda

Once you have created the executable file, upload it to the AWS Lambda code editor. Change the name of the Handler under Runtime Settings to "main".

### Testing the Slack Bot

To test the Slack bot, go back to Slack and try out your slash command. If everything is set up correctly, you should see a modal that takes in requests.

Note: It is recommended that you test your Lambda function locally before deploying it to AWS Lambda. You can do this by setting up a local environment and running your function on your local machine.

