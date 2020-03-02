package requests

import (
	"time"
)

// Status types:
// 1 (Новый)
// 2 (В работе)
// 3 (Ожидает)
// 4 (Закрыт)
type Ticket struct {
	Status    int
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
	Priority    int
}
