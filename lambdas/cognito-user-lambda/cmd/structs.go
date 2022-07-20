package cmd

type UserLogin struct {
	Username string `json:"USERNAME"`
	Password string `json:"PASSWORD"`
}

type UserLoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
