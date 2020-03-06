package me

import (
	"github.com/lueurxax/go-admitad-client/enums"
	"github.com/lueurxax/go-admitad-client/internal"
	"github.com/lueurxax/go-admitad-client/responses"
)

type Me struct {
	*internal.BaseClient
}

// todo: naming, testing, function description
// https://www.admitad.com/ru/developers/doc/api_ru/methods/private/payment-settings/
func (c *Me) Settings() error {
	answer := new(responses.PaymentDataInformation)
	errResponse := new(responses.ErrorResponse)

	resp, err := c.R().EnableTrace().
		SetAuthToken(c.Token).
		SetResult(answer).
		SetError(errResponse).
		Get("/me/payment/settings/")

	if err != nil {
		return err
	}

	if resp.Error() != nil {
		return errResponse
	}

	return nil
}

// todo: naming, testing, function description
// https://www.admitad.com/ru/developers/doc/api_ru/methods/private/payment-settings/
func (c *Me) SettingsByCurrency(currency enums.Currency) error {
	answer := new(responses.PaymentDataInformation)
	errResponse := new(responses.ErrorResponse)

	resp, err := c.R().EnableTrace().
		SetAuthToken(c.Token).
		SetResult(answer).
		SetError(errResponse).
		Get("/me/payment/settings/" + string(currency))

	if err != nil {
		return err
	}

	if resp.Error() != nil {
		return errResponse
	}

	return nil
}

// todo: naming, testing, function description
// https://www.admitad.com/ru/developers/doc/api_ru/methods/private/user-balance/
func (c *Me) Balance() error {
	answer := new(responses.Balance)
	errResponse := new(responses.ErrorResponse)

	resp, err := c.R().EnableTrace().
		SetAuthToken(c.Token).
		SetResult(answer).
		SetError(errResponse).
		Get("/me/balance/")

	if err != nil {
		return err
	}

	if resp.Error() != nil {
		return errResponse
	}

	return nil
}

// todo: naming, testing, function description
// https://www.admitad.com/ru/developers/doc/api_ru/methods/private/user-balance/
func (c *Me) BalanceExtended() error {
	answer := new(responses.Balance)
	errResponse := new(responses.ErrorResponse)

	resp, err := c.R().EnableTrace().
		SetAuthToken(c.Token).
		SetResult(answer).
		SetError(errResponse).
		Get("/me/balance/extended/")

	if err != nil {
		return err
	}

	if resp.Error() != nil {
		return errResponse
	}

	return nil
}

// todo: naming, testing, function description
// https://www.admitad.com/ru/developers/doc/api_ru/methods/private/user/
func (c *Me) Me() error {
	answer := new(responses.Me)
	errResponse := new(responses.ErrorResponse)

	resp, err := c.R().EnableTrace().
		SetAuthToken(c.Token).
		SetResult(answer).
		SetError(errResponse).
		Get("/me/")

	if err != nil {
		return err
	}

	if resp.Error() != nil {
		return errResponse
	}

	return nil
}
