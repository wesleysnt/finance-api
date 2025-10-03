package responses

type RegisterResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}
