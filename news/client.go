package news

import (
	client "github.com/lueurxax/go-admitad-client"
	"github.com/lueurxax/go-admitad-client/requests"
	"github.com/lueurxax/go-admitad-client/responses"
	"strconv"
)

type News struct {
	baseClient *client.AClient
}

// todo: naming, testing, function description
// https://www.admitad.com/ru/developers/doc/api_ru/methods/news/news/
func (c *News) NewsByID(id int) error {
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
// https://www.admitad.com/ru/developers/doc/api_ru/methods/news/news/
func (c *News) News(news requests.News) error {
	answer := new(responses.News)
	errResponse := new(responses.ErrorResponse)

	params := map[string]string{}

	if news.Offset != 0 {
		params["offset"] = strconv.Itoa(news.Offset)
	}

	if news.Limit != 0 {
		params["limit"] = strconv.Itoa(news.Limit)
	}

	if news.Language != "" {
		params["language"] = string(news.Language)
	}

	resp, err := c.baseClient.R().EnableTrace().
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
