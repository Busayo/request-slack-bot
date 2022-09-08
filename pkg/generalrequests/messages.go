package generalrequests

import (
	"fmt"


	"github.com/slack-go/slack"
)


func addSimpleSection(blocks []slack.Block, title, body string) []slack.Block {
	// Add title
	blocks = append(blocks, slack.NewSectionBlock(
		slack.NewTextBlockObject(
			slack.MarkdownType,
			fmt.Sprintf("*%s:*", title),
			false,
			false,
		),
		nil,
		nil,
	))
	// Add body
	blocks = append(blocks, slack.NewSectionBlock(
		slack.NewTextBlockObject(
			slack.MarkdownType,
			body,
			false,
			false,
		),
		nil,
		nil,
	))
	return blocks
}

func BuildRequestNotificationMessageBlocks(request GeneralRequest) []slack.Block {
	blocks := make([]slack.Block, 0)
	// Header block
	blocks = append(blocks, slack.NewHeaderBlock(
		slack.NewTextBlockObject(
			slack.PlainTextType,
			"Angkor General Request",
			false,
			false,
		),
	))
	blocks = append(blocks, slack.NewDividerBlock())
	// Requester block
	blocks = addSimpleSection(blocks, "From", fmt.Sprintf("<@%s>", request.Requester))
	// Title block
	blocks = addSimpleSection(blocks, "Title", request.Title)
	// Category block
	blocks = addSimpleSection(blocks, "Category", request.Category)
	// Description block
	blocks = addSimpleSection(blocks, "Body", request.Description)
  // Urgency block
	blocks = addSimpleSection(blocks, "Urgency", request.Urgency)
	// Divider block
	blocks = append(blocks, slack.NewDividerBlock())
	return blocks
}