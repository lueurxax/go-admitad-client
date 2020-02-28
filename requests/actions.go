package requests

import (
	"github.com/lueurxax/go-admitad-client/enums"
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
	ActionID           int
	Banner             string
	ActionType         enums.ActionType
	Processed          enums.Processed
	Paid               enums.Paid
	OrderBy            []string
}
