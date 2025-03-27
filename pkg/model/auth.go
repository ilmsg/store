package model

type Auth struct {
	ID       string `json:"id" firestore:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthRegister struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type AuthLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token string `json:"token"`
}
