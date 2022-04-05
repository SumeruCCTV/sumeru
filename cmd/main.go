package main

import (
	"fmt"
	"github.com/SumeruCCTV/sumeru"
	"github.com/SumeruCCTV/sumeru/pkg"
	"github.com/SumeruCCTV/sumeru/pkg/config"
	"github.com/SumeruCCTV/sumeru/pkg/utils"
	"github.com/SumeruCCTV/sumeru/service/camera"
	"github.com/SumeruCCTV/sumeru/service/database"
	"github.com/SumeruCCTV/sumeru/service/web"
	"github.com/SumeruCCTV/sumeru/service/web/routes"
	"go.uber.org/zap"
	"math/rand"
	"os"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func mustgetenv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Sprintf("%s is not set", key))
	}
	return value
}

func main() {
	host := mustgetenv("DB_HOST")
	passwCRDB := mustgetenv("PASS_CRDB")
	passRedis := mustgetenv("PASS_REDIS")
	tempConfig := &config.Config{
		Database: &config.Database{
			DisableGormLogger: true,
			PgDSN: fmt.Sprintf(
				"host=%s port=%d user=%s password=%s dbname=%s",
				host, 26257, "sumeru", passwCRDB, "sumeru",
			),
			RedisDSN:      fmt.Sprintf("%s:6381", host),
			RedisPassword: passRedis,
		},
		Web: &config.Web{
			Port: 3000,
		},
	}

	logger := createLogger()
	app := pkg.New(tempConfig, logger)
	sumeru.App = app

	app.AddService(&database.Service{})

	app.AddService(&web.Service{})
	routes.Init()

	app.AddService(&camera.Service{})

	if err := app.Run(); err != nil {
		logger.Error(err)
	}
}

func createLogger() *utils.Logger {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	return logger.Sugar()
}
