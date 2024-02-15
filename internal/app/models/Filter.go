package models

type PageRequest struct {
	Limit int `json:"limit"`
	Next  int `json:"next"`
}
