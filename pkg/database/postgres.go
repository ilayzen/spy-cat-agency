package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type DBConfig struct {
	Port               string `yaml:"port"`
	Username           string `yaml:"username"`
	Host               string `yaml:"host"`
	DBName             string `yaml:"dbName"`
	Password           string `yaml:"password"`
	SSLMode            string `yaml:"sslmode"`
	MaxIdleConnections int    `yaml:"maxIdleConnections"`
	MaxOpenConnections int    `yaml:"maxOpenConnections"`
}

type DB struct {
	sqlx.ExtContext
	closeFn   func() error
	begintxFn func(context.Context, *sql.TxOptions) (*sqlx.Tx, error)
}

func BuildPostgresURL(db DBConfig) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", db.Username, db.Password, db.Host, db.Port, db.DBName, db.SSLMode)
}


func NewDB(cfg DBConfig, log logrus.FieldLogger) (*DB, error) {
	connConfig, err := pgx.ParseConfig(BuildPostgresURL(cfg))
	if err != nil {
		return nil, err
	}

	connStr := stdlib.RegisterConnConfig(connConfig)

	db, err := sqlx.Connect("pgx", connStr)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(cfg.MaxIdleConnections)
	db.SetMaxOpenConns(cfg.MaxOpenConnections)

	return &DB{
		ExtContext: db,
		closeFn:    db.Close,
		begintxFn:  db.BeginTxx,
	}, nil
}

func (db *DB) Close() error {
	return db.closeFn()
}

func (db *DB) Beginx(ctx context.Context) (*sqlx.Tx, error) {
	return db.begintxFn(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
}

func (db *DB) RollbackTx(tx *sqlx.Tx) {
	err := tx.Rollback()
	if err != nil && !errors.Is(err, sql.ErrTxDone) {
		logrus.Infof("[RollbackTx] Transaction rollback error. Tx: %v, err: %v", tx, err)
	}
}
