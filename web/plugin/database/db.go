package database

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"io"
	"study-go/web/common"
	"study-go/web/log"
	"study-go/web/plugin"
	"time"
)

type DBStorage interface {
	Transaction(func(*sqlx.Tx) error) error
	Exec(tx *sqlx.Tx, sql string, args ...interface{}) (sql.Result, error)
	Query(tx *sqlx.Tx, sql string, data interface{}, args ...interface{}) error
	BeginTx() (*sqlx.Tx, error)
	Commit(tx *sqlx.Tx)
	Rollback(tx *sqlx.Tx)

	io.Closer
}

type DB struct {
	db  *sqlx.DB
	cfg Config
	Log *log.Logger
}

func init() {
	plugin.RegisterFactory("database", New)
}

func New() (plugin.Plugin, error) {
	var cfg Config
	err := common.LoadConfig(&cfg)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return NewDB(cfg)
}

func NewDB(cfg Config) (*DB, error) {
	db, err := sqlx.Open(cfg.Database.Type, cfg.Database.URL)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	db.SetMaxOpenConns(cfg.Database.MaxOpenConns)
	db.SetMaxIdleConns(cfg.Database.MaxIdleConns)
	db.SetConnMaxLifetime(time.Duration(cfg.Database.ConnMaxLifetime) * time.Second)

	err = db.Ping()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &DB{
		db:  db,
		cfg: cfg,
		Log: log.With(log.Any("plugin", "database")),
	}, nil
}

func (d *DB) Close() (err error) {
	err = d.db.Close()
	err = errors.WithStack(err)
	return
}

func (d *DB) Transaction(handler func(*sqlx.Tx) error) (err error) {
	tx, err := d.db.Beginx()
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	err = handler(tx)
	err = errors.WithStack(err)
	return
}

func (d *DB) Exec(tx *sqlx.Tx, sql string, args ...interface{}) (res sql.Result, err error) {
	if tx == nil {
		res, err = d.db.Exec(sql, args...)
	} else {
		res, err = tx.Exec(sql, args...)
	}
	return
}

func (d *DB) Query(tx *sqlx.Tx, sql string, data interface{}, args ...interface{}) (err error) {
	if tx == nil {
		err = d.db.Select(data, sql, args...)
	} else {
		err = tx.Select(data, sql, args...)
	}
	err = errors.WithStack(err)
	return
}

func (d *DB) BeginTx() (*sqlx.Tx, error) {
	return d.db.Beginx()
}

func (d *DB) Commit(tx *sqlx.Tx) {
	if tx == nil {
		return
	}
	err := tx.Commit()
	if err != nil {
		log.Error(err)
	}
}

func (d *DB) Rollback(tx *sqlx.Tx) {
	if tx == nil {
		return
	}
	err := tx.Rollback()
	if err != nil {
		log.Error(err)
	}
}
