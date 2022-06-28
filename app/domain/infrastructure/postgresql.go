package infrastructure

import "github.com/fenriz07/27jun1749/app/domain/entity"

type FindLink func(link entity.Link) (entity.Link, error)
type CreateLink func(link entity.Link) error
