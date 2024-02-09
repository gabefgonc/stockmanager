package database

import "go.uber.org/fx"

var Module = fx.Provide(
	NewMigrationRunner,
	NewDatabaseConnection,
)
