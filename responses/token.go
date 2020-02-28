package responses

type Token struct {
	AccessToken  string `json:"access_token"`
	ExpiresIN    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}
