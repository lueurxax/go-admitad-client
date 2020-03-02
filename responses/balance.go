package responses

// https://www.admitad.com/ru/developers/doc/api_ru/methods/private/user-balance/
type Balance []struct {
	Currency   string `json:"currency"`             // Код валюты
	Balance    string `json:"balance"`              // Сумма баланса
	Processing string `json:"processing,omitempty"` // В обработке
	Today      string `json:"today,omitempty"`      // Заработок за сегодня
	Stalled    string `json:"stalled,omitempty"`    // Сколько задержано
}
