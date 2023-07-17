package model

type GetIPInfoRequest struct {
	IPs []string `json:"ips"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
