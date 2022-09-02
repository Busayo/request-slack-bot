package vms

import (

	"github.com/slack-go/slack"
)

const RequestModalCallbackId = "general-request-modal"

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



// Category  block data
const UrgencyBlockId = "urgency"
const UrgencyActionId = "URGENCY"

func createUrgencyOptions() []*slack.OptionBlockObject {
	optionBlockObjects := make([]*slack.OptionBlockObject, 0, len(SupportedUrgency))
	for _, option := range SupportedUrgency {
		optionText := slack.NewTextBlockObject(slack.PlainTextType, option.Name, false, false)
		optionBlockObjects = append(optionBlockObjects, slack.NewOptionBlockObject(option.Value, optionText, nil))
	}
	return optionBlockObjects
}

func createUrgencyBlock() *slack.InputBlock {
	urgencyInput := slack.NewInputBlock(
		UrgencyBlockId,
		slack.NewTextBlockObject(
			slack.PlainTextType,
			"Urgency",
			false,
			false,
		),
		slack.NewTextBlockObject(
			slack.PlainTextType,
			"Urgency",
			false,
			false,
		),
		slack.NewOptionsSelectBlockElement(
			slack.OptTypeStatic,
			nil,
			UrgencyActionId,
			createUrgencyOptions()...,
		),
	)
	urgencyInput.DispatchAction = true 
	return urgencyInput 
}


func BuildVMRequestModal2() slack.ModalViewRequest {
	// Modal texts
	titleText2 := slack.NewTextBlockObject(slack.PlainTextType, "Angkor Request", false, false)
	closeText2 := slack.NewTextBlockObject(slack.PlainTextType, "Cancel", false, false)
	submitText2 := slack.NewTextBlockObject(slack.PlainTextType, "Submit", false, false)
	// Header section
	headerSection2 := createHeader2()
	// Name input
	titleBlock := createTitleBlock()
	// Category input
	categoryBlock := createCategoryBlock()
	// Description input
	requestDescriptionBlock := createRequestDescriptionBlock()
	// Urgency input
	requestUrgencyBlock := createUrgencyBlock()



	// Blocks
	blocks2 := slack.Blocks{
		BlockSet: []slack.Block{
			headerSection2,
			slack.NewDividerBlock(),
			titleBlock,
			categoryBlock,
			requestDescriptionBlock,
			requestUrgencyBlock,
			slack.NewDividerBlock(),
		},
	}
	// Modal
	var modalRequest slack.ModalViewRequest
	modalRequest.CallbackID = RequestModalCallbackId
	modalRequest.Type = slack.ViewType("modal")
	modalRequest.Title = titleText2
	modalRequest.Close = closeText2
	modalRequest.Submit = submitText2
	modalRequest.Blocks = blocks2
	return modalRequest
}