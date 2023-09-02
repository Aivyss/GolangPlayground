package sqlx

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL
)

const dbDriver = "postgres"
const datasourceName = "user=test1 password=test1 dbname=test1 sslmode=disable"
const maxConnSize = 10
const maxIdleSize = 5

func SetDB() *sqlx.DB {
	db := sqlx.MustConnect(dbDriver, datasourceName)

	db.SetMaxOpenConns(maxConnSize)
	db.SetMaxIdleConns(maxIdleSize)

	return db
}
