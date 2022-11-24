package model

type CallRequest struct {
	Duration int64 `json:"duration" validate:"required,gt=0"`
}
