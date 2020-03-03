package requests

import (
	"github.com/lueurxax/go-admitad-client/enums"
)

type News struct {
	Offset   int
	Limit    int
	Language enums.Language
}
