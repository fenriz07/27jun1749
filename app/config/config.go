package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type config struct {

	//Gearbox env
	serverPort         string
	serverReadTimeOut  *time.Duration
	serverWriteTimeOut *time.Duration

	//PostgreSQL env
	postgreSQLMaxIdleConn int
	postgreSQLMaxOpenConn int
	postgreSQLAutoMigrate bool
	postgresSQLBatchSize  int

	//PostgreSQL scr
	postgreSQLHostName string
	postgreSQLPort     int
	postgreSQLDBName   string
	postgreSQLUsername string
	postgreSQLPWD      string
}

var conf config

func LoadConfiguration() {

	c := &config{}

	err := godotenv.Load()
	if err != nil {
		log.
			Error().
			Msg("Environment variables not found. Using environments established by deployment.yaml and envgonsul")
	}

	errors := make([]string, 0)

	errors = c.loadGearboxConfiguration(errors)
	errors = c.loadPostgreSQLConfiguration(errors)
	errors = c.loadPostgreSQLSCRT(errors)

	if len(errors) > 0 {
		log.Error().
			Interface("errors", errors).
			Msg("Required enviroments not found")
		os.Exit(1)
	}

	conf = *c
}

func (c *config) loadPostgreSQLSCRT(errors []string) []string {

	postgreSQLHostName := os.Getenv("database.postgres.hostName")
	if postgreSQLHostName == "" {
		errors = append(errors, "Error reading database.postgres.hostName from .env/envgonsul")
	}

	postgreSQLPort, err := strconv.Atoi(os.Getenv("database.postgres.port"))
	if err != nil {
		errors = append(errors, "Error reading/parsing database.postgres.port from .env/envgonsul")
	}

	postgreSQLDBName := os.Getenv("database.postgres.db.name")
	if postgreSQLDBName == "" {
		errors = append(errors, "Error reading database.postgres.db.name from .env/envgonsul")
	}

	postgreSQLUsername := os.Getenv("database.postgres.username")
	if postgreSQLUsername == "" {
		errors = append(errors, "Error reading database.postgres.username from .env/envgonsul")
	}

	postgreSQLPW := os.Getenv("database.postgres.pwd")
	if postgreSQLPW == "" {
		errors = append(errors, "Error reading database.postgres.pwd from .env/envgonsul")
	}

	//postgresql scr
	c.postgreSQLHostName = postgreSQLHostName
	c.postgreSQLPort = postgreSQLPort
	c.postgreSQLDBName = postgreSQLDBName
	c.postgreSQLUsername = postgreSQLUsername
	c.postgreSQLPWD = postgreSQLPW

	return errors
}

func (c *config) loadPostgreSQLConfiguration(errors []string) []string {

	postgreSQLAutoMigrate, err := strconv.ParseBool(os.Getenv("database.postgres.automigrate"))
	if err != nil {
		errors = append(errors, "Error reading/parsing database.postgres.automigrate from .env/deployment.yaml")
	}

	postgreSQLMaxIdleConn, err := strconv.Atoi(os.Getenv("database.postgres.maxIdleConn"))
	if err != nil {
		errors = append(errors, "Error reading database.postgres.maxIdleConn from .env/deployment.yaml")
	}

	postgreSQLMaxOpenConn, err := strconv.Atoi(os.Getenv("database.postgres.maxOpenConn"))
	if err != nil {
		errors = append(errors, "Error reading database.postgres.maxOpenConn from .env/deployment.yaml")
	}

	postgreSQLBatchSize, err := strconv.Atoi(os.Getenv("database.postgres.batchSize"))
	if err != nil {
		errors = append(errors, "Error reading database.postgres.batchSize from .env/deployment.yaml")
	}

	//postgresql config
	c.postgreSQLMaxIdleConn = postgreSQLMaxIdleConn
	c.postgreSQLMaxOpenConn = postgreSQLMaxOpenConn
	c.postgreSQLAutoMigrate = postgreSQLAutoMigrate
	c.postgresSQLBatchSize = postgreSQLBatchSize

	return errors
}

func (c *config) loadGearboxConfiguration(errors []string) []string {

	serverPort := os.Getenv("server.port")
	if serverPort == "" {
		errors = append(errors, "Error reading server.port from .env/deployment.yaml")
	}
	serverReadTimeOut, err := time.ParseDuration(os.Getenv("server.readTimeout"))
	if err != nil {
		errors = append(errors, "Error reading/parsing server.readTimeout from .env/deployment.yaml")
	}
	serverWriteTimeOut, err := time.ParseDuration(os.Getenv("server.writeTimeout"))
	if err != nil {
		errors = append(errors, "Error reading/parsing server.writeTimeout from .env/deployment.yaml")
	}

	//Gearbox config
	c.serverPort = serverPort
	c.serverReadTimeOut = &serverReadTimeOut
	c.serverWriteTimeOut = &serverWriteTimeOut

	return errors
}

func ServerPort() string {
	return conf.serverPort
}
func ServerReadTimeOut() time.Duration {
	return *conf.serverReadTimeOut
}
func ServerWriteTimeOut() time.Duration {
	return *conf.serverWriteTimeOut
}

func PostgreSQLMaxIdleConn() int {
	return conf.postgreSQLMaxIdleConn
}

func PostgreSQLMaxOpenConn() int {
	return conf.postgreSQLMaxOpenConn
}

func PostgreSQLAutoMigrate() bool {
	return conf.postgreSQLAutoMigrate
}

func PostgreSQLHostName() string {
	return conf.postgreSQLHostName
}
func PostgreSQLPort() int {
	return conf.postgreSQLPort
}
func PostgreSQLDBName() string {
	return conf.postgreSQLDBName
}
func PostgreSQLUsername() string {
	return conf.postgreSQLUsername
}
func PostgreSQLPWD() string {
	return conf.postgreSQLPWD
}

func PostgresSQLBatchSize() int {
	return conf.postgresSQLBatchSize
}
