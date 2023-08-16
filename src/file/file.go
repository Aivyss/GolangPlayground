package file

import "os"

func GetFile(name string) (*os.File, error) {
	return os.OpenFile(
		name,
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0666,
	)
}
