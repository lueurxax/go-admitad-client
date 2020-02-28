package responses

import "fmt"

type ErrorResponse struct {
	ErrorDescription string `json:"error_description"`
	StatusCode       int    `json:"status_code"`
	ErrorCode        int    `json:"error_code"`
	ErrorField       string `json:"error"`
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("api error: %s", e.ErrorDescription)
}
