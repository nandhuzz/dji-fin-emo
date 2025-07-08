package db

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // required for file:// migrations
	_ "github.com/jackc/pgx/v5/stdlib"

	"context"
	"database/sql"

	"github.com/jackc/pgx/v5/pgxpool"
)

func RunMigrations(dbPool *pgxpool.Pool, migrationsPath string) {
	// Convert pgxpool.Pool to *sql.DB for migrate compatibility
	stdDB := stdlibOpen(dbPool)

	driver, err := postgres.WithInstance(stdDB, &postgres.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to create migration driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+migrationsPath,
		"postgres", driver,
	)
	if err != nil {
		log.Fatalf("❌ Failed to create migrator: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("❌ Migration failed: %v", err)
	}

	log.Println("✅ Database migrated successfully")
}

func stdlibOpen(pool *pgxpool.Pool) *sql.DB {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Fatalf("Failed to acquire connection: %v", err)
	}
	defer conn.Release()

	// Create a stdlib-compliant connection using the DSN
	dsn := pool.Config().ConnString()
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("Failed to open stdlib connection: %v", err)
	}
	return db
}
