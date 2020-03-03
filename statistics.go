package client

import (
	"fmt"
	"github.com/lueurxax/go-admitad-client/requests"
	"github.com/lueurxax/go-admitad-client/responses"
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
	data, errResponse := new(responses.Actions), new(responses.ErrorResponse)

	httpParams, logParams := request.Params()

	s.logger.Debug("make request", "/statistics/actions/", logParams, nil)

	resp, err := s.R().
		SetQueryParams(httpParams).
		EnableTrace().
		SetAuthToken(s.auth.token).
		SetResult(data).
		SetError(errResponse).
		Get("/statistics/actions/")

	if err != nil {
		s.logger.Error("http error", "/statistics/actions/", logParams, err)
		return nil, err
	}

	s.metrics.Collect("/statistics/actions/", resp.StatusCode(), errResponse.StatusCode, resp.Time())

	if resp.Error() != nil {
		s.logger.Error("app error", "/statistics/actions/", errResponse.ErrLogParams(logParams), errResponse)
		return nil, errResponse
	}

	s.logger.Debug("success request", "/statistics/actions/", logParams, nil)

	return data, nil
}

func (s *Statistics) ActionsTotal(request requests.Actions) (*responses.ActionsTotal, error) {
	panic("method not implemented")
}
