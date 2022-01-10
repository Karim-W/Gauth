package Models

import "gorm.io/gorm"

type Token struct {
	gorm.Model
	ClientID            string `json:"client_id"`
	UserID              string `json:"user_id"`
	RedirectURI         string `json:"redirect_uri"`
	Scope               string `json:"scope"`
	Code                string `json:"code"`
	CodeChallenge       string `json:"code_challenge"`
	CodeChallengeMethod string `json:"code_challenge_method"`
	CodeCreateAt        string `json:"code_create_at"`
	CodeExpiresIn       string `json:"code_expires_in"`
	Access              string `json:"access"`
	AccessCreateAt      string `json:"access_create_at"`
	AccessExpiresIn     string `json:"access_expires_in"`
	Refresh             string `json:"refresh"`
	RefreshCreateAt     string `json:"refresh_create_at"`
	RefreshExpiresIn    string `json:"refresh_expires_in"`
}
