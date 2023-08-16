package file

import "os"

func GetSimpleFile(name string) (*os.File, error) {
	return os.OpenFile(
		"log.log",
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0666,
	)
}
