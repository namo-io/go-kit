package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/namo-io/go-kit/log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Configuration postgresql configuration
type Configuration struct {
	URL          string
	Port         int
	Database     string
	UserID       string
	UserPassword string

	MaxConnectionCount  int
	IdleConnectionCount int
}

// Postgres postrgres instance
type Postgres struct {
	*gorm.DB
}

// New create postgresql instance
func New(ctx context.Context, cfg *Configuration) (*Postgres, error) {
	logger := log.New().WithContext(ctx)

	logger.Debug("creating postgresql instance ...")
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=disable TimeZone=Asia/Seoul",
			cfg.URL,
			cfg.UserID,
			cfg.UserPassword,
			cfg.Database,
			cfg.Port,
		),
	}), &gorm.Config{})

	if err != nil {
		logger.Error(err)
		return nil, err
	}

	// connection pool setup
	sqlDB, err := db.DB()
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	sqlDB.SetMaxIdleConns(cfg.IdleConnectionCount)
	sqlDB.SetMaxOpenConns(cfg.MaxConnectionCount)
	sqlDB.SetConnMaxLifetime(time.Hour)

	postgres := &Postgres{
		DB: db.Debug(),
	}

	logger.Debug("create postgresql instance")
	return postgres, nil
}

// Health ping test
func (p *Postgres) Health() error {
	sqlDB, err := p.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Ping()
}

// Close connection close
func (p *Postgres) Close(ctx context.Context) {
	logger := log.New().WithContext(ctx)
	logger.Debug("postgresql connection close")

	p.Close(ctx)
}
