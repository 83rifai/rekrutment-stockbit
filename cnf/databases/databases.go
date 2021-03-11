package databases

import (
	"context"
	"fmt"
	"stockbit/cnf/env"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
)

var (
	mysqlMaster     *sqlx.DB
	lockMysqlMaster sync.Mutex
)

// CreateDBConnection function for creating database connection
func CreateDBConnection(descriptor string, maxIdle int, MaxOpen int) (*sqlx.DB, error) {
	db, err := sqlx.ConnectContext(context.Background(), "mysql", descriptor)
	if err != nil {

		defer db.Close()
		return db, err
	}

	db.SetMaxIdleConns(maxIdle)
	db.SetMaxOpenConns(MaxOpen)
	db.SetConnMaxLifetime(time.Second * 10)
	// defer db.Close()
	err = db.Ping()
	if err != nil {
		// log.WithField("error", constants.MYSQLErrorConnect).Error(err)
		log.Info("not connect database")
		log.Fatal(err)
		return db, err
	}

	return db, err
}

// MysqlDB function for creating database connection
func MysqlDB() (mysqlMaster *sqlx.DB, err error) {

	lockMysqlMaster.Lock()
	defer lockMysqlMaster.Unlock()

	if mysqlMaster == nil {
		mysqlMaster, err = CreateDBConnection(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", env.Conf.DBUser,
			env.Conf.DBPass, env.Conf.DBHost, env.Conf.DBPort, env.Conf.DBName), env.Conf.MaxIdle, env.Conf.MaxOpenConn)
	}

	return mysqlMaster, err

}
