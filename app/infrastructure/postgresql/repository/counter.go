package repository

import (
	"errors"

	"github.com/fenriz07/27jun1749/app/domain/entity"
	"github.com/fenriz07/27jun1749/app/infrastructure/postgresql/models"
	"gorm.io/gorm"
)

func (r *PostgresRepository) IncrementCounter(link entity.Link) error {

	conn := r.conn.GetConnection()

	counterModel := models.Counter{
		Code: link.Code,
	}

	tx := conn.Create(&counterModel)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (r *PostgresRepository) CountLink(link entity.Link) (*entity.Link, error) {

	conn := r.conn.GetConnection()

	var count int64

	tx := conn.Model(&models.Counter{}).Where("code = ?", link.Code)

	if tx.Error != nil {

		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, entity.ErrEntityNotFound
		}

		return nil, tx.Error

	}

	tx.Count(&count)

	link.Count = count

	return &link, nil
}
