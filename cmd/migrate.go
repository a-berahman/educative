package cmd

import (
	"os"
	"path/filepath"
	"strings"

	"log"

	libmigrate "github.com/golang-migrate/migrate"
	// ...
	_ "github.com/golang-migrate/migrate/source/file"
	// ...
	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/database"
	"github.com/golang-migrate/migrate/database/postgres"
	"github.com/spf13/cobra"

	"github.com/a-berahman/educative/config"
)

var (
	steps           int
	migrationsPath  string
	migrationsTable string
)

var migrateDatabaseCMD = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migrations",
	Run: func(cmd *cobra.Command, args []string) {
		migrateDB()
	},
}

func init() {
	migrateDatabaseCMD.Flags().StringVarP(&migrationsPath, "migrations-path", "m", "", "path to migrations directory")
	migrateDatabaseCMD.Flags().StringVarP(&migrationsTable, "migrations-table", "t", "migrations", "database table holding migrations")
	migrateDatabaseCMD.Flags().IntVarP(&steps, "steps", "n", 0, "number of steps to migrate. positive steps for up and negative steps for down. zero to upgrade all.")
	rootCMD.AddCommand(migrateDatabaseCMD)
}

func getDriver() database.Driver {
	var driver database.Driver
	var err error
	cfg := config.LoadConfig(configPath)
	dtb := cfg.PostgresConnection
	driver, err = postgres.WithInstance(dtb, &postgres.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return driver
}

func migrateDB() {
	if migrationsPath == "" {
		log.Fatal("migrations path is required")
	}

	if !(strings.HasPrefix(migrationsPath, "/")) {
		path, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		migrationsPath, err = filepath.Abs(filepath.Join(path, migrationsPath))
		if err != nil {
			log.Fatal("cannot resolve full migration path")
		}
	}

	driver := getDriver()

	m, err := libmigrate.NewWithDatabaseInstance(
		"file://"+migrationsPath,
		config.CFG.Postgres.DBName,
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	m.Log = migrationLogger{}

	if steps == 0 {
		err = m.Up()
	} else {
		err = m.Steps(steps)
	}

	if err != nil {
		if err.Error() == "no change" {
			log.Println(err)
		} else {
			log.Fatal(err)
		}
	}
}

type migrationLogger struct{}

func (l migrationLogger) Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func (l migrationLogger) Verbose() bool { return true }
