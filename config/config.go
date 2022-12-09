package config

import (
	// "os"
)

type DBConfig struct {
	DB_DRIVE string `json:"db_drive"`
	DB_HOST  string `json:"db_host"`
	DB_PORT  string `json:"db_port"`
	DB_USER  string `json:"db_user"`
	DB_PASS  string `json:"db_pass"`
	DB_NAME  string `json:"db_name"`
	DB_DSN   string `json:"-"`
}

func getConfig() *DBConfig {
	default_conf := DBConfig{
			DB_DRIVE: "mysql",
			DB_HOST: "deposito-db.cmfdgmbsgaqu.sa-east-1.rds.amazonaws.com",
			DB_PORT: "3306",
			DB_USER: "admin",
			DB_PASS: "cursogolang",
			DB_NAME: "deposito",
			DB_DSN: "admin:cursogolang@tcp(deposito-db.cmfdgmbsgaqu.sa-east-1.rds.amazonaws.com:3306)/deposito",
	}
	return &default_conf
}
