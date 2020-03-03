package referrals

import (
	client "github.com/lueurxax/go-admitad-client"
	"github.com/lueurxax/go-admitad-client/requests"
	"github.com/lueurxax/go-admitad-client/responses"
	"strconv"
)

type Referrals struct {
	baseClient *client.AClient
}

// todo: naming, testing, function description
// https://www.admitad.com/ru/developers/doc/api_ru/methods/referrals/referrals/
func (c *Referrals) ReferralsByID(id int) error {
	answer := new(responses.News)
	errResponse := new(responses.ErrorResponse)

	resp, err := c.baseClient.R().EnableTrace().
		SetAuthToken(c.baseClient.Token).
		SetResult(answer).
		SetError(errResponse).
		Get("/news/" + strconv.Itoa(id))

	if err != nil {
		return err
	}

	if resp.Error() != nil {
		return errResponse
	}

	return nil
}

// todo: naming, testing, function description
// https://www.admitad.com/ru/developers/doc/api_ru/methods/referrals/referrals/
func (c *Referrals) Referrals(news requests.Referrals) error {
	answer := new(responses.News)
	errResponse := new(responses.ErrorResponse)

	params := map[string]string{}

	if news.Offset != 0 {
		params["offset"] = strconv.Itoa(news.Offset)
	}

	if news.Limit != 0 {
		params["limit"] = strconv.Itoa(news.Limit)
	}

	if news.DateStart != nil {
		params["date_start"] = news.DateStart.Format("02.01.2006")
	}

	if news.DateEnd != nil {
		params["date_start"] = news.DateEnd.Format("02.01.2006")
	}

	resp, err := c.baseClient.R().EnableTrace().
		SetQueryParams(params).
		SetAuthToken(c.baseClient.Token).
		SetResult(answer).
		SetError(errResponse).
		Get("/news/")

	if err != nil {
		return err
	}

	if resp.Error() != nil {
		return errResponse
	}

	return nil
}
