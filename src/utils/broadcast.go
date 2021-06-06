package utils

import (
	"container/list"
	"gobds/src/config"
	"sync"
)

// Broadcast object
type Broadcast struct {
	list *list.List
	lock sync.Mutex
}

func NewBroadcast() *Broadcast {
	return &Broadcast{list: list.New().Init()}
}

// add a chan into broadcast.list
func (s *Broadcast) New() (Promise, *list.Element) {
	s.lock.Lock()
	defer s.lock.Unlock()
	v := make(chan interface{}, config.ChannelSize)
	el := s.list.PushBack(v)
	return v, el
}

// add a chan into broadcast.list
func (s *Broadcast) Close(v *list.Element) {
	s.lock.Lock()
	defer s.lock.Unlock()
	select {
	case _, ok := <-v.Value.(chan interface{}):
		if ok {
			close(v.Value.(chan interface{}))
		}
	default:
		close(v.Value.(chan interface{}))
	}
	s.list.Remove(v)
}

// send messenge into chan
func (s *Broadcast) Say(v interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	for i := s.list.Front(); i != nil; i = i.Next() {
		select {
		case i.Value.(chan interface{}) <- v:
		default:
			s.Close(i)
		}
	}
}
