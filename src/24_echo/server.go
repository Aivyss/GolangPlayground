package main

import (
	"com.playground/24_echo/router"
	dbPkg "com.playground/sqlx"
)

func main() {
	dbPkg.BuildDbConn()
	e := router.GetRouter()

	_ = e.Start(":8080")
}
