package fileserver

import "errors"

const (
	ModeRead       int8 = 0
	ModeWrite      int8 = 1
	ModeCreateFile int8 = 2
	ModeCreateDir  int8 = 3
	ModeDelete     int8 = 4
	ModeRename     int8 = 5
)

var (
	ErrUnknowMode error = errors.New("unknow mode")
	ErrIsDir      error = errors.New("file is dir")
)
