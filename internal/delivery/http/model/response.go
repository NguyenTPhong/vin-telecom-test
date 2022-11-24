package model

type ErrorResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type BillingData struct {
	CallCount  int64 `json:"call_count"`
	BlockCount int64 `json:"block_count"`
}
