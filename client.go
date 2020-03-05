package client

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/lueurxax/go-admitad-client/defaults"
	"github.com/lueurxax/go-admitad-client/news"
	"github.com/lueurxax/go-admitad-client/referrals"
	"github.com/lueurxax/go-admitad-client/responses"
	"github.com/lueurxax/go-admitad-client/tickets"
	"time"
)

type AClient struct {
	Statistics
	news.News
	referrals.Referrals
	tickets.Ticket

	bc *baseClient
}

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

func New(url string, logger Logger, mc MetricsCollector) *AClient {
	if logger == nil {
		logger = defaults.NewLogrus()
	}
	if mc == nil {
		mc = &defaults.EmptyMetrics{}
	}
	bc := newBaseClient(url, logger, mc)
	return &AClient{Statistics: Statistics{bc}, bc: bc}
}

type auth struct {
	tokenExpiredAt time.Time
	token          string
	refreshToken   string
}

type baseClient struct {
	*resty.Client
	auth     *auth
	clientID string
	scope    string

	logger  Logger
	metrics MetricsCollector
}

func newBaseClient(url string, logger Logger, metrics MetricsCollector) *baseClient {
	return &baseClient{Client: resty.New().SetHostURL(url), logger: logger, metrics: metrics}
}

func (ac *AClient) SetAuth(clientID, clientSecret, scope string) error {
	ac.clientID = clientID
	ac.scope = scope
	answer := new(responses.Token)
	errResponse := new(responses.ErrorResponse)
	resp, err := ac.R().EnableTrace().
		SetBasicAuth(clientID, clientSecret).
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

	ac.setToken(answer)

	return nil
}

func (b *baseClient) autoRefresh(force bool) error {
	if b.auth == nil {
		return fmt.Errorf("need use SetAuth for set auth data")
	}

	if b.auth.tokenExpiredAt.Sub(time.Now()) > 0 && !force {
		b.logger.Debug("token is actual", "/token/", nil, nil)
		return nil
	}
	answer, errResponse := new(responses.Token), new(responses.ErrorResponse)

	logParams := map[string]interface{}{"force": force}

	b.logger.Debug("make request", "/token/", logParams, nil)

	resp, err := b.R().EnableTrace().
		SetFormData(map[string]string{
			"grant_type":    "refresh_token",
			"client_id":     b.clientID,
			"scope":         b.scope,
			"refresh_token": b.auth.refreshToken,
		}).
		SetResult(answer).
		SetError(errResponse).
		Post("/token/")

	if err != nil {
		b.logger.Error("http error", "/statistics/actions/", logParams, err)
		return err
	}

	b.metrics.Collect("/statistics/actions/", resp.StatusCode(), errResponse.StatusCode, resp.Time())

	if resp.Error() != nil {
		b.logger.Error("app error", "/statistics/actions/", errResponse.ErrLogParams(logParams), errResponse)
		return errResponse
	}

	b.setToken(answer)

	b.logger.Debug("success response", "/token/", nil, nil)

	return nil
}

func (b *baseClient) setToken(token *responses.Token) {
	b.auth = &auth{
		tokenExpiredAt: time.Now().Add(time.Duration(token.ExpiresIN)*time.Second - 5*time.Minute),
		token:          token.AccessToken,
		refreshToken:   token.RefreshToken,
	}
}
