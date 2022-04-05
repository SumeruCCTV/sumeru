package database

import (
	"context"
	"fmt"
	"github.com/SumeruCCTV/sumeru/pkg/config"
	"github.com/SumeruCCTV/sumeru/pkg/svcstat"
	"github.com/SumeruCCTV/sumeru/service/database/db"
	"github.com/SumeruCCTV/sumeru/service/database/models"
	"github.com/go-redis/redis/v8"
	"go.uber.org/atomic"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type Service struct {
	cfg *config.Config
	*db.Database

	closed atomic.Bool
}

func (Service) Name() string {
	return "database"
}

func (svc *Service) Start() error {
	gormConfig := new(gorm.Config)
	if svc.cfg.Database.DisableGormLogger {
		gormConfig.Logger = logger.Default.LogMode(logger.Silent)
	}
	gdb, err := gorm.Open(postgres.New(postgres.Config{DSN: svc.cfg.Database.PgDSN}), gormConfig)
	if err != nil || gdb == nil {
		panic(fmt.Errorf("failed to connect to postgres database: %w", err))
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     svc.cfg.Database.RedisDSN,
		Password: svc.cfg.Database.RedisPassword,
	})
	err = rdb.Ping(context.Background()).Err()
	if err != nil {
		panic(fmt.Errorf("failed to connect to redis database: %w", err))
	}

	svc.Database = &db.Database{
		DB:    gdb,
		Redis: rdb,
	}

	go func() {
		// Periodically check if Redis connection is closed.
		timer := time.NewTicker(10 * time.Second)
		for {
			select {
			case <-timer.C:
				svc.closed.Store(svc.Redis.Ping(context.Background()).Err() != nil)
			}
		}
	}()

	return svc.startMigration()
}

func (svc *Service) Stop() error {
	// gorm doesn't have a Close method
	return svc.Redis.Close()
}

func (svc *Service) Status() svcstat.Status {
	if svc.Database != nil && svc.DB != nil && svc.Redis != nil && !svc.closed.Load() {
		return svcstat.StatusHealthy
	}
	return svcstat.StatusUnhealthy
}

func (svc *Service) startMigration() error {
	return svc.AutoMigrate(
		&models.Account{},
		&models.Camera{},
	)
}
