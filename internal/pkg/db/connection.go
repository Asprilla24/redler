package db

import (
	"fmt"
	"time"

	"redler/internal/pkg/config"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	dbTimeoutConnection = 30 * time.Second
)

// New creates new instance of DAO for database operations.
// dbDriver is database driver name,
// see http://doc.gorm.io/database.html#connecting-to-a-database for available database driver
func New(conf *config.Config) (*gorm.DB, error) {
	logrus.Print("Waiting database...")
	sslMode := "" // omit sslmode so it set to default (enabled)
	if !conf.DBSSLEnabled {
		sslMode = "sslmode=disable"
	}
	dns := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s connect_timeout=5 TimeZone=Asia/Jakarta "+sslMode,
		conf.DBHost, conf.DBPort, conf.DBName, conf.DBUsername, conf.DBPassword)

	timeout := time.Now().Add(dbTimeoutConnection)
	var postgresORM *gorm.DB
	var err error
	retryCounter := 0
	for time.Now().Before(timeout) {
		postgresORM, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
		if err == nil {
			break
		}
		retryCounter++
		logrus.Printf("database connection retry counter: %d", retryCounter)
	}
	if err != nil {
		return nil, errors.Wrapf(err, "unable to connect to database: timeout: %v", err)
	}
	if postgresORM == nil {
		return nil, errors.Wrapf(err, "unable to initiate DAO: %v", err)
	}

	err = RunMigrations(postgresORM)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to migrate the database: %v", err)
	}

	return postgresORM, nil
}
