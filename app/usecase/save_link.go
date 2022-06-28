package usecase

import (
	"github.com/fenriz07/27jun1749/app/domain/entity"
	"github.com/fenriz07/27jun1749/app/domain/infrastructure"
)

type SaveLink struct {
	infrastructure.CreateLink
	infrastructure.FindLink
}

type SaveLinkUseCase func(link entity.Link) (*entity.Link, error)

func (u *SaveLink) Launch(link entity.Link) (*entity.Link, error) {

	errCreatingLink := u.CreateLink(link)

	if errCreatingLink != nil {
		return nil, errCreatingLink
	}

	return nil, nil
}
