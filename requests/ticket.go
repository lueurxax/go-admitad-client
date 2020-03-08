package requests

import (
	"github.com/lueurxax/go-admitad-client/enums"
	"time"
)

// Status types:
// 1 (Новый)
// 2 (В работе)
// 3 (Ожидает)
// 4 (Закрыт)
type Ticket struct {
	Status    enums.StatusID
	DateStart *time.Time // %d.%m.%Y
	DateEnd   *time.Time // %d.%m.%Y
	Limit     int
	Offset    int
}

// Priority types
// 1 (Низкий)
// 2 (Нормальный)
// 3 (Срочный)
type CreateTicket struct {
	Subject     string
	Text        string
	AdvCampaign int
	Category    int
	Priority    enums.Priority
}

type TicketComment struct {
	ID      int
	Comment string
}
