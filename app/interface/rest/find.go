package rest

import (
	"github.com/fenriz07/27jun1749/app/usecase"
	"github.com/gogearbox/gearbox"
)

type FindShortHandler struct {
}

func NewFindShortUrlHandler(gb gearbox.Gearbox, u usecase.SaveLinkUseCase) {

	h := &FindShortHandler{}

	gb.Get(BASE_PATH+":code", h.Find)
}

func (h *FindShortHandler) Find(ctx gearbox.Context) {

	//code := ctx.Param("code")

}
