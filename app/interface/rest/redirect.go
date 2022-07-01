package rest

import (
	"github.com/fenriz07/27jun1749/app/domain/entity"
	"github.com/fenriz07/27jun1749/app/interface/rest/response"
	"github.com/fenriz07/27jun1749/app/usecase"
	"github.com/gogearbox/gearbox"
)

type RedirectHandler struct {
	usecase.RedirectLinkUseCase
}

func NewRedirectHandler(gb gearbox.Gearbox, u usecase.RedirectLinkUseCase) {
	h := RedirectHandler{
		RedirectLinkUseCase: u,
	}

	gb.NotFound(h.Redirect)
}

func (h *RedirectHandler) Redirect(ctx gearbox.Context) {

	lastPathSegment := ctx.Context().URI().LastPathSegment()

	link := entity.Link{
		Code: string(lastPathSegment),
	}

	l, err := h.RedirectLinkUseCase(link)

	if err != nil {
		ctx.Status(gearbox.StatusNotFound).SendJSON(response.ErrValidatingRequestResponse{
			Message: "not found",
		})
		return
	}

	ctx.Context().Redirect(l.ID, 302)
}
