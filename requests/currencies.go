package requests

import (
	"github.com/lueurxax/go-admitad-client/enums"
	"time"
)

type Currency struct {
	Base   *enums.Currency
	Target *enums.Currency
	Date   *time.Time
}

func (c *Currency) Params() (map[string]string, map[string]interface{}) {
	params := map[string]string{}
	logParams := map[string]interface{}{}
	if c == nil {
		return params, logParams
	}

	if c.Date != nil {
		logParams["date"] = c.Date
		params["date"] = c.Date.Format("02.01.2006")
	}

	if c.Base != nil {
		logParams["base"] = c.Base
		params["base"] = string(*c.Base)
	}

	if c.Target != nil {
		logParams["target"] = c.Target
		params["target"] = string(*c.Target)
	}

	return params, logParams
}
