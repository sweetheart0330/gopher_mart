package pg

import (
	"context"
	"database/sql"
	"fmt"
	"path/filepath"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sweetheart0330/gopher_mart/internal/repository"
	"go.uber.org/zap"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

const (
	postgreDriver = "postgres"
	migrationsDir = "migrations"
)

const (
	checkTablesIfExist = `SELECT
    EXISTS (
        SELECT 1 FROM information_schema.tables
        WHERE table_schema = 'public' AND table_name = 'users'
    ) AS users_exists,
    EXISTS (
        SELECT 1 FROM information_schema.tables
        WHERE table_schema = 'public' AND table_name = 'orders'
    ) AS orders_exists,
    EXISTS (
        SELECT 1 FROM information_schema.tables
        WHERE table_schema = 'public' AND table_name = 'balances'
    ) AS balances_exists,
    EXISTS (
        SELECT 1 FROM information_schema.tables
        WHERE table_schema = 'public' AND table_name = 'withdrawals'
    ) AS withdrawals_exists;`
)

type Database struct {
	pg *pgxpool.Pool
}

func NewDatabase(ctx context.Context, connStr string, log *zap.SugaredLogger) (repository.IRepository, error) {
	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		return nil, fmt.Errorf("could not connect to database: %w", err)
	}
	//	db := &Database{pg: pool}
	//	err = db.migrateTable(ctx, connStr, log)
	//	if err != nil {
	//		log.Errorw("could not migrate table", "error", err)
	//	}

	return &Database{pg: pool}, nil
}

func (db *Database) Close() {
	db.pg.Close()
}

func (db *Database) Ping(ctx context.Context) error {
	err := db.pg.Ping(ctx)
	if err != nil {
		return fmt.Errorf("could not ping database: %w", err)
	}

	return nil
}

func (db *Database) migrate(url string, log *zap.SugaredLogger) error {
	pg, err := sql.Open(postgreDriver, url)
	if err != nil {
		return err
	}
	driver, err := postgres.WithInstance(pg, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("create migrate driver: %w", err)
	}

	absPath, _ := filepath.Abs(migrationsDir)
	m, err := migrate.NewWithDatabaseInstance(
		"file://"+absPath,
		postgreDriver, driver,
	)
	if err != nil {
		return fmt.Errorf("new migrate instance: %w", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migrate up: %w", err)
	}

	log.Debug("migrated up")
	return nil
}

func (db *Database) migrateTable(ctx context.Context, url string, log *zap.SugaredLogger) error {
	var exists bool
	if err := db.pg.QueryRow(ctx, checkTablesIfExist).Scan(&exists); err != nil {
		return fmt.Errorf("check table exists: %w", err)
	}

	if !exists {
		err := db.migrate(url, log)
		if err != nil {
			return fmt.Errorf("failed to migrate, err: %w", err)
		}
	} else {
		log.Debug("skipping migration")
	}
	return nil
}
