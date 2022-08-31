package models

type APIResponse struct {
	Status  bool              `json:"status"`
	Message string            `json:"message"`
	Error   map[string]string `json:"error,omitempty"`
	Data    interface{}       `json:"data,omitempty"`
}
