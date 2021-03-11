package cnf

import "github.com/jmoiron/sqlx"

type Config struct {
	DB *sqlx.DB
}
