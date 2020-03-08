package responses

type Me struct {
	Username        string `json:"username"`         // 	Идентификатор веб-мастера
	FirstName       string `json:"first_name"`       // Логин веб-мастера
	LastName        string `json:"last_name"`        // Имя веб-мастера
	ID              int    `json:"id"`               // Фамилия веб-мастера
	Language        string `json:"language"`         // Язык веб-мастера
	DefaultCurrency string `json:"default_currency"` // Валюта веб-мастера
	Phone           string `json:"phone,omitempty"`  // Телефон веб-мастера
	Email           string `json:"email,omitempty"`  // Адрес электронной почты веб-мастера
}
