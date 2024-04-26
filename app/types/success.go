package types

type Success[T any] struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}
