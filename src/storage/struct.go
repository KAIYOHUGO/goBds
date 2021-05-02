package storage

import "os"

type Storage struct {
	Struct interface{}
	file   *os.File
}
