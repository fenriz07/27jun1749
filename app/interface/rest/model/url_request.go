package model

type CreateShortRequestModel struct {
	URL string `json:"url" validate:"required,url"`
}
