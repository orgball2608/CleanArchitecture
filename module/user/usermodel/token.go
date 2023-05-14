package usermodel

import "LearnGo/component/tokenprovider"

type Token struct {
	AccessToken  *tokenprovider.Token `json:"access_token"`
	RefreshToken *tokenprovider.Token `json:"refresh_token""`
}

func NewToken(at, rt *tokenprovider.Token) *Token {
	return &Token{
		AccessToken:  at,
		RefreshToken: rt,
	}
}

type RefreshTokenResponse struct {
	RefreshToken *tokenprovider.Token `json:"refresh_token""`
}

func NewRefreshTokenResponse(at *tokenprovider.Token) *RefreshTokenResponse {
	return &RefreshTokenResponse{
		RefreshToken: at,
	}
}
