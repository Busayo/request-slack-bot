package vms

import (

	"github.com/slack-go/slack"
)

const RequestModalCallbackId = "request-modal"

func createHeader() *slack.SectionBlock {
	return slack.NewSectionBlock(
		slack.NewTextBlockObject(
			slack.MarkdownType,
			"*Please fill out the request form below:*",
			false,
			false,
		),
		nil,
		nil,
	)
}

// Name block data
const VMNameBlockId = "vm_name"
const VMNameActionId = "VM_NAME"

func createNameBlock() *slack.InputBlock {
	return slack.NewInputBlock(
		VMNameBlockId,
		slack.NewTextBlockObject(
			slack.PlainTextType,
			"Project Name",
			false,
			false,
		),
		slack.NewTextBlockObject(
			slack.PlainTextType,
			"Project Name",
			false,
			false,
		),
		slack.NewPlainTextInputBlockElement(
			nil,
			VMNameActionId,
		), 
	)
}


// Description block data
const VMDescriptionBlockId = "vm_desc"
const VMDescriptionActionId = "VM_DESC"

func createDescriptionBlock() *slack.InputBlock {
	textInputBlock := slack.NewPlainTextInputBlockElement(
		nil,
		VMDescriptionActionId,
	)
	textInputBlock.Multiline = true
	return slack.NewInputBlock(
		VMDescriptionBlockId,
		slack.NewTextBlockObject(
			slack.PlainTextType,
			"Why do you need this VM?",
			false,
			false,
		),
		slack.NewTextBlockObject(
			slack.PlainTextType,
			"Why do you need this VM?",
			false,
			false,
		),
		textInputBlock,
	)
}


// SSH Key block data
const VMSSHKeyBlockId = "ssh_key"
const VMSSHKeyActionId = "SSH_KEY"

func createSSHKeyBlock() *slack.InputBlock {
	return slack.NewInputBlock(
		VMSSHKeyBlockId,
		slack.NewTextBlockObject(
			slack.PlainTextType,
			"Public SSH Key",
			false,
			false,
		),
		slack.NewTextBlockObject(
			slack.PlainTextType,
			"Public SSH Key",
			false,
			false,
		),
		slack.NewPlainTextInputBlockElement(
			nil,
			VMSSHKeyActionId,
		),
	)
}

func createSSHKeyExplanationBlock() *slack.SectionBlock {
	return slack.NewSectionBlock(
		slack.NewTextBlockObject(
			slack.MarkdownType,
			"_For How to create a Public SSH Key or other details read <https://www.linode.com/docs/guides/use-public-key-authentication-with-ssh/#public-key-authentication-on-linux-and-macos|here>_",
			false,
			false,
		),
		nil,
		nil,
	)
}


// VM provider block data
const VMProviderBlockId = "vm_provider"
const VMProviderActionId = "VM_PROVIDER"

func createProviderOptions() []*slack.OptionBlockObject {
	optionBlockObjects := make([]*slack.OptionBlockObject, 0, len(SupportedProviders))
	for _, option := range SupportedProviders {
		optionText := slack.NewTextBlockObject(slack.PlainTextType, option.Name, false, false)
		optionBlockObjects = append(optionBlockObjects, slack.NewOptionBlockObject(option.Value, optionText, nil))
	}
	return optionBlockObjects
}

func createProviderBlock() *slack.InputBlock {
	providerInput := slack.NewInputBlock(
		VMProviderBlockId,
		slack.NewTextBlockObject(
			slack.PlainTextType,
			"Select a Provider",
			false,
			false,
		),
		slack.NewTextBlockObject(
			slack.PlainTextType,
			"Select a Provider",
			false,
			false,
		),
		slack.NewOptionsSelectBlockElement(
			slack.OptTypeStatic,
			nil,
			VMProviderActionId,
			createProviderOptions()...,
		),
	)
	providerInput.DispatchAction = true
	return providerInput
}


// VM Additional block data
const VMAdditionalBlockId = "vm_additional"
const VMAdditionalActionId = "VM_ADDITIONAL"

const VMPrivateIPValue = "USE_PRIVATE_IP"

func createPrivateIpBlock() *slack.OptionBlockObject {
	return slack.NewOptionBlockObject(
		VMPrivateIPValue,
		slack.NewTextBlockObject(
			slack.PlainTextType,
			"Use Private Ip",
			false,
			false,
		),
		nil,
	)
}

func createAdditionalInputBlock() *slack.SectionBlock {
	section := slack.NewSectionBlock(
		slack.NewTextBlockObject(
			slack.MarkdownType,
			"*Additional*",
			false,
			false,
		),
		nil,
		slack.NewAccessory(
			slack.NewCheckboxGroupsBlockElement(
				VMAdditionalActionId,
				createPrivateIpBlock(),
			),
		),
	)
	section.BlockID = VMAdditionalBlockId
	return section
}

func BuildVMRequestModal() slack.ModalViewRequest {
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


