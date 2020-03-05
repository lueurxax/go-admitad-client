package responses

import "fmt"

type ErrorResponse struct {
	ErrorDescription string `json:"error_description"`
	StatusCode       int    `json:"status_code"`
	ErrorCode        int    `json:"error_code"`
	ErrorField       string `json:"error"`
	Detail           string `json:"detail"`
}

func (e *ErrorResponse) Error() string {
	if e.Detail != "" {
		return fmt.Sprintf("api error: %s", e.Detail)
	}
	return fmt.Sprintf("api error: %s", e.ErrorDescription)
}

func (e *ErrorResponse) ErrLogParams(params map[string]interface{}) map[string]interface{} {
	errParams := params
	errParams["status_code"] = e.StatusCode
	errParams["error_code"] = e.ErrorCode
	return errParams
}
