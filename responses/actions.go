package responses

import (
	"encoding/json"
	"github.com/lueurxax/go-admitad-client/enums"
	"strconv"
	"time"
)

type Result struct {
	Status           enums.Status     `json:"status"`
	Comment          string           `json:"comment"`
	ConversionTime   int              `json:"conversion_time"`
	Keyword          string           `json:"keyword"`
	AdvcampaignName  string           `json:"advcampaign_name"`
	AdvcampaignID    int              `json:"advcampaign_id"`
	Cart             float64          `json:"cart"`
	Subid            string           `json:"subid"`
	Subid1           string           `json:"subid1"`
	Subid2           string           `json:"subid2"`
	Subid3           string           `json:"subid3"`
	Subid4           string           `json:"subid4"`
	ClickUserReferer string           `json:"click_user_referer"`
	Currency         enums.Currency   `json:"currency"`
	ActionDate       *time.Time       `json:"action_date"`
	WebsiteName      string           `json:"website_name"`
	Action           string           `json:"action"`
	ClickDate        *time.Time       `json:"click_date"`
	Payment          float64          `json:"payment"`
	TariffID         int              `json:"tariff_id"`
	BannerID         int              `json:"banner_id"`
	ActionID         int              `json:"action_id"`
	Processed        enums.Processed  `json:"processed"`
	Paid             enums.Paid       `json:"paid"`
	OrderID          string           `json:"order_id"`
	ClosingDate      *time.Time       `json:"closing_date"`
	ActionType       enums.ActionType `json:"action_type"`
	StatusUpdated    *time.Time       `json:"status_updated"`
	Positions        []Position       `json:"positions"`
}

func (r *Result) UnmarshalJSON(data []byte) error {
	type Alias Result
	aux := &struct {
		ActionDate    string `json:"action_date"`
		ClickDate     string `json:"click_date"`
		ClosingDate   string `json:"closing_date"`
		StatusUpdated string `json:"status_updated"`
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	if aux.ClosingDate != "" {
		closingDate, err := time.Parse("2006-01-02", aux.ClosingDate)
		if err != nil {
			return err
		}
		r.ClosingDate = &closingDate
	}

	if aux.ActionDate != "" {
		actionDate, err := time.Parse("2006-01-02 15:04:05", aux.ActionDate)
		if err != nil {
			return err
		}
		r.ActionDate = &actionDate
	}

	if aux.ClickDate != "" {
		clickDate, err := time.Parse("2006-01-02 15:04:05", aux.ClickDate)
		if err != nil {
			return err
		}
		r.ClickDate = &clickDate
	}

	if aux.StatusUpdated != "" {
		statusUpdated, err := time.Parse("2006-01-02 15:04:05", aux.StatusUpdated)
		if err != nil {
			return err
		}
		r.StatusUpdated = &statusUpdated
	}

	return nil
}

type Meta struct {
	Count  int `json:"count"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type Actions struct {
	Results []Result `json:"results"`
	Meta    `json:"_meta"`
}
type Position struct {
	TariffID            int       `json:"tariff_id"`
	Payment             float64   `json:"payment"`
	Rate                string    `json:"rate"`
	RateID              int       `json:"rate_id"`
	Datetime            time.Time `json:"datetime"`
	Amount              float64   `json:"amount"`
	Percentage          bool      `json:"percentage"`
	ProductID           string    `json:"product_id"`
	ProductName         string    `json:"product_name"`
	ProductImage        string    `json:"product_image"`
	ProductUrl          string    `json:"product_url"`
	ProductCategoryID   string    `json:"product_category_id"`
	ProductCategoryName string    `json:"product_category_name"`
	ID                  int       `json:"id"`
}

func (p *Position) UnmarshalJSON(data []byte) error {
	type Alias Position
	aux := &struct {
		Datetime string `json:"datetime"`
		Payment  string `json:"payment"`
		Amount   string `json:"amount"`
		*Alias
	}{
		Alias: (*Alias)(p),
	}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	datetime, err := time.Parse("2006-01-02 15:04:05", aux.Datetime)
	if err != nil {
		return err
	}
	p.Datetime = datetime

	payment, err := strconv.ParseFloat(aux.Payment, 64)
	if err != nil {
		return err
	}
	p.Payment = payment

	amount, err := strconv.ParseFloat(aux.Payment, 64)
	if err != nil {
		return err
	}
	p.Amount = amount

	return nil
}

type ActionsTotal struct {
	Currency   enums.Currency `json:"currency"`
	PaymentSum float64        `json:"payment_sum"`
	Cart       float64        `json:"cart"`
}
