package utils

import (
	"container/list"
	"gobds/src/config"
	"sync"
)

// Broadcast object
type Broadcast struct {
	list *list.List
	sync.Mutex
}

func NewBroadcast() *Broadcast {
	return &Broadcast{list: list.New().Init()}
}

// add a chan into broadcast.list
func (s *Broadcast) Add() await {
	v := make(chan interface{}, config.ChannelBufferSize)
	s.list.PushBack(v)
	return v
}

// send messenge into chan
func (s *Broadcast) Say(v interface{}) {
	s.Lock()
	for i := s.list.Front(); i != nil; i = i.Next() {
		select {
		case i.Value.(chan interface{}) <- v:
		default:
			s.list.Remove(i)
		}
	}
	s.Unlock()
}
