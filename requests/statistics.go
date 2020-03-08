package requests

import (
	"github.com/lueurxax/go-admitad-client/enums"
	"strconv"
	"time"
)

// https://www.admitad.com/en/developers/doc/api_ru/methods/statistics/statistics-sub-id/
type StatisticsSubID struct {
	Offset      *int
	Limit       int
	DateStart   *time.Time
	DateEnd     *time.Time
	Website     *int
	Campaign    *int
	SubID       string
	SubID1      string
	SubID2      string
	SubID3      string
	SubID4      string
	OrderBy     enums.Sorting
	Total       *bool
	GroupSubIDs *int
}

func (r *StatisticsSubID) Params() (map[string]string, map[string]interface{}) {
	params := map[string]string{}
	logParams := map[string]interface{}{}
	if r == nil {
		return params, logParams
	}

	if r.Limit != 0 {
		logParams["limit"] = r.Limit
		params["limit"] = strconv.Itoa(r.Limit)
	}

	if r.Offset != nil {
		logParams["offset"] = r.Offset
		params["offset"] = strconv.Itoa(*r.Offset)
	}

	if r.Campaign != nil {
		logParams["campaign"] = r.Campaign
		params["campaign"] = strconv.Itoa(*r.Campaign)
	}

	if r.GroupSubIDs != nil {
		logParams["group_siubids"] = r.GroupSubIDs
		params["group_subids"] = strconv.Itoa(*r.GroupSubIDs)
	}

	if r.SubID != "" {
		logParams["subid"] = r.SubID
		params["subid"] = r.SubID
	}

	if r.Total != nil {
		logParams["total"] = r.Total
		params["total"] = strconv.FormatBool(*r.Total)
	}

	if r.Website != nil {
		logParams["website"] = r.Website
		params["website"] = strconv.Itoa(*r.Website)
	}

	if r.DateEnd != nil {
		logParams["date_end"] = r.DateEnd
		params["date_end"] = r.DateEnd.Format("02.01.2006")
	}

	if r.DateStart != nil {
		logParams["date_start"] = r.DateStart
		params["date_start"] = r.DateStart.Format("02.01.2006")
	}
	return params, logParams
}
