package usefull

import (
	"encoding/base64"
	"gobds/src/config"
	"math/rand"
	"time"

	"github.com/syndtr/goleveldb/leveldb/errors"
)

var (
	Session = &session{}
)

type session struct {
	list map[string]User
}

// User ...
// ID is a map[string] in leveldb

// Get ...
// get session , return user struct
func (s *session) Get(v string) (User, error) {
	l, ok := s.list[v]
	if !ok {
		return User{}, errors.New("not find")
	}
	if l.time+config.MaxSessionLiveTime < time.Now().Unix() {
		s.Del(v)
		return User{}, errors.New("died")
	}
	return l, nil
}

// Add ...
// add a sesson ,return session id
func (s *session) Add(v string) (string, error) {
	token := make([]byte, 64)
	if _, err := rand.Read(token); err != nil {
		return "", err
	}
	n := base64.URLEncoding.EncodeToString(token)
	_, ok := s.list[n]
	if !ok {
		return "", errors.New("aready exist")
	}
	s.list[n] = User{
		Name: v,
		time: time.Now().Unix(),
	}

	return n, nil
}

// Del ...
// delete a sesson ,return nil or err
func (s *session) Del(v string) error {
	_, ok := s.list[v]
	if !ok {
		return errors.New("not find")
	}
	delete(s.list, v)
	return nil
}

type User struct {
	Name string
	time int64
}

func CheckTime() {
	for t := range time.NewTimer(time.Duration(config.MaxSessionLiveTime)).C {
		for i, el := range Session.list {
			if el.time+config.MaxSessionLiveTime < t.Unix() {
				Session.Del(i)
			}
		}
	}
}
