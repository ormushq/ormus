package migrate

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/cassandra"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"os"
)

type Manager struct {
	dir      string
	host     string
	keyspace string
}

const MigrationDBURL = "cassandra://%s/%s"

func New(dir string, host string, keyspace string) *Manager {
	return &Manager{
		dir:      dir,
		host:     host,
		keyspace: keyspace,
	}
}

func (m *Manager) Run() error {
	sourceDriver, err := iofs.New(os.DirFS(m.dir), "migrations")
	if err != nil {
		return fmt.Errorf("create migrations source: %w", err)
	}

	mSource, err := migrate.NewWithSourceInstance("iofs", sourceDriver, fmt.Sprintf(MigrationDBURL, m.host, m.keyspace))
	if err != nil {
		return fmt.Errorf("create migration instance: %w", err)
	}

	err = mSource.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("apply migrations: %w", err)
	}

	err, _ = mSource.Close()
	if err != nil {
		return fmt.Errorf("close migration instance: %w", err)
	}

	return nil
}
