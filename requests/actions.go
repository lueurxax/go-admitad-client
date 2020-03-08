package requests

import (
	"github.com/lueurxax/go-admitad-client/enums"
	"strconv"
	"strings"
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
	Website            *int       `json:"website"`
	Campaign           *int       `json:"campaign"`
	Subid              string
	Subid1             string
	Subid2             string
	Subid3             string
	Subid4             string
	Source             string
	Status             *enums.Status
	Keyword            string
	Action             string
	ActionID           *int
	Banner             string
	ActionType         *enums.ActionType
	Processed          *enums.Processed
	Paid               *enums.Paid
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

	if a.Website != nil {
		logParams["website"] = a.Website
		params["website"] = strconv.Itoa(*a.Website)
	}

	if a.Campaign != nil {
		logParams["campaign"] = a.Campaign
		params["campaign"] = strconv.Itoa(*a.Campaign)
	}

	if a.Paid != nil {
		logParams["paid"] = a.Paid
		params["paid"] = strconv.Itoa(int(*a.Paid))
	}

	if a.ActionType != nil {
		logParams["action_type"] = a.Paid
		params["action_type"] = string(*a.Paid)
	}

	if a.Status != nil {
		logParams["status"] = a.Status
		params["status"] = string(*a.Status)
	}

	if len(a.OrderBy) > 0 {
		temp := strings.Join(a.OrderBy, ",")
		logParams["order_by"] = temp
		params["order_by"] = temp
	}

	if a.Source != "" {
		logParams["source"] = a.Source
		params["source"] = a.Source
	}

	if a.Keyword != "" {
		logParams["keyword"] = a.Keyword
		params["keyword"] = a.Keyword
	}

	if a.Action != "" {
		logParams["action"] = a.Action
		params["action"] = a.Action
	}

	if a.Banner != "" {
		logParams["banner"] = a.Banner
		params["banner"] = a.Banner
	}

	if a.DateEnd != nil {
		logParams["date_end"] = a.DateEnd
		params["date_end"] = a.DateEnd.Format("02.01.2006")
	}

	if a.DateStart != nil {
		logParams["date_start"] = a.DateStart
		params["date_start"] = a.DateStart.Format("02.01.2006")
	}
	return params, logParams
}

func (a *Actions) ParamsTotal() (map[string]string, map[string]interface{}) {
	params := map[string]string{}
	logParams := map[string]interface{}{}
	if a == nil {
		return params, logParams
	}

	// forced
	params["total"] = "1"
	logParams["total"] = 1

	if a.Limit != 0 {
		logParams["limit"] = a.Limit
		params["limit"] = strconv.Itoa(a.Limit)
	}
	if a.Offset != nil {
		logParams["offset"] = *a.Offset
		params["offset"] = strconv.Itoa(*a.Offset)
	}

	if a.Website != nil {
		logParams["website"] = a.Website
		params["website"] = strconv.Itoa(*a.Website)
	}

	if a.Campaign != nil {
		logParams["campaign"] = a.Campaign
		params["campaign"] = strconv.Itoa(*a.Campaign)
	}

	if a.Paid != nil {
		logParams["paid"] = a.Paid
		params["paid"] = strconv.Itoa(int(*a.Paid))
	}

	if a.ActionType != nil {
		logParams["action_type"] = a.Paid
		params["action_type"] = string(*a.Paid)
	}

	if a.Status != nil {
		logParams["status"] = a.Status
		params["status"] = string(*a.Status)
	}

	if len(a.OrderBy) > 0 {
		temp := strings.Join(a.OrderBy, ",")
		logParams["order_by"] = temp
		params["order_by"] = temp
	}

	if a.Source != "" {
		logParams["source"] = a.Source
		params["source"] = a.Source
	}

	if a.Keyword != "" {
		logParams["keyword"] = a.Keyword
		params["keyword"] = a.Keyword
	}

	if a.Action != "" {
		logParams["action"] = a.Action
		params["action"] = a.Action
	}

	if a.Banner != "" {
		logParams["banner"] = a.Banner
		params["banner"] = a.Banner
	}

	if a.DateEnd != nil {
		logParams["date_end"] = a.DateEnd
		params["date_end"] = a.DateEnd.Format("02.01.2006")
	}

	if a.DateStart != nil {
		logParams["date_start"] = a.DateStart
		params["date_start"] = a.DateStart.Format("02.01.2006")
	}

	return params, logParams
}
