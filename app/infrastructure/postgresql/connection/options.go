package connection

import (
	"fmt"
	"os"

	"github.com/fenriz07/27jun1749/app/config"
)

type Options struct {
	databaseName *string
	host         *string
	port         *int
	user         *string
	password     *string
}

func Config() *Options {
	return new(Options)
}

func (c *Options) DatabaseName(databaseName string) *Options {
	c.databaseName = &databaseName
	return c
}

func (c *Options) Host(host string) *Options {
	c.host = &host

	return c
}

func (c *Options) Port(port int) *Options {
	c.port = &port

	return c
}

func (c *Options) User(user string) *Options {
	c.user = &user

	return c
}

func (c *Options) Password(password string) *Options {
	c.password = &password

	return c
}

func MergeOptions(opts ...*Options) *Options {
	option := new(Options)

	for _, opt := range opts {
		if opt.databaseName != nil {
			option.databaseName = opt.databaseName
		}

		if opt.host != nil {
			option.host = opt.host
		}

		if opt.port != nil {
			option.port = opt.port
		}

		if opt.user != nil {
			option.user = opt.user
		}

		if opt.password != nil {
			option.password = opt.password
		}
	}
	return option
}

func (c *Options) GetURLConnection() string {
	urlCockroachFormat := "host=%v user=%v password=%v dbname=%v port=%v"

	environment := os.Getenv("ENV")
	if environment == "local" || environment == "" {
		urlCockroachFormat = "host=%v user=%v password=%v dbname=%v port=%v sslmode=disable"
	}

	if c.port == nil {
		p := config.PostgreSQLPort()
		c.port = &p
	}

	urlFormat := fmt.Sprintf(urlCockroachFormat, *c.host, *c.user, *c.password, *c.databaseName, *c.port)

	return urlFormat
}
