package client

import (
	"fmt"
	"github.com/lueurxax/go-admitad-client/requests"
	"github.com/lueurxax/go-admitad-client/responses"
	"strconv"
)

type Statistics struct {
	*baseClient
}

func (s *Statistics) Actions(request *requests.Actions) (*responses.Actions, error) {
	if err := s.baseClient.autoRefresh(false); err != nil {
		return nil, err
	}
	resp, err := s.actions(request)
	if err != nil {
		if !IsAuthError(err) {
			return nil, err
		}
		if err2 := s.autoRefresh(true); err2 != nil {
			return nil, fmt.Errorf("request err: %s, refresh token err: %s", err.Error(), err2.Error())
		}
		return s.actions(request)
	}
	return resp, nil
}

func (s *Statistics) actions(request *requests.Actions) (*responses.Actions, error) {
	data := new(responses.Actions)
	errResponse := new(responses.ErrorResponse)
	params := map[string]string{}
	if request.Limit != 0 {
		params["limit"] = strconv.Itoa(request.Limit)
	}
	if request.Offset != nil {
		params["offset"] = strconv.Itoa(*request.Offset)
	}
	if request.DateStart != nil {
		params["date_start"] = request.DateStart.Format("02.01.2006")
	}
	resp, err := s.R().
		SetQueryParams(params).
		EnableTrace().
		SetAuthToken(s.auth.token).
		SetResult(data).
		SetError(errResponse).
		Get("/statistics/actions/")
	if err != nil {
		return nil, err
	}
	if resp.Error() != nil {
		return nil, errResponse
	}
	return data, nil
}

func (s *Statistics) ActionsTotal(request requests.Actions) (*responses.ActionsTotal, error) {
	panic("method not implemented")
}
