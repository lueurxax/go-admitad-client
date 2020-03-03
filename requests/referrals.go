package requests

import (
	"time"
)

type Referrals struct {
	Limit     int
	Offset    int
	DateStart *time.Time // %d.%m.%Y
	DateEnd   *time.Time // %d.%m.%Y
}
