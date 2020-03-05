package requests

import (
    "github.com/lueurxax/go-admitad-client/enums"
    "time"
)

// https://www.admitad.com/en/developers/doc/api_ru/methods/statistics/statistics-sub-id/
type Statistics struct {
    Offset int
    Limit int
    DateStart time.Time
    DateEnd time.Time
    Website int
    Campaign int
    SubID int
    OrderBy enums.Sorting
    Total bool
    GroupSubIDs int // ???
}
