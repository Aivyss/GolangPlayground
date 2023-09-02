package main

import "com.playground/24_echo/router"

func main() {
	e := router.GetRouter()

	_ = e.Start(":8080")
}
