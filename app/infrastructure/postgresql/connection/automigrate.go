package connection

import "github.com/rs/zerolog/log"

type migrate struct {
	connection *PostgreSQLConnection
}

func NewMigrate(connection *PostgreSQLConnection) *migrate {
	return &migrate{connection: connection}
}

func (m *migrate) AutoMigrateAll(tables ...interface{}) {
	db := m.connection.GetConnection()
	err := db.AutoMigrate(tables...)

	if err != nil {
		log.Err(db.Error).Msg("Error migrating tables")
	}
}
