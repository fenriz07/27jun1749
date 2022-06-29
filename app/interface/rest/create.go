package rest

import (
	"github.com/fenriz07/27jun1749/app/interface/rest/mapper"
	"github.com/fenriz07/27jun1749/app/interface/rest/model"
	"github.com/fenriz07/27jun1749/app/interface/rest/response"
	"github.com/fenriz07/27jun1749/app/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/gogearbox/gearbox"
	"github.com/rs/zerolog/log"
)

type CreateShortHandler struct {
	*validator.Validate
	usecase.SaveLinkUseCase
}

func NewCreateShortUrlHandler(gb gearbox.Gearbox, validator *validator.Validate, u usecase.SaveLinkUseCase) {

	h := &CreateShortHandler{
		SaveLinkUseCase: u,
		Validate:        validator,
	}

	gb.Post(BASE_PATH, h.CreateShortUrl)

}

func (h *CreateShortHandler) CreateShortUrl(ctx gearbox.Context) {

	var bodyRequest model.CreateShortRequestModel

	if err := ctx.ParseBody(&bodyRequest); err != nil {
		log.Error().
			Err(err).
			Msg("error parsing request in create short url")
		ctx.Status(gearbox.StatusBadRequest)
		return
	}

	//Validate Body
	if err := h.Validate.Struct(bodyRequest); err != nil {

		log.Error().
			Err(err).
			Msg("bad request in create short url")

		if _, ok := err.(*validator.InvalidValidationError); ok {

			ctx.Status(gearbox.StatusBadRequest)
			return
		}

		ctx.Status(gearbox.StatusBadRequest).SendJSON(response.NewErrValidatingRequestResponse(err))
		return
	}

	link := mapper.RequestCreateToEntity(bodyRequest)

	linkResult, errUseCase := h.SaveLinkUseCase(link)

	if errUseCase != nil {
		ctx.Status(gearbox.StatusUnprocessableEntity).SendJSON(
			response.ErrValidatingRequestResponse{
				Message: errUseCase.Error(),
			},
		)

		return

	}

	ctx.Status(gearbox.StatusOK).SendJSON(response.LinkResponse{
		Url:  linkResult.ID,
		Code: linkResult.Code,
	})

	return
}
