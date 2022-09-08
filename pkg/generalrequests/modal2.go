package generalrequests

import (
	"fmt"

	"github.com/slack-go/slack"
)

const GeneralRequestModalCallbackId = "general-request-modal"

func createHeader() *slack.SectionBlock {
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
const TitleBlockId = "general_request_title_name"
const TitleActionId = "GENERAL_REQUEST_TITLE_NAME"

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
const CategoryBlockId = "general_request_category"
const CategoryActionId = "GENERAL_REQUEST_CATEGORY"

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
const RequestDescriptionBlockId = "general_request_desc"
const RequestDescriptionActionId = "GENERAL_REQUEST_DESC"

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
const UrgencyBlockId = "general_request_urgency"
const UrgencyActionId = "general_request_URGENCY"

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

func BuildGeneralRequestModal() slack.ModalViewRequest {
	// Modal texts
	titleText := slack.NewTextBlockObject(slack.PlainTextType, "Angkor Request", false, false)
	closeText := slack.NewTextBlockObject(slack.PlainTextType, "Cancel", false, false)
	submitText := slack.NewTextBlockObject(slack.PlainTextType, "Submit", false, false)
	// Header section
	headerSection := createHeader()
	// Name input
	titleBlock := createTitleBlock()
	// Category input
	categoryBlock := createCategoryBlock()
	// Description input
	requestDescriptionBlock := createRequestDescriptionBlock()
	// Urgency input
	requestUrgencyBlock := createUrgencyBlock()

	// Blocks
	blocks := slack.Blocks{
		BlockSet: []slack.Block{
			headerSection,
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
	modalRequest.CallbackID = GeneralRequestModalCallbackId
	modalRequest.Type = slack.ViewType("modal")
	modalRequest.Title = titleText
	modalRequest.Close = closeText
	modalRequest.Submit = submitText
	modalRequest.Blocks = blocks
	return modalRequest
}

// const DefaultBotTag = "slack-bot-generalrequest"

func NewGeneralRequestFromModal(i slack.InteractionCallback) (GeneralRequest, []ModalError) {
	stateValues := i.View.State.Values

	title := stateValues[TitleBlockId][TitleActionId].Value

	category := stateValues[CategoryBlockId][CategoryActionId].SelectedOption.Value

	description := stateValues[RequestDescriptionBlockId][RequestDescriptionActionId].Value

	urgency := stateValues[UrgencyBlockId][UrgencyActionId].SelectedOption.Value

	data := GeneralRequest{
		Requester:   i.User.ID,
		Title:       title,
		Category:    category,
		Description: description,
		Urgency:     urgency,
	}

	errs := validateRequest(data)
	return data, errs
}

type ModalError struct {
	BlockID string
	Err     error
	ErrMsg  string
}

func validateRequest(data GeneralRequest) []ModalError {
	errs := make([]ModalError, 0, 10)

	check := func(value, blockId, errMsg string) {
		if value == "" {
			errs = append(errs, ModalError{
				BlockID: blockId,
				Err:     fmt.Errorf(errMsg),
				ErrMsg:  errMsg,
			})
		}
	}
	toCheck := []struct {
		value   string
		blockId string
		errMsg  string
	}{
		{data.Title, TitleBlockId, "Missing Title!"},
		{data.Category, CategoryBlockId, "Missing Category!"},
		{data.Description, RequestDescriptionBlockId, "Missing Description!"},
		{data.Urgency, UrgencyBlockId, "Missing Urgency!"},
	}

	for _, checkData := range toCheck {
		check(checkData.value, checkData.blockId, checkData.errMsg)
	}
	return errs
}
