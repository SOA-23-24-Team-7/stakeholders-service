package dto

type TokenResponse struct {
	Id          int64  `json:"id"`
	AccessToken string `json:"accessToken"`
}