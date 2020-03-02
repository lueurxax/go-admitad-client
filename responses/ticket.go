package responses

import (
	"github.com/lueurxax/go-admitad-client/enums"
)

type Ticket struct {
	AdvCampaign  interface{}    `json:"advcampaign"`
	Status       int            `json:"status"` // Статус (1 - Новый, 2 - В работе, 3 - Ожидает, 4 - Закрыт)
	Category     Category       `json:"category"`
	DateModified string         `json:"date_modified"` // Дата и время последнего обновления тикета
	Text         string         `json:"text"`          // HTML текст
	PlainText    string         `json:"plain_text"`    // Текст без html
	Messages     []Messages     `json:"messages"`
	Priority     enums.Priority `json:"priority"`     // Приоритет (1 - Низкий, 2 - Нормальный, 3 - Срочный)
	DateCreated  string         `json:"date_created"` // Дата и время создания тикета
	ID           int            `json:"id"`           // Идентификатор тикета
	Subject      string         `json:"subject"`      // Тема тикета
	Comments     []interface{} `json:"comments,omitempty"`
}
type Category struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}
type CreatedBy struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
type Messages struct {
	Files       []interface{} `json:"files"`
	Text        string        `json:"text"`
	PlainText   string        `json:"plain_text"`
	ID          int           `json:"id"`
	CreatedBy   CreatedBy     `json:"created_by"`
	DateCreated string        `json:"date_created"`
}
