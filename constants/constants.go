package constants

import (
	"github.com/LukasJatmiko/dating-app-api/types"
)

/*
driver
*/
const DriverMysql types.Database = "MYSQL"
const DriverPostgres types.Database = "POSTGRES"

/*
environments
*/
const ENVAppName types.Environment = "APP_NAME"
const ENVAppPort types.Environment = "APP_PORT"

const ENVDBURI types.Environment = "DB_URI"
const ENVDBDriverType types.Environment = "DB_DRIVER_TYPE"
const ENVMaxOpenConn types.Environment = "DB_MAX_OPEN_CONN"
const ENVMaxConnLifetime types.Environment = "DB_MAX_CONN_LIFETIME"
const ENVMaxIdleConn types.Environment = "DB_MAX_IDLE_CONN"
const ENVAuthJWTPrivateKey types.Environment = "AUTH_JWT_PRIVATE_KEY"
const ENVAuthJWTPublicKey types.Environment = "AUTH_JWT_PUBLIC_KEY"
