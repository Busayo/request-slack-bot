package vms

import (

	"github.com/slack-go/slack"
)

func BuildVMRequestModal2() slack.ModalViewRequest {
	// Modal texts
	titleText := slack.NewTextBlockObject(slack.PlainTextType, "VM Request", false, false)
	closeText := slack.NewTextBlockObject(slack.PlainTextType, "Cancel", false, false)
	submitText := slack.NewTextBlockObject(slack.PlainTextType, "Submit", false, false)
	// Header section
	headerSection := createHeader()
	// Name input
	vmNameBlock := createNameBlock()
	// Description input
	vmDescriptionBlock := createDescriptionBlock()
	// Public SSH Key
	vmPublicSSHKey := createSSHKeyBlock()
	vmPublicSSHKeyExplanation := createSSHKeyExplanationBlock()
	// Provider input
	vmProviderBlock := createProviderBlock()
	// Additional inputs
	vmAdditionalBlock := createAdditionalInputBlock()
	// Blocks
	blocks := slack.Blocks{
		BlockSet: []slack.Block{
			headerSection,
			slack.NewDividerBlock(),
			vmNameBlock,
			vmDescriptionBlock,
			vmPublicSSHKey,
			vmPublicSSHKeyExplanation,
			vmProviderBlock,
			vmAdditionalBlock,
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