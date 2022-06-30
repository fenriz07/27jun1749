package rest

import (
	"errors"

	"github.com/fenriz07/27jun1749/app/domain/entity"
	"github.com/fenriz07/27jun1749/app/interface/rest/response"
	"github.com/fenriz07/27jun1749/app/usecase"
	"github.com/gogearbox/gearbox"
)

type CountRedirectShortHandler struct {
	usecase.CountLinkUseCase
}

func NewCountRedirectUrlHandler(gb gearbox.Gearbox, u usecase.CountLinkUseCase) {

	h := &CountRedirectShortHandler{
		CountLinkUseCase: u,
	}

	gb.Get(COUNT_PATH+":code", h.Count)
}

func (h *CountRedirectShortHandler) Count(ctx gearbox.Context) {

	code := ctx.Param("code")

	linkResult, err := h.CountLinkUseCase(entity.Link{
		Code: code,
	})

	if err != nil {
		if errors.Is(err, entity.ErrEntityNotFound) {

			ctx.Status(gearbox.StatusNotFound).SendJSON(
				response.ErrValidatingRequestResponse{
					Message: err.Error(),
				},
			)

			return
		}

		ctx.Status(gearbox.StatusUnprocessableEntity)
		return
	}

	ctx.Status(gearbox.StatusOK).SendJSON(response.CountResponse{
		Count: linkResult.Count,
		Code:  linkResult.Code,
	})
	return

}
