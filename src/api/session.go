package api

import (
	"encoding/base64"
	"math/rand"
	"time"

	"github.com/syndtr/goleveldb/leveldb/errors"
)

type Session struct {
	list map[string]User
}
type User struct {
	Id   []byte
	Info interface{}
}

func (s *Session) Get(v string) (l User, err error) {
	l, ok := s.list[v]
	if !ok {
		err = errors.New("not find")
		return
	}
	err = nil
	return
}
func (s *Session) Add() (string, error) {
	rand.Seed(time.Now().UnixNano())
	token := make([]byte, 64)
	if _, err := rand.Read(token); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(token), nil
}
func (s *Session) Del(v string) error {
	_, ok := s.list[v]
	if !ok {
		return errors.New("not find")
	}
	delete(s.list, v)
	return nil
}
