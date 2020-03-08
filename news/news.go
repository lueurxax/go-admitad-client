package news

import (
	"github.com/lueurxax/go-admitad-client/internal"
	"github.com/lueurxax/go-admitad-client/requests"
	"github.com/lueurxax/go-admitad-client/responses"
	"strconv"
)

type News struct {
	// *client.AClient
	*internal.BaseClient
}

// todo: naming, testing, function description
// https://www.admitad.com/ru/developers/doc/api_ru/methods/news/news/
func (c *News) ByID(id int) error {
	answer := new(responses.News)
	errResponse := new(responses.ErrorResponse)

	resp, err := c.R().EnableTrace().
		SetAuthToken(c.Token).
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
func (c *News) Get(news requests.News) error {
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

	resp, err := c.R().EnableTrace().
		SetAuthToken(c.Token).
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
