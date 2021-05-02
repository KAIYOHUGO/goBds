package utils

import (
	"encoding/base64"
	"errors"
	"gobds/src/config"
	"math/rand"
	"sync"
	"time"
)

var (
	Session = &session{}
)

type session struct {
	list map[string]User
	sync.Mutex
}

// User ...
// ID is a map[string] in leveldb

// Get ...
// get session , return user struct
func (s *session) Get(v string) (User, error) {
	s.Lock()
	l, ok := s.list[v]
	if !ok {
		return User{}, errors.New("not find")
	}
	if l.time+config.MaxSessionLiveTime < time.Now().Unix() {
		s.Unlock()
		s.Del(v)
		return User{}, errors.New("died")
	}
	l.time = time.Now().Unix()
	s.Unlock()
	return l, nil
}

// Add ...
// add a sesson ,return session id
func (s *session) Add() (string, error) {
	token := make([]byte, 64)
	if _, err := rand.Read(token); err != nil {
		return "", err
	}
	n := base64.URLEncoding.EncodeToString(token)
	s.Lock()
	_, ok := s.list[n]
	s.Unlock()
	if !ok {
		return "", errors.New("aready exist")
	}
	s.list[n] = User{
		// wip
		// info: &db.User{},
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

func CheckTime() {
	for t := range time.NewTimer(time.Duration(config.MaxSessionLiveTime)).C {
		for i, el := range Session.list {
			if el.time+config.MaxSessionLiveTime < t.Unix() {
				Session.Del(i)
			}
		}
	}
}

type User struct {
	time int64
	// info *db.User
}
