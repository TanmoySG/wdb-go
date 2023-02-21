package apiResponse

import (
	"encoding/json"
	"fmt"
)

const (
	StatusSuccess = "success"
	StatusFailure = "failure"
)

type Response struct {
	Action string       `json:"action"`
	Status string       `json:"status"`
	Error  *Error       `json:"error,omitempty"`
	Data   *interface{} `json:"data,omitempty"`
}

type Error struct {
	Code  string     `json:"code,omitempty"`
	Stack ErrorStack `json:"stack,omitempty"`
}

type ErrorStack []string

func (r Response) IsSuccess() bool {
	switch r.Status {
	case StatusSuccess:
		return true
	case StatusFailure:
		return false
	default:
		return false
	}
}

func (r Response) MarshalData() ([]byte, error) {
	dataByteArray, err := json.Marshal(r.Data)
	if err != nil {
		return nil, fmt.Errorf("error marshalling data map : %s", err)
	}
	return dataByteArray, nil
}
