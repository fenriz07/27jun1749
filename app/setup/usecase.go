package setup

import "github.com/fenriz07/27jun1749/app/usecase"

type usesCases struct {
	usecase.SaveLink
}

func NewUsesCases() usesCases {

	i := newInfrastructure()

	saveLinkUseCase := usecase.SaveLink{
		CreateLink: i.PostgresRepository.CreateLink,
		FindLink:   i.PostgresRepository.FindLink,
	}

	return usesCases{
		SaveLink: saveLinkUseCase,
	}
}