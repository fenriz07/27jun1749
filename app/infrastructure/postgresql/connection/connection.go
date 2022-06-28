package connection

import (
	"database/sql"
	"errors"
	"os"
	"sync"
	"time"

	"github.com/fenriz07/27jun1749/app/config"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLog "gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

type PostgreSQLRepository interface {
	GetConnection() *gorm.DB
}

type PostgreSQLConnection struct {
	opts       *Options
	url        string
	connection *gorm.DB
	once       sync.Once
}

func (r *PostgreSQLConnection) SetConnection(gormConnection *gorm.DB) {
	r.connection = gormConnection
}

func NewPostgreSQLConnection(opts ...*Options) *PostgreSQLConnection {
	databaseOptions := MergeOptions(opts...)
	url := databaseOptions.GetURLConnection()

	if url == "" {

		log.Err(errors.New("url is empty"))
		os.Exit(1)
	}

	return &PostgreSQLConnection{
		opts: databaseOptions,
		url:  url,
	}
}

var lock = &sync.Mutex{}

func (r *PostgreSQLConnection) GetConnection() *gorm.DB {
	if r.connection == nil {
		sqlDb, err := sql.Open("pgx", r.url)

		if err != nil {
			log.Err(err)
			os.Exit(1)
		}

		connection, err := gorm.Open(postgres.New(postgres.Config{
			Conn:                 sqlDb,
			PreferSimpleProtocol: true,
		}), &gorm.Config{
			Logger:               gormLog.Default.LogMode(gormLog.Silent),
			FullSaveAssociations: true,
			CreateBatchSize:      config.PostgresSQLBatchSize(),
		})

		if err != nil {
			log.Err(err).Msg("error trying to connect to DB")
		} else {
			_, err := connection.DB()
			if err != nil {
				log.Err(err).Msg("error trying to connect to DB")
			}
			_ = connection.Use(dbresolver.Register(dbresolver.Config{}).
				SetMaxIdleConns(config.PostgreSQLMaxIdleConn()).
				SetMaxOpenConns(config.PostgreSQLMaxOpenConn()).
				SetConnMaxLifetime(time.Hour).
				SetConnMaxIdleTime(time.Hour))

			r.connection = connection
		}
	}
	return r.connection

}
