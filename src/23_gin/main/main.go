package main

import "com.playground/23_gin/router"

func main() {
	r := router.GetRouter()
	_ = r.Run(":8080")
}
