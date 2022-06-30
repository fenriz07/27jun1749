package rest

import (
	"errors"

	"github.com/fenriz07/27jun1749/app/domain/entity"
	"github.com/fenriz07/27jun1749/app/interface/rest/response"
	"github.com/fenriz07/27jun1749/app/usecase"
	"github.com/gogearbox/gearbox"
)

type DeleteShortHandler struct {
	usecase.DeleteLinkUseCase
}

func NewDeleteShortUrlHandler(gb gearbox.Gearbox, u usecase.DeleteLinkUseCase) {

	h := &DeleteShortHandler{
		DeleteLinkUseCase: u,
	}

	gb.Delete(BASE_PATH+":code", h.Delete)
}

func (h *DeleteShortHandler) Delete(ctx gearbox.Context) {

	code := ctx.Param("code")

	err := h.DeleteLinkUseCase(entity.Link{
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

	ctx.Status(gearbox.StatusOK)
	return

}
