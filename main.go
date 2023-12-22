package main

import (
	"fmt"
	"time"

	"github.com/LukasJatmiko/dating-app-api/constants"
	"github.com/LukasJatmiko/dating-app-api/driver"
	"github.com/LukasJatmiko/dating-app-api/types"
	"github.com/LukasJatmiko/dating-app-api/utils"
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	appName := utils.GetEnvOrDefaultString(string(constants.ENVAppName), "dating-app-api")

	//init db connection
	dbdriver := driver.NewDriver(&driver.Options{
		Database:              types.Database(utils.GetEnvOrDefaultString(string(constants.ENVDBDriverType), "POSTGRES")),
		URI:                   utils.GetEnvOrDefaultString(string(constants.ENVDBURI), "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disabled"),
		MaxOpenConnection:     utils.GetEnvOrDefaultInt(string(constants.ENVMaxOpenConn), 5),
		MaxIdleConnection:     utils.GetEnvOrDefaultInt(string(constants.ENVMaxIdleConn), 1),
		MaxConnectionLifetime: utils.GetEnvOrDefaultDuration(string(constants.ENVMaxConnLifetime), (1800 * time.Second)),
	})
	if e := dbdriver.Init(); e != nil {
		panic(e)
	}

	app := fiber.New()
	app.Name(appName)

	app.Use(cors.New(cors.Config{
		//to be improve
		//configurable origins
		//parse allowed origins from environments
		AllowOrigins: "*",
	}))

	appPort := utils.GetEnvOrDefaultString(string(constants.ENVAppPort), "8080")
	app.Listen(fmt.Sprintf(":%v", appPort))
}
