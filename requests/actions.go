package requests

import (
	"github.com/lueurxax/go-admitad-client/enums"
	"strconv"
	"time"
)

type Actions struct {
	Offset             *int       `json:"offset"`
	Limit              int        `json:"limit"`
	DateStart          *time.Time `json:"date_start"`
	DateEnd            *time.Time `json:"date_end"`
	ClosingDateStart   *time.Time `json:"closing_date_start"`
	ClosingDateEnd     *time.Time `json:"closing_date_end"`
	StatusUpdatedStart *time.Time `json:"status_updated_start"`
	StatusUpdatedEnd   *time.Time `json:"status_updated_end"`
	Website            int        `json:"website"`
	Campaign           int        `json:"campaign"`
	Subid              string
	Subid1             string
	Subid2             string
	Subid3             string
	Subid4             string
	Source             string
	Status             enums.Status
	Keyword            string
	Action             string
	ActionID           *int
	Banner             string
	ActionType         enums.ActionType
	Processed          enums.Processed
	Paid               enums.Paid
	OrderBy            []string
}

func (a *Actions) Params() (map[string]string, map[string]interface{}) {
	params := map[string]string{}
	logParams := map[string]interface{}{}
	if a == nil {
		return params, logParams
	}
	if a.ActionID != nil {
		params["action_id"] = strconv.Itoa(*a.ActionID)
	}
	if a.Limit != 0 {
		logParams["limit"] = a.Limit
		params["limit"] = strconv.Itoa(a.Limit)
	}
	if a.Offset != nil {
		logParams["offset"] = *a.Offset
		params["offset"] = strconv.Itoa(*a.Offset)
	}
	if a.DateStart != nil {
		logParams["date_start"] = a.DateStart
		params["date_start"] = a.DateStart.Format("02.01.2006")
	}
	return params, logParams
}
