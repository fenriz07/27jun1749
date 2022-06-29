package usecase

import (
	"errors"

	"github.com/fenriz07/27jun1749/app/domain/entity"
	"github.com/fenriz07/27jun1749/app/domain/infrastructure"
	"github.com/google/uuid"
)

type SaveLink struct {
	infrastructure.CreateLink
	infrastructure.FindLink
}

type SaveLinkUseCase func(link entity.Link) (*entity.Link, error)

func (u *SaveLink) Launch(l entity.Link) (*entity.Link, error) {

	link, errFindingLink := u.FindLink(l)

	if errFindingLink != nil {
		if errors.Is(entity.ErrEntityNotFound, errFindingLink) {

			l.Code = u.GetCode()

			errCreatingLink := u.CreateLink(l)

			if errCreatingLink != nil {
				return nil, entity.ErrCreatingEntity
			}

			return &l, nil

		} else {
			return nil, entity.ErrCreatingEntity
		}
	}

	return link, nil

}

func (u *SaveLink) GetCode() string {
	uuid := uuid.New()

	return uuid.String()[:8]

}
