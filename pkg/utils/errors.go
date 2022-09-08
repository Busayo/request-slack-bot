package utils

import (
	"encoding/json"

	"github.com/request-slack-bot/pkg/generalrequests"
//	"github.com/request-slack-bot/pkg/vms"
)

type responseError struct {
	Action string            `json:"response_action"`
	Errs   map[string]string `json:"errors"`
}




func BuildGeneralResponseErrorsBody(modalErrs []generalrequests.ModalError) (string, error) {
	resp := responseError{
		Action: "errors",
		Errs:   map[string]string{},
	}

	for _, modalErr := range modalErrs {
		resp.Errs[modalErr.BlockID] = modalErr.ErrMsg
	}

	body, err := json.Marshal(resp)
	return string(body), err
}