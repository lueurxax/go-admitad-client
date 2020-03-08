package requests

import (
	"strconv"
	"time"
)

type OfflineReceipts struct {
	WebsiteID           *int
	ID                  string
	DatetimeCreatedFrom *time.Time
	DatetimeCreatedTo   *time.Time
	DateTimeUpdatedFrom *time.Time
	DateTimeUpdatedTo   *time.Time
	Limit               int
	Offset              *int
	PurchaserID         string
}

func (r *OfflineReceipts) Params() (map[string]string, map[string]interface{}) {
	params := map[string]string{}
	logParams := map[string]interface{}{}
	if r == nil {
		return params, logParams
	}

	if r.WebsiteID != nil {
		logParams["website_id"] = r.WebsiteID
		params["website_id"] = strconv.Itoa(*r.WebsiteID)
	}

	if r.ID != "" {
		logParams["id"] = r.ID
		params["id"] = r.ID
	}

	if r.Limit != 0 {
		logParams["limit"] = r.ID
		params["limit"] = strconv.Itoa(r.Limit)
	}

	if r.Offset != nil {
		logParams["offset"] = r.Offset
		params["offset"] = strconv.Itoa(*r.Offset)
	}

	if r.PurchaserID != "" {
		logParams["purchaser_id"] = r.PurchaserID
		params["purchaser_id"] = r.PurchaserID
	}

	if r.DatetimeCreatedFrom != nil {
		logParams["datetime_created_from"] = r.DatetimeCreatedFrom
		params["datetime_created_from"] = r.DatetimeCreatedFrom.Format("02.01.2006 15:04")
	}

	if r.DatetimeCreatedTo != nil {
		logParams["datetime_created_to"] = r.DatetimeCreatedTo
		params["datetime_created_to"] = r.DatetimeCreatedTo.Format("02.01.2006 15:04")
	}

	if r.DateTimeUpdatedFrom != nil {
		logParams["datetime_updated_from"] = r.DateTimeUpdatedFrom
		params["datetime_updated_from"] = r.DateTimeUpdatedFrom.Format("02.01.2006 15:04")
	}

	if r.DateTimeUpdatedTo != nil {
		logParams["datetime_updated_to"] = r.DateTimeUpdatedTo
		params["datetime_updated_to"] = r.DateTimeUpdatedTo.Format("02.01.2006 15:04")
	}

	return params, logParams
}
