package util

import (
	"log"
	"os"
)

func GetGoPath() string {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		log.Fatalln("GOPATH is not set")
	}

	return gopath
}
