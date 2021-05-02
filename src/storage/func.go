package storage

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func Open(v string) (*Storage, error) {
	file, err := os.OpenFile(v, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		return &Storage{}, err
	}
	return &Storage{file: file}, nil
}
func (s *Storage) Close() {
	s.file.Close()
}
func (s *Storage) Read(v interface{}) error {
	b, err := ioutil.ReadAll(s.file)
	if err != nil {
		return err
	}
	json.Unmarshal(b, &v)
	s.Struct = v
	return nil
}
func (s *Storage) Write(v interface{}) error {
	b, err := json.Marshal(v)
	println(string(b))
	if err != nil {
		return err
	}
	_, err = s.file.Write(b)
	return err
}
