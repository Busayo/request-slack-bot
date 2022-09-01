# request-slack-bot
## golang lambda function Slack bot 


This is a slack bot that causes separate modals to appear when either the slash commands "/request" or "/request-vm" are used

For this to run, a slack app needs to have been created with the appropriate OAUTH permissions assigned to it. Then the two slash commands can be added.

For the request URL to make use of, create an AWS Lambda Function with a Golang runtime, add an APIGateway (REST) trigger and use the given URL as the request URL for both Slash Commands. (Don't forget to add SLACK_BOT_TOKEN and SLACK_SIGNING_SECRET as environment variables.

Then on your local machine, build the goland program using 

GOOS=linux go build main.go 

(where main.go is the name of the main golang file in this program)

Then zip the resulting executable file by typing this on the terminal

zip function.zip main 

Go back to AWS Lambda and upload your zipped file to the code editor. Change the name of the Handler under Runtime Settings to main. 

Then go back to Slack to test your slash command. If all works well, you should have a modal that takes in requests.

The interaction for the modal will be added later on.
