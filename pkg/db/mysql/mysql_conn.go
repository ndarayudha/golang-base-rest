package mysql

import (
	"database/sql"
	"fmt"
	"golang_restfull_api/config"
	"golang_restfull_api/pkg/utils"

	"time"
)

const (
	maxIddleConns = 5
	maxOpenConns  = 20
	maxLifeTime   = 60 * time.Minute
	maxIddleTime  = 10 * time.Minute
)

func NewDB(c *config.Config) *sql.DB {
	connStr := fmt.Sprintf("%s:%s@%s(%s:%s)/%s",
		c.Mysql.MysqlUser,
		c.Mysql.MysqlPassword,
		c.Mysql.MysqlNetwork,
		c.Mysql.MysqlHost,
		c.Mysql.MysqlPort,
		c.Mysql.MysqlDbName,
	)

	db, err := sql.Open(c.Mysql.MysqlDriver, connStr)
	utils.PanicIfError(err)

	db.SetMaxIdleConns(maxIddleConns)
	db.SetMaxOpenConns(maxOpenConns)
	db.SetConnMaxLifetime(maxLifeTime)
	db.SetConnMaxIdleTime(maxIddleTime)

	if err = db.Ping(); err != nil {
		utils.PanicIfError(err)
	}

	return db
}
