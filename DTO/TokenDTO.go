package DTO

type TokenDTO struct {
	AccessToken  string `json:"access_token"`  // Short lived access token, used to authenticate the user when making requests to the API
	RefreshToken string `json:"refresh_token"` // Long lived refresh token, used to get a new access token when the access token expires
}