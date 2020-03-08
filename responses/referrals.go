package responses

type Referrals struct {
	Results []ReferralResults `json:"results"`
	Meta    Meta              `json:"_meta"`
}
type ReferralResults struct {
	Username string      `json:"username"`
	ID       int         `json:"id"`
	Payment  interface{} `json:"payment"`
}
