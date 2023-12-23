package driver

import (
	"database/sql"
	"fmt"

	"gorm.io/gorm"
)

type MysqlConnectionPool struct {
	Options      *Options
	Instance     *sql.DB
	GormInstance *gorm.DB
}

func (cp *MysqlConnectionPool) Init() error {
	//to do
	return fmt.Errorf("MYSQL driver is not yet supported")
}

func (cp *MysqlConnectionPool) GetInstance() *sql.DB {
	return cp.Instance
}

func (cp *MysqlConnectionPool) GetWrapperInstance() interface{} {
	return cp.GormInstance
}

func (cp *MysqlConnectionPool) GetOptions() *Options {
	return cp.Options
}
