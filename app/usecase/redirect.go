package usecase

import (
	"github.com/fenriz07/27jun1749/app/domain/entity"
	"github.com/fenriz07/27jun1749/app/domain/infrastructure"
)

type RedirectLink struct {
	infrastructure.IncrementCounter
	infrastructure.FindLink
}

type RedirectLinkUseCase func(l entity.Link) (*entity.Link, error)

func (u *RedirectLink) Launch(l entity.Link) (*entity.Link, error) {

	link, err := u.FindLink(l)

	if err != nil {
		return nil, err
	}

	_ = u.IncrementCounter(*link)

	return link, nil

}
