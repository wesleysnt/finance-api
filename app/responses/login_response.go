package responses

type LoginResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}
