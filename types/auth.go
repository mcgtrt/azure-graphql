package types

type AuthParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Email   string `json:"email"`
	Expires int64  `json:"expires"`
	Token   string `json:"token"`
}

type AuthEmployee struct {
	ID                string
	Email             string
	EncryptedPassword string
}
