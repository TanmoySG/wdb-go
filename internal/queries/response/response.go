package apiResponse

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
