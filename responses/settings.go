package responses

// https://www.admitad.com/ru/developers/doc/api_ru/methods/private/payment-settings/
type PaymentDataInformation []struct {
	City             string   `json:"city"`
	Name             string   `json:"name"`
	Country          string   `json:"country"`
	Currency         []string `json:"currency"`
	ConversionRate   string   `json:"conversion_rate"`
	Address          string   `json:"address"`
	WebmoneyAccount  string   `json:"webmoney_account,omitempty"`
	WithdrawalType   string   `json:"withdrawal_type"`
	ID               int      `json:"id"`
	ZipCode          string   `json:"zip_code"`
	PaypalAccount    string   `json:"paypal_account,omitempty"`
	BankAccountOwner string   `json:"bank_account_owner,omitempty"`
	BankAccount      string   `json:"bank_account,omitempty"`
	BankCode         string   `json:"bank_code,omitempty"`
	BankName         string   `json:"bank_name,omitempty"`
	BankBIC          string   `json:"bank_bic,omitempty"` // BIC/SWIFT
	BankIBAN         string   `json:"bank_iban,omitempty"`
}
