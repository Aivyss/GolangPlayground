package sqlx

import "github.com/jmoiron/sqlx"

var Db *sqlx.DB

func BuildDbConn() {
	Db = SetDB()
}
