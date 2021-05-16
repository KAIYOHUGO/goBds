package database

import (
	"bytes"
	"encoding/gob"
)

func Encode(t interface{}) ([]byte, error) {
	var b bytes.Buffer
	err := gob.NewEncoder(&b).Encode(t)
	return b.Bytes(), err
}

// t should be a pointer
func Decode(b []byte, t interface{}) error {
	return gob.NewDecoder(bytes.NewBuffer(b)).Decode(t)
}
