package mysql

import (
	"fmt"
	"sync"

	"github.com/marmotedu/errors"
	"gorm.io/gorm"

	"goer-startup/internal/apiserver/store"
	"goer-startup/internal/pkg/logger"
	genericoptions "goer-startup/internal/pkg/options"
	"goer-startup/pkg/db"
)

type datastore struct {
	db *gorm.DB

	// can include two database instance if needed
	// docker *gorm.DB
}

func (ds *datastore) Close() error {
	db, err := ds.db.DB()
	if err != nil {
		return errors.Wrap(err, "get gorm db instance failed")
	}

	return db.Close()
}

var (
	mysqlFactory store.Factory
	once         sync.Once
)

// GetMySQLFactoryOr create mysql factory with the given config.
func GetMySQLFactoryOr(opts *genericoptions.MySQLOptions) (store.Factory, error) {
	if opts == nil && mysqlFactory == nil {
		return nil, fmt.Errorf("failed to get mysql store factory")
	}

	var err error
	var dbIns *gorm.DB
	once.Do(func() {
		options := &db.Options{
			Host:                  opts.Host,
			Username:              opts.Username,
			Password:              opts.Password,
			Database:              opts.Database,
			MaxIdleConnections:    opts.MaxIdleConnections,
			MaxOpenConnections:    opts.MaxOpenConnections,
			MaxConnectionLifeTime: opts.MaxConnectionLifeTime,
			LogLevel:              opts.LogLevel,
			Logger:                logger.New(opts.LogLevel),
		}
		dbIns, err = db.New(options)

		// uncomment the following line if you need auto migration the given models
		// not suggested in production environment.
		// migrateDatabase(dbIns)

		mysqlFactory = &datastore{dbIns}
	})

	if mysqlFactory == nil || err != nil {
		return nil, fmt.Errorf("failed to get mysql store factory, mysqlFactory: %+v, error: %w", mysqlFactory, err)
	}

	return mysqlFactory, nil
}

// cleanDatabase tear downs the database tables.
// nolint:unused // may be reused in the feature, or just show a migration usage.
func cleanDatabase(db *gorm.DB) error {
	return nil
}

// migrateDatabase run auto migration for given models, will only add missing fields,
// won't delete/change current data.
// nolint:unused // may be reused in the feature, or just show a migration usage.
func migrateDatabase(db *gorm.DB) error {
	return nil
}

// resetDatabase resets the database tables.
// nolint:unused // may be reused in the feature, or just show a migration usage.
func resetDatabase(db *gorm.DB) error {
	if err := cleanDatabase(db); err != nil {
		return err
	}
	if err := migrateDatabase(db); err != nil {
		return err
	}

	return nil
}
