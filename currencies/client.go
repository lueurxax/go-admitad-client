package currencies

import (
	"github.com/lueurxax/go-admitad-client/internal"
	"github.com/lueurxax/go-admitad-client/requests"
	"github.com/lueurxax/go-admitad-client/responses"
)

type Currencies struct {
	*internal.BaseClient
}

// https://www.admitad.com/en/developers/doc/api_ru/methods/public/currency_exchange_rate/
func (c Currencies) Get(request requests.Currency) (*responses.Currency, error) {
	data, errResponse := new(responses.Currency), new(responses.ErrorResponse)

	httpParams, logParams := request.Params()

	c.Logger.Debug("make request", "/currencies/rate/", logParams, nil)

	resp, err := c.R().
		SetQueryParams(httpParams).
		EnableTrace().
		SetAuthToken(c.Auth.Token).
		SetResult(data).
		SetError(errResponse).
		Get("/currencies/rate/")

	if err != nil {
		c.Logger.Error("http error", "/currencies/rate/", logParams, err)
		return nil, err
	}

	c.Metrics.Collect("/currencies/rate/", resp.StatusCode(), errResponse.StatusCode, resp.Time())

	if resp.Error() != nil {
		c.Logger.Error("app error", "/currencies/rate/", errResponse.ErrLogParams(logParams), errResponse)
		return nil, errResponse
	}

	c.Logger.Debug("success request", "/currencies/rate/", logParams, nil)

	return data, nil
}
