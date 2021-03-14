package usefull

import (
	"container/list"
	"sync"
)

type Broadcast struct {
	list *list.List
	sync.Mutex
}

func NewBroadcast() *Broadcast {
	return &Broadcast{list: list.New()}
}
func (s *Broadcast) Add(v chan interface{}) *list.Element {
	return s.list.PushBack(v)
}
func (s *Broadcast) Del(v *list.Element) {
	s.list.Remove(v)
}
func (s *Broadcast) Say(v interface{}) {
	s.Mutex.Lock()
	for i := s.list.Front(); i != nil; i = i.Next() {
		select {
		case <-i.Value.(chan interface{}):
			s.Del(i)
		default:
			i.Value.(chan interface{}) <- v
		}
	}
	s.Mutex.Unlock()
}
