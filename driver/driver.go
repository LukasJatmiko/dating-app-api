package driver

import (
	"database/sql"
	"time"

	"github.com/LukasJatmiko/dating-app-api/constants"
	"github.com/LukasJatmiko/dating-app-api/types"
)

type Options struct {
	MaxOpenConnection     int
	MaxIdleConnection     int
	MaxConnectionLifetime time.Duration
	Database              types.Database
	URI                   string
}

type Driver interface {
	Init() error
	GetInstance() *sql.DB
	GetWrapperInstance() interface{}
	GetOptions() *Options
}

func NewDriver(opts *Options) Driver {
	switch opts.Database {

	//if mysql
	case constants.DriverMysql:
		return &MysqlConnectionPool{
			Options: opts,
		}

	//set default driver to postgres
	default:
		return &PostgresConnectionPool{
			Options: opts,
		}
	}
}
