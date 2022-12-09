package database

import (
	"database/sql"
	"github.com/faelp22/tcs_curso/projetogolangcerto/config"
	"stoq/config"
)

type DatabaseInterface interface {
	GetDB() (DB *sql.DB)
	Close() error
}

type dabase_pool struct {
	DB *sql.DB
}

var dbpool = &dabase_pool{}

func NewDB(conf *config.DBConfig) *dabase_pool {
	println(conf.DB_DRIVE)
	dbpool = MySQLConn(conf)
	return dbpool
}

func (d *dabase_pool) Close() error {

	err := d.DB.Close()
	if err != nil {
		return err
	}

	dbpool = &dabase_pool{}

	return err
}

func (d *dabase_pool) GetDB() (DB *sql.DB) {
	return d.DB
}
