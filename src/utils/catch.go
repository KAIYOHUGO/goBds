package utils

import (
	"errors"
	"gobds/src/config"
)

type Catch struct {
	list map[string]interface{}
	ram  int
}

func (s *Catch) Add(n string, v interface{}) {
	s.list[n] = v
	s.ram++
}
func (s *Catch) Get(v string) (interface{}, error) {
	r, ok := s.list[v]
	if !ok {
		return "", errors.New("")
	}
	return r, nil
}
func (s *Catch) Del(v string) {
	delete(s.list, v)
}

func (s *Catch) GC() {
	if s.ram > config.MaxCatchRam {

	}
}
