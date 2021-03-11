package model

type RequestPayload struct {
	Search string `json:"search"`
	Page   string `json:"page"`
}
