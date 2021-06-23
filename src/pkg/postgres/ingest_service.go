package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/raghanag/my-project/pkg/weather"
)

type IngestService struct {
	DbUserName string
	DbPassword string
	DbURL      string
	DbName     string
}

func (t *IngestService) Initialise() error {
	targetSchemaVersion := 1
	dbConnectionString := t.getDBConnectionString()
	db, err := sql.Open("pgx", dbConnectionString)
	if err != nil {

		return err
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}
	version, dirty, err := driver.Version()

	if dirty {
		log.Fatalf("ERROR: The current database schema is reported as being dirty. A manual resolution is needed")
	}
	log.Printf("Target database schema version is: %d and current database schema version is: %d", targetSchemaVersion, version)
	if version != targetSchemaVersion {
		log.Printf("Migrating database schema from version: %d to version %d", version, targetSchemaVersion)
		m, err := migrate.NewWithDatabaseInstance("file://../../pkg/postgres/migrations", t.DbName, driver)
		if err != nil {
			return err
		}
		err = m.Steps(targetSchemaVersion)
		if err != nil {
			return err
		}
		return nil
	} else {
		log.Println("No database schema migrations need to be performed.")
	}
	if err != nil {
		log.Fatalf("ERROR: Could not determine the current database schema version")
	}

	return nil
}

func (t *IngestService) Create() (*string, error) {
	error := "Ingestion Error"
	minutely := weather.QueryWeather()
	if len(minutely) == 0 {
		return &error, nil
	}
	insertSQL := "insert into weather_ts.minutely_weather(dt, data_type, value) values ($1, $2, $3)"
	ctx := context.Background()
	dbPool := t.getConnection()
	defer dbPool.Close()
	tx, err := dbPool.Begin(ctx)
	if err != nil {
		return &error, err
	}

	for i := 0; i < len(minutely); i++ {
		_, err = tx.Exec(ctx, insertSQL, time.Unix(int64(minutely[i].Dt), 0), "precipitation", minutely[i].Precipitation)
		if err != nil {
			log.Println("ERROR: Could not save the item due to the error:", err)
			rollbackErr := tx.Rollback(ctx)
			log.Fatal("ERROR: Transaction rollback failed due to the error: ", rollbackErr)
			return &error, nil
		}

	}
	err = tx.Commit(ctx)
	if err != nil {
		return &error, err
	}
	succcess := "Ingestion Success"
	return &succcess, nil
}

func (t *IngestService) getConnection() *pgxpool.Pool {
	dbPool, err := pgxpool.Connect(context.Background(), t.getDBConnectionString())
	if err != nil {
		log.Fatal(err)
	}
	return dbPool
}

func (t *IngestService) getDBConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", t.DbUserName, t.DbPassword, t.DbURL, t.DbName)
}
