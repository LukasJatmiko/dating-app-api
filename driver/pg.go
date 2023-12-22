package driver

import (
	"database/sql"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresConnectionPool struct {
	Options      *Options
	Instance     *sql.DB
	GormInstance *gorm.DB
}

func (cp *PostgresConnectionPool) Init() error {
	if gormdb, e := gorm.Open(postgres.Open(cp.Options.URI), &gorm.Config{}); e != nil {
		return e
	} else {
		cp.GormInstance = gormdb
		if sqlDB, e := gormdb.DB(); e != nil {
			return e
		} else {
			cp.Instance = sqlDB
			sqlDB.SetMaxOpenConns(cp.Options.MaxOpenConnection)
			sqlDB.SetMaxIdleConns(cp.Options.MaxIdleConnection)
			sqlDB.SetConnMaxLifetime(cp.Options.MaxConnectionLifetime)

			return nil
		}
	}
}

func (cp *PostgresConnectionPool) GetInstance() *sql.DB {
	return cp.Instance
}

func (cp *PostgresConnectionPool) GetWrapperInstance() interface{} {
	return cp.GormInstance
}

func (cp *PostgresConnectionPool) GetOptions() *Options {
	return cp.Options
}
