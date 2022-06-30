package usecase

import (
	"github.com/fenriz07/27jun1749/app/domain/entity"
	"github.com/fenriz07/27jun1749/app/domain/infrastructure"
)

type FindLink struct {
	infrastructure.FindLink
}

type FindLinkUseCase func(l entity.Link) (*entity.Link, error)

func (u *FindLink) Launch(l entity.Link) (*entity.Link, error) {

	link, errFindingLink := u.FindLink(l)

	if errFindingLink != nil {
		return nil, errFindingLink
	}

	return link, nil
}
