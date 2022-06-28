package repository

import (
	"errors"

	"github.com/fenriz07/27jun1749/app/infrastructure/postgresql/connection"
	"github.com/rs/zerolog/log"
)

const (
	PostgresRepositoryStructName = "PostgresRepository"
)

type PostgresRepository struct {
	conn *connection.PostgreSQLConnection
}

func NewPostgresRepository(conn *connection.PostgreSQLConnection) *PostgresRepository {
	return &PostgresRepository{conn: conn}
}

func (r *PostgresRepository) CheckDBHealth() error {
	pgConn := r.conn.GetConnection()

	db, err := pgConn.DB()
	if err != nil {
		log.Error().Err(err).Msg("error connecting to DB")
		return errors.New("error connecting to DB")
	}
	err = db.Ping()
	if err != nil {
		log.Error().Err(err).Msg("error in PING to DB")

		return errors.New("error in PING to DB")
	}
	return nil
}
