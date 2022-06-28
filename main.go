package main

import (
	"os"

	"github.com/fenriz07/27jun1749/app/config"
	"github.com/fenriz07/27jun1749/app/setup"
	"github.com/rs/zerolog/log"
)

func main() {

	gb := setup.Setup()

	setup.NewUsesCases()

	listUseCase := setup.NewUsesCases()
	setup.SetInterface(gb, listUseCase)

	if err := gb.Start(":" + config.ServerPort()); err != nil {
		log.Error().
			Err(err).
			Msg("Error starting server")

		os.Exit(1)
	}

}
