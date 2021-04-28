package usefull

import (
	"container/list"
	"sync"
)

type Broadcast struct {
	List *list.List
	sync.Mutex
}

// func (s *Broadcast) Init() {
// s.list = list.New()
// }
func (s *Broadcast) Add(v chan interface{}) *list.Element {
	return s.List.PushBack(v)
}
func (s *Broadcast) Del(v *list.Element) {
	s.List.Remove(v)
}
func (s *Broadcast) Say(v interface{}) {
	s.Lock()
	for i := s.List.Front(); i != nil; i = i.Next() {
		select {
		case i.Value.(chan interface{}) <- v:
		default:
			s.Del(i)
		}
	}
	s.Unlock()
}
