package vms

import (

	"github.com/slack-go/slack"
)

func createHeader2() *slack.SectionBlock {
	return slack.NewSectionBlock(
		slack.NewTextBlockObject(
			slack.MarkdownType,
			"*Hi there :wave:* \n \n *Please fill out the request form below :slightly_smiling_face:*",
			false,
			false,
		),
		nil,
		nil,
	)
}

// Title block data
const TitleBlockId = "title_name"
const TitleActionId = "TITLE_NAME"

func createTitleBlock() *slack.InputBlock {
	return slack.NewInputBlock(
		TitleBlockId,
		slack.NewTextBlockObject(
			slack.PlainTextType,
			"Title",
			false,
			false,
		),
		slack.NewTextBlockObject(
			slack.PlainTextType,
			"Title",
			false,
			false,
		),
		slack.NewPlainTextInputBlockElement(
			nil,
			TitleActionId,
		), 
	)
}


// Category  block data
const CategoryBlockId = "category"
const CategoryActionId = "CATEGORY"

func createCategoryOptions() []*slack.OptionBlockObject {
	optionBlockObjects := make([]*slack.OptionBlockObject, 0, len(SupportedCategories))
	for _, option := range SupportedCategories {
		optionText := slack.NewTextBlockObject(slack.PlainTextType, option.Name, false, false)
		optionBlockObjects = append(optionBlockObjects, slack.NewOptionBlockObject(option.Value, optionText, nil))
	}
	return optionBlockObjects
}

func createCategoryBlock() *slack.InputBlock {
	categoryInput := slack.NewInputBlock(
		CategoryBlockId,
		slack.NewTextBlockObject(
			slack.PlainTextType,
			"Select a Category",
			false,
			false,
		),
		slack.NewTextBlockObject(
			slack.PlainTextType,
			"Category",
			false,
			false,
		),
		slack.NewOptionsSelectBlockElement(
			slack.OptTypeStatic,
			nil,
			CategoryActionId,
			createCategoryOptions()...,
		),
	)
	categoryInput.DispatchAction = true 
	return categoryInput 
}

// Description block data
const RequestDescriptionBlockId = "request_desc"
const RequestDescriptionActionId = "REQUEST_DESC"

func createRequestDescriptionBlock() *slack.InputBlock {
	textInputBlock := slack.NewPlainTextInputBlockElement(
		nil,
		RequestDescriptionActionId,
	)
	textInputBlock.Multiline = true
	return slack.NewInputBlock(
		RequestDescriptionBlockId,
		slack.NewTextBlockObject(
			slack.PlainTextType,
			"Body",
			false,
			false,
		),
		slack.NewTextBlockObject(
			slack.PlainTextType,
			"What is your request?",
			false,
			false,
		),
		textInputBlock,
	)
}



func BuildVMRequestModal2() slack.ModalViewRequest {
	// Modal texts
	titleText := slack.NewTextBlockObject(slack.PlainTextType, "Angkor Request", false, false)
	closeText := slack.NewTextBlockObject(slack.PlainTextType, "Cancel", false, false)
	submitText := slack.NewTextBlockObject(slack.PlainTextType, "Submit", false, false)
	// Header section
	headerSection2 := createHeader2()
	// Name input
	titleBlock := createTitleBlock()
	// Category input
	categoryBlock := createCategoryBlock()
	// Description input
	requestDescriptionBlock := createRequestDescriptionBlock()



	// Blocks
	blocks := slack.Blocks{
		BlockSet: []slack.Block{
			headerSection2,
			slack.NewDividerBlock(),
			titleBlock,
			categoryBlock,
			requestDescriptionBlock,

		},
	}
	// Modal
	var modalRequest slack.ModalViewRequest
	modalRequest.CallbackID = RequestModalCallbackId
	modalRequest.Type = slack.ViewType("modal")
	modalRequest.Title = titleText
	modalRequest.Close = closeText
	modalRequest.Submit = submitText
	modalRequest.Blocks = blocks
	return modalRequest
}