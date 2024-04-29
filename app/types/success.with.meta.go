package types

type SuccessWithMeta[T any] struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	Data    T      `json:"data"`
	MetaData MetaData `json:"meta"`
}