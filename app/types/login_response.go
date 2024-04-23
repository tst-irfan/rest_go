package types

type LoginResponse struct {
	Error string `json:"error"`
	Token string `json:"token"`
}
