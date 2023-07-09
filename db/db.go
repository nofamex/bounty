package db

import (
	"bounty/config"
	"context"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

var DB *pgxpool.Pool

func ConnectDatabase() error {
	pool, err := pgxpool.New(context.Background(), config.ENVConfig.DBURL)
	if err != nil {
		log.Error().Err(err).Msg("failed connecting to database")
		return err
	}

	DB = pool

	log.Info().Msg("success connecting to database")

	return nil
}

func MigrateDatabase() {
	m, err := migrate.New("file://db/migrations", config.ENVConfig.DBURL)
	if err != nil {
		log.Warn().Err(err).Msg("error migrating database")
	}
	m.Up()
	log.Info().Msg("success migrating database")
}
