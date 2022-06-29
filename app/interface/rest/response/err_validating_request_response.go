package response

import (
	"github.com/go-playground/validator/v10"
)

type ErrValidatingRequestResponse struct {
	Message string `json:"message"`
}

func NewErrValidatingRequestResponse(err error) map[string]ErrValidatingRequestResponse {
	list := make(map[string]ErrValidatingRequestResponse)

	for _, err := range err.(validator.ValidationErrors) {

		var message string

		message = err.Tag()

		list[err.Field()] = ErrValidatingRequestResponse{
			Message: message,
		}

	}

	return list
}
