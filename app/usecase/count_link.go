package usecase

import (
	"github.com/fenriz07/27jun1749/app/domain/entity"
	"github.com/fenriz07/27jun1749/app/domain/infrastructure"
)

type CountLink struct {
	infrastructure.CountLink
}

type CountLinkUseCase func(l entity.Link) (*entity.Link, error)

func (u *CountLink) Launch(l entity.Link) (*entity.Link, error) {

	return u.CountLink(l)

}
