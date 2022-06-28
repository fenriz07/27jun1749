package setup

import (
	"time"

	"github.com/fenriz07/27jun1749/app/interface/rest"
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

func setEndpoints(gb gearbox.Gearbox, listUseCase usesCases) {

	rest.NewCreateShortUrlHandler(gb, listUseCase.Launch)
}
