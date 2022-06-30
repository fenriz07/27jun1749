package usecase

import (
	"github.com/fenriz07/27jun1749/app/domain/entity"
	"github.com/fenriz07/27jun1749/app/domain/infrastructure"
)

type DeleteLink struct {
	infrastructure.DeleteLink
}

type DeleteLinkUseCase func(l entity.Link) error

func (u *DeleteLink) Launch(l entity.Link) error {

	return u.DeleteLink(l)

}
