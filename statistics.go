package client

import (
	"fmt"
	"github.com/lueurxax/go-admitad-client/internal"
	"github.com/lueurxax/go-admitad-client/requests"
	"github.com/lueurxax/go-admitad-client/responses"
)

type Statistics struct {
	b *internal.BaseClient
}

func (s *Statistics) Actions(request *requests.Actions) (*responses.Actions, error) {
	if err := s.b.AutoRefresh(false); err != nil {
		return nil, err
	}
	resp, err := s.actions(request)
	if err != nil {
		if !IsAuthError(err) {
			return nil, err
		}
		if err2 := s.b.AutoRefresh(true); err2 != nil {
			return nil, fmt.Errorf("request err: %s, refresh token err: %s", err.Error(), err2.Error())
		}
		return s.actions(request)
	}
	return resp, nil
}

func (s *Statistics) actions(request *requests.Actions) (*responses.Actions, error) {
	data, errResponse := new(responses.Actions), new(responses.ErrorResponse)

	httpParams, logParams := request.Params()

	s.b.Logger.Debug("make request", "/statistics/actions/", logParams, nil)

	resp, err := s.b.R().
		SetQueryParams(httpParams).
		EnableTrace().
		SetAuthToken(s.b.Auth.Token).
		SetResult(data).
		SetError(errResponse).
		Get("/statistics/actions/")

	if err != nil {
		s.b.Logger.Error("http error", "/statistics/actions/", logParams, err)
		return nil, err
	}

	s.b.Metrics.Collect("/statistics/actions/", resp.StatusCode(), errResponse.StatusCode, resp.Time())

	if resp.Error() != nil {
		s.b.Logger.Error("app error", "/statistics/actions/", errResponse.ErrLogParams(logParams), errResponse)
		return nil, errResponse
	}

	s.b.Logger.Debug("success request", "/statistics/actions/", logParams, nil)

	return data, nil
}

func (s *Statistics) ActionsTotal(request requests.Actions) (*responses.ActionsTotal, error) {
	panic("method not implemented")
}
