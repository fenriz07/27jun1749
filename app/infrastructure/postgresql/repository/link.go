package repository

import (
	"errors"

	"github.com/fenriz07/27jun1749/app/domain/entity"
	"github.com/fenriz07/27jun1749/app/infrastructure/postgresql/models"
	"gorm.io/gorm"
)

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

func (r *PostgresRepository) DeleteLink(link entity.Link) error {

	conn := r.conn.GetConnection()

	tx := conn.Where("code = ?", link.Code).Delete(&models.Link{})

	if tx.RowsAffected == 0 {
		return entity.ErrEntityNotFound
	}

	if tx.Error != nil {
		return tx.Error
	}

	txCounter := conn.Where("code = ?", link.Code).Delete(&models.Counter{})

	if txCounter.Error != nil {
		return txCounter.Error
	}

	return nil
}

func (r *PostgresRepository) FindLink(link entity.Link) (*entity.Link, error) {

	conn := r.conn.GetConnection()

	linkModel := models.Link{}

	tx := conn.First(&linkModel, "id = ?", link.ID)

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, entity.ErrEntityNotFound
		}

		return nil, entity.ErrEntityNotFound
	}

	return &entity.Link{
		ID:   linkModel.ID,
		Code: linkModel.Code,
	}, nil
}

func (r *PostgresRepository) FindLinkByCode(link entity.Link) (*entity.Link, error) {

	conn := r.conn.GetConnection()

	linkModel := models.Link{}

	tx := conn.First(&linkModel, "code = ?", link.Code)

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, entity.ErrEntityNotFound
		}

		return nil, entity.ErrEntityNotFound
	}

	return &entity.Link{
		ID:   linkModel.ID,
		Code: linkModel.Code,
	}, nil
}
