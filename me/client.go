package me

import (
	client "github.com/lueurxax/go-admitad-client"
	"github.com/lueurxax/go-admitad-client/enums"
	"github.com/lueurxax/go-admitad-client/responses"
)

type Payment struct {
	baseClient *client.AClient
}

// todo: naming, testing, function description
// https://www.admitad.com/ru/developers/doc/api_ru/methods/private/payment-settings/
func (c *Payment) Settings() error {
	answer := new(responses.PaymentDataInformation)
	errResponse := new(responses.ErrorResponse)

	resp, err := c.baseClient.R().EnableTrace().
		SetAuthToken(c.baseClient.Token).
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
func (c *Payment) SettingsByCurrency(currency enums.Currency) error {
	answer := new(responses.PaymentDataInformation)
	errResponse := new(responses.ErrorResponse)

	resp, err := c.baseClient.R().EnableTrace().
		SetAuthToken(c.baseClient.Token).
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
func (c *Payment) Balance() error {
	answer := new(responses.Balance)
	errResponse := new(responses.ErrorResponse)

	resp, err := c.baseClient.R().EnableTrace().
		SetAuthToken(c.baseClient.Token).
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
func (c *Payment) BalanceExtended() error {
	answer := new(responses.Balance)
	errResponse := new(responses.ErrorResponse)

	resp, err := c.baseClient.R().EnableTrace().
		SetAuthToken(c.baseClient.Token).
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
func (c *Payment) Me() error {
	answer := new(responses.Me)
	errResponse := new(responses.ErrorResponse)

	resp, err := c.baseClient.R().EnableTrace().
		SetAuthToken(c.baseClient.Token).
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
