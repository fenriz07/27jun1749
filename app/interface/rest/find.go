package rest

import (
	"errors"

	"github.com/fenriz07/27jun1749/app/domain/entity"
	"github.com/fenriz07/27jun1749/app/interface/rest/response"
	"github.com/fenriz07/27jun1749/app/usecase"
	"github.com/gogearbox/gearbox"
)

type FindShortHandler struct {
	usecase.FindLinkUseCase
}

func NewFindShortUrlHandler(gb gearbox.Gearbox, u usecase.FindLinkUseCase) {

	h := &FindShortHandler{
		FindLinkUseCase: u,
	}

	gb.Get(BASE_PATH+":code", h.Find)
}

func (h *FindShortHandler) Find(ctx gearbox.Context) {

	code := ctx.Param("code")

	linkResult, err := h.FindLinkUseCase(entity.Link{
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

	ctx.Status(gearbox.StatusOK).SendJSON(response.LinkResponse{
		Url:  linkResult.ID,
		Code: linkResult.Code,
	})
	return

}
