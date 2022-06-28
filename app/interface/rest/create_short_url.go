package rest

import (
	"github.com/fenriz07/27jun1749/app/domain/entity"
	"github.com/fenriz07/27jun1749/app/usecase"
	"github.com/gogearbox/gearbox"
)

type CreateShortHandler struct {
	usecase.SaveLinkUseCase
}

func NewCreateShortUrlHandler(gb gearbox.Gearbox, u usecase.SaveLinkUseCase) {

	h := &CreateShortHandler{
		SaveLinkUseCase: u,
	}

	gb.Post(BASE_PATH, h.CreateShortUrl)

}

func (h *CreateShortHandler) CreateShortUrl(ctx gearbox.Context) {

	h.SaveLinkUseCase(entity.Link{
		ID:   "id",
		Code: "epa",
	})

	ctx.SendString("hello")
}
