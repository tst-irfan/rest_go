package types

type UserResponse struct {
	Id        uint   `json:"id"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type RegisterResponse struct {
	User  UserResponse `json:"user"`
	Error string       `json:"error"`
}
