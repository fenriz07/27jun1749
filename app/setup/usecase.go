package setup

import "github.com/fenriz07/27jun1749/app/usecase"

type usesCases struct {
	usecase.SaveLink
	usecase.FindLink
	usecase.DeleteLink
	usecase.RedirectLink
	usecase.CountLink
}

func NewUsesCases() usesCases {

	i := newInfrastructure()

	saveLinkUseCase := usecase.SaveLink{
		CreateLink: i.PostgresRepository.CreateLink,
		FindLink:   i.PostgresRepository.FindLink,
	}

	findLinkUseCase := usecase.FindLink{
		FindLink: i.PostgresRepository.FindLinkByCode,
	}

	deleteLinkUseCase := usecase.DeleteLink{
		DeleteLink: i.PostgresRepository.DeleteLink,
	}

	redirectLinkUseCase := usecase.RedirectLink{
		FindLink:         i.PostgresRepository.FindLinkByCode,
		IncrementCounter: i.PostgresRepository.IncrementCounter,
	}

	countLinkUseCase := usecase.CountLink{
		CountLink: i.PostgresRepository.CountLink,
	}

	return usesCases{
		SaveLink:     saveLinkUseCase,
		FindLink:     findLinkUseCase,
		DeleteLink:   deleteLinkUseCase,
		RedirectLink: redirectLinkUseCase,
		CountLink:    countLinkUseCase,
	}
}
