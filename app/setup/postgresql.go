package setup

import (
	"os"

	"github.com/fenriz07/27jun1749/app/config"
	"github.com/fenriz07/27jun1749/app/infrastructure/postgresql/connection"
	"github.com/fenriz07/27jun1749/app/infrastructure/postgresql/models"
	"github.com/fenriz07/27jun1749/app/infrastructure/postgresql/repository"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func createPostgreDBConnection() *connection.PostgreSQLConnection {
	return connection.NewPostgreSQLConnection(connection.Config().
		Host(config.PostgreSQLHostName()).
		Port(config.PostgreSQLPort()).
		DatabaseName(config.PostgreSQLDBName()).
		User(config.PostgreSQLUsername()).
		Password(config.PostgreSQLPWD()),
	)
}

func setupPostgresSQL() *repository.PostgresRepository {

	conn := createPostgreDBConnection()

	initDBMigrations(conn.GetConnection())

	postgresqlRepository := repository.NewPostgresRepository(conn)

	return postgresqlRepository
}

func initDBMigrations(connection *gorm.DB) {
	migrator := connection.Migrator()

	migrations := []interface{}{
		&models.Link{},
	}

	err := migrator.AutoMigrate(migrations...)
	if err != nil {

		log.Error().
			Err(err).
			Msg("Error executing migrations")

		os.Exit(1)

	}

}
