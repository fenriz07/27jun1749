package repository

import (
	"github.com/fenriz07/27jun1749/app/domain/entity"
	"github.com/fenriz07/27jun1749/app/infrastructure/postgresql/models"
)

//type FindLink func(link entity.Link) (entity.Link, error)

func (r *PostgresRepository) CreateLink(link entity.Link) error {

	conn := r.conn.GetConnection()

	linkModel := models.Link{
		ID:   link.ID,
		Code: link.Code,
	}

	tx := conn.Create(&linkModel)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
