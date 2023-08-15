package main

import (
	log "com.playground/logger"
	"fmt"
	"os"
)

func main() {
	logger, logFile := log.GetLogger()
	defer logFile.Close()

	file, err := os.Open("no-file.txt")
	if err != nil {
		//logger.Println("error happened", err)
		//logger.Fatalln("fatal error happened", err)
		logger.Panicln("panic error happened", err)
	}
	defer file.Close()

	fmt.Println("about to exit")
}
