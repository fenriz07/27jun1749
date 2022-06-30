package setup

import (
	"time"

	"github.com/fenriz07/27jun1749/app/interface/rest"
	"github.com/go-playground/validator/v10"
	"github.com/gogearbox/gearbox"
)

func gearBox() gearbox.Gearbox {

	gb := gearbox.New(
		&gearbox.Settings{
			ReadTimeout:  time.Second * 10,
			WriteTimeout: time.Second * 10,
		},
	)

	return gb
}

func setEndpoints(gb gearbox.Gearbox, v *validator.Validate, listUseCase usesCases) {

	rest.NewRedirectHandler(gb, listUseCase.RedirectLink.Launch)

	rest.NewCreateShortUrlHandler(gb, v, listUseCase.SaveLink.Launch)
	rest.NewFindShortUrlHandler(gb, listUseCase.FindLink.Launch)
	rest.NewDeleteShortUrlHandler(gb, listUseCase.DeleteLink.Launch)
	rest.NewCountRedirectUrlHandler(gb, listUseCase.CountLink.Launch)

}
