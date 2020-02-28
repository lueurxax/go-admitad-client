package client

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/lueurxax/go-admitad-client/responses"
	"time"
)

type AClient struct {
	Statistics
	bc *baseClient
}

func New(url string) *AClient {
	bc := newBaseClient(url)
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
}

func newBaseClient(url string) *baseClient {
	return &baseClient{Client: resty.New().SetHostURL(url)}
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
		return nil
	}
	answer := new(responses.Token)
	errResponse := new(responses.ErrorResponse)
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
		return err
	}
	if resp.Error() != nil {
		return errResponse
	}

	b.setToken(answer)

	return nil
}

func (b *baseClient) setToken(token *responses.Token) {
	b.auth = &auth{
		tokenExpiredAt: time.Now().Add(time.Duration(token.ExpiresIN)*time.Second - 5*time.Minute),
		token:          token.AccessToken,
		refreshToken:   token.RefreshToken,
	}
}
