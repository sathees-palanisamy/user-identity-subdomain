package users

type User struct {
	Id          string `json:"id"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	TokenStatus string `json:"tokenStatus"`
	AccessToken string `json:"accessToken"`
	DateCreated string `json:"date_created"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Email       string `json:"email"`
	AccessToken string `json:"accessToken"`
}

type Users []User

type TokenDetail struct {
	AccessToken string `json:"accessToken"`
	TokenStatus string `json:"tokenStatus"`
}
