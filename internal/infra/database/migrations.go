package database

import (
	"context"

	"github.com/gabefgonc/stockmanager/internal/domain/admins"
	"github.com/gabefgonc/stockmanager/internal/domain/products"
	"github.com/gabefgonc/stockmanager/internal/domain/stocks"
	"github.com/gabefgonc/stockmanager/internal/domain/stores"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type MigrationRunner struct {
	DB *gorm.DB
}

func (m *MigrationRunner) RunMigrations() error {
	err := m.DB.AutoMigrate(&stores.Store{})
	if err != nil {
		return err
	}
	err = m.DB.AutoMigrate(&admins.Admin{})
	if err != nil {
		return err
	}
	err = m.DB.AutoMigrate(&stocks.Stock{})
	if err != nil {
		return err
	}
	err = m.DB.AutoMigrate(&products.Product{})
	if err != nil {
		return err
	}

	return nil
}

func NewMigrationRunner(lc fx.Lifecycle, db *gorm.DB) *MigrationRunner {
	migrationRunner := &MigrationRunner{DB: db}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return migrationRunner.RunMigrations()
		},
	})
	return migrationRunner
}
