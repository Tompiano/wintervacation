package model

type Conf struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectURL  string `json:"redirect_url"`
}
type TokenString struct {
	AccessToken string `json:"access_token"`
}
