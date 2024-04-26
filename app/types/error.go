package types

type Error struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
