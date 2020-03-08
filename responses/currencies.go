package responses

import (
	"github.com/lueurxax/go-admitad-client/enums"
	"time"
)

type Currency struct {
	Date   time.Time      `json:"date"`
	Rate   string         `json:"rate"`
	Base   enums.Currency `json:"base"`
	Target enums.Currency `json:"target"`
}
