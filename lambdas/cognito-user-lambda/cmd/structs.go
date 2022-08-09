package cmd

type UserLogin struct {
	Username string `json:"USERNAME"`
	Password string `json:"PASSWORD"`
}

type UserLoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type UserRegister struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password1 string `json:"password1"`
	Password2 string `json:"password2"`
}

type UserRegisterResponse struct {
	IsRegistered string `json:"isRegistered"`
}
