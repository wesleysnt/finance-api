package commands

import (
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"github.com/wesleysnt/go-base/app/config"
)

var migrateCommand = &cobra.Command{
	Use:   "migrate",
	Short: "Run migrate database",
	Run: func(cmd *cobra.Command, args []string) {
		m, err := getMigrate()

		if err != nil {
			color.Yellowln(err)
			return
		}

		if m == nil {
			color.Yellowln("Please fill database config first")
			return
		}

		if err = m.Up(); err != nil && err != migrate.ErrNoChange {
			color.Redln("Migration failed:", err.Error())
			return
		}

		color.Greenln("Migration success")
	},
}

var migrateRefreshCmd = &cobra.Command{
	Use:   "migrate:refresh",
	Short: "Rollback all migrations and re-run them",
	Run: func(cmd *cobra.Command, args []string) {
		m, err := getMigrate()

		if err != nil {
			color.Redf("error drop: %v \n", err)
			return
		}

		err = m.Down()
		if err != nil {
			color.Redf("error drop: %v \n", err)
			return
		}

		err = m.Up()
		if err != nil && err != migrate.ErrNoChange {
			color.Redf("error up: %v \n", err)
			return
		}

		color.Greenln("Database refreshed")
	},
}

var migrateResetCmd = &cobra.Command{
	Use:   "migrate:reset",
	Short: "Rollback all migrations",
	Run: func(cmd *cobra.Command, args []string) {
		m, err := getMigrate()

		if err != nil {
			color.Redf("error down: %v \n", err)
			return
		}

		err = m.Down()
		if err != nil && err != migrate.ErrNoChange {
			color.Redf("error down: %v \n", err)
			return
		}

		color.Greenln("Database reset")
	},
}

var migrateRollbackCmd = &cobra.Command{
	Use:   "migrate:rollback [step]",
	Short: "Rollback the last migration",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		m, err := getMigrate()

		if err != nil {
			color.Yellowln("Please fill database config first")
			return
		}

		step := "-" + args[0]

		convStep, err := strconv.Atoi(step)

		if err != nil {
			color.Redf("error step: %v \n", err)
			return
		}

		err = m.Steps(convStep)
		if err != nil && err != migrate.ErrNoChange {
			color.Redf("error process step: %v \n", err)
			return
		}

		color.Greenln("Rolled back the last migration")
	},
}

var migrateStatusCmd = &cobra.Command{
	Use:   "migrate:status",
	Short: "Check status migration",
	Run: func(cmd *cobra.Command, args []string) {
		m, err := getMigrate()

		if err != nil {
			color.Yellowln("Please fill database config first")
			return
		}
		ver, dirty, err := m.Version()

		if err != nil {
			color.Redf("error step: %v \n", err)
			return
		}
		color.Greenln("Migration status: ")
		color.Greenf("version: %v \n", ver)
		color.Greenf("Dirty? (%v) \n", dirty)
	},
}

func init() {
	gobaseCommand.AddCommand(migrateCommand)
	gobaseCommand.AddCommand(migrateRefreshCmd)
	gobaseCommand.AddCommand(migrateResetCmd)
	gobaseCommand.AddCommand(migrateRollbackCmd)
	gobaseCommand.AddCommand(migrateStatusCmd)
}

func getMigrate() (*migrate.Migrate, error) {
	dir := "file://./database/migrations"
	gormDB, err := config.InitDb(config.GetEnv().Database)

	if err != nil {
		return nil, err
	}

	db, err := gormDB.DB()

	if err != nil {
		return nil, err
	}

	driver, errDriver := postgres.WithInstance(db, &postgres.Config{
		MigrationsTable: "migrations",
	})

	if errDriver != nil {
		return nil, errDriver
	}

	m, errMigrate := migrate.NewWithDatabaseInstance(
		dir,
		"postgres",
		driver,
	)

	if errMigrate != nil {
		return nil, errMigrate
	}

	return m, nil
}
