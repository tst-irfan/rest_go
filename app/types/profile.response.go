package types

type Profile struct {
	Id        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	DoB       string `json:"dob"`
	User 			UserResponse `json:"user"`
}
type ProfileResponse struct {
	Profile Profile `json:"profile"`
	Error   string  `json:"error"`	
}