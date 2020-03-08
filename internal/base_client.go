package internal

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/lueurxax/go-admitad-client/responses"
	"time"
)

// msg - human readable message, url - using url path, params - map of parameters in request, err - error
type Logger interface {
	Panic(msg string, url string, params map[string]interface{}, err error)
	Error(msg string, url string, params map[string]interface{}, err error)
	Warn(msg string, url string, params map[string]interface{}, err error)
	Info(msg string, url string, params map[string]interface{}, err error)
	Debug(msg string, url string, params map[string]interface{}, err error)
	Trace(msg string, url string, params map[string]interface{}, err error)
}

type MetricsCollector interface {
	Collect(url string, httpCode int, appCode int, duration time.Duration)
}

type auth struct {
	tokenExpiredAt time.Time
	Token          string
	refreshToken   string
}

type BaseClient struct {
	*resty.Client
	Auth     *auth
	clientID string
	scope    string

	Logger
	Metrics MetricsCollector
}

func NewBaseClient(url string, logger Logger, metrics MetricsCollector) *BaseClient {
	return &BaseClient{Client: resty.New().SetHostURL(url), Logger: logger, Metrics: metrics}
}

func (b *BaseClient) AutoRefresh(force bool) error {
	if b.Auth == nil {
		return fmt.Errorf("need to use SetAuth to set auth data")
	}

	if b.Auth.tokenExpiredAt.Sub(time.Now()) > 0 && !force {
		b.Logger.Debug("the token is up-to-date", "/token/", nil, nil)
		return nil
	}
	answer, errResponse := new(responses.Token), new(responses.ErrorResponse)

	logParams := map[string]interface{}{"force": force}

	b.Logger.Debug("make request", "/token/", logParams, nil)

	resp, err := b.R().EnableTrace().
		SetFormData(map[string]string{
			"grant_type":    "refresh_token",
			"client_id":     b.clientID,
			"scope":         b.scope,
			"refresh_token": b.Auth.refreshToken,
		}).
		SetResult(answer).
		SetError(errResponse).
		Post("/token/")

	if err != nil {
		b.Logger.Error("http error", "/statistics/actions/", logParams, err)
		return err
	}

	b.Metrics.Collect("/statistics/actions/", resp.StatusCode(), errResponse.StatusCode, resp.Time())

	if resp.Error() != nil {
		b.Logger.Error("app error", "/statistics/actions/", errResponse.ErrLogParams(logParams), errResponse)
		return errResponse
	}

	b.setToken(answer)

	b.Logger.Debug("success response", "/token/", nil, nil)

	return nil
}

func (b *BaseClient) setToken(token *responses.Token) {
	b.Auth = &auth{
		tokenExpiredAt: time.Now().Add(time.Duration(token.ExpiresIn)*time.Second - 5*time.Minute),
		Token:          token.AccessToken,
		refreshToken:   token.RefreshToken,
	}
}

func (b *BaseClient) SetAuth(clientID string, secret string, scope string) error {
	b.clientID = clientID
	b.scope = scope
	answer := new(responses.Token)
	errResponse := new(responses.ErrorResponse)
	resp, err := b.R().EnableTrace().
		SetBasicAuth(clientID, secret).
		SetFormData(map[string]string{
			"grant_type": "client_credentials",
			"client_id":  clientID,
			"scope":      scope,
		}).
		SetResult(answer).
		SetError(errResponse).
		Post("/token/")
	if err != nil {
		return err
	}

	if resp.Error() != nil {
		return errResponse
	}

	b.setToken(answer)
	return nil
}

func IsAuthError(err error) bool {
	switch e := err.(type) {
	case *responses.ErrorResponse:
		if e.StatusCode == 401 {
			return true
		}
	}
	return false
}
