package setup

import "github.com/fenriz07/27jun1749/app/infrastructure/postgresql/repository"

type listInfrastructure struct {
	PostgresRepository *repository.PostgresRepository
}

func newInfrastructure() listInfrastructure {

	postgresRepository := setupPostgresSQL()

	i := listInfrastructure{
		PostgresRepository: postgresRepository,
	}

	return i
}
