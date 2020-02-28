package client

import "github.com/lueurxax/go-admitad-client/responses"

func IsAuthError(err error) bool {
	switch e := err.(type) {
	case *responses.ErrorResponse:
		if e.StatusCode == 401 {
			return true
		}
	}
	return false
}
