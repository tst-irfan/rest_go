package types

type MetaData struct {
	TotalPages int `json:"total_pages"`
	TotalItems int `json:"total_items"`
	Page       int `json:"page"`
	Size       int `json:"size"`
}
