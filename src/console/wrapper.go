package console

import (
	"container/list"
	"errors"
	"gobds/src/config"
	"gobds/src/utils"
	"sync"
	"time"
)

type Wrapper struct {
	// set as exec.cmd(Command)
	Command string `json:"file,omitempty"`
	// set as exec.cmd.dir
	Path string `json:"path,omitempty"`
	// 3 stopping
	// 2 started
	// 1 starting
	// 0 init
	Status    uint8            `json:"status,omitempty"`
	broadcast *utils.Broadcast `json:"-"`
	console   *Console         `json:"-"`
	queue     chan string
	err       chan error
	lock      sync.Mutex
}

type Log struct {
	Time, Level, Output string
}

func NewWrapper(v string, c string) *Wrapper {
	r := &Wrapper{
		Path:      v,
		Command:   c,
		Status:    0,
		broadcast: utils.NewBroadcast(),
		queue:     make(chan string, config.ChannelSize),
		err:       make(chan error),
	}
	go r.worker()
	return r
}
func (s *Wrapper) InputQueue(v string) bool {
	select {
	case s.queue <- v:
	default:
		return false
	}
	return true
}
func (s *Wrapper) Err() error {
	return <-s.err
}
func (s *Wrapper) Join() (utils.Promise, *list.Element) {
	return s.broadcast.New()
}
func (s *Wrapper) Leave(v *list.Element) {
	s.broadcast.Close(v)
}
func (s *Wrapper) GC() {
	s.console.Kill()
	s = nil
}
func (s *Wrapper) error(v error) {
	if v == nil {
		return
	}
	select {
	case s.err <- v:
	default:
		<-s.err
		s.err <- v
	}
}
func (s *Wrapper) worker() {
	for i := range s.queue {
		switch i {
		case "$start":
			go func() {
				if err := s.start(); err != nil {
					s.error(err)
				}
			}()
		case "$kill":
			if err := s.console.Kill(); err != nil {
				s.error(err)
				break
			}
			s.Status = 0
		default:
			s.console.Input(i)
		}
	}
}

func (s *Wrapper) start() error {
	defer func() {
		s.Status = 0
	}()
	s.lock.Lock()
	if s.Status > 0 {
		s.lock.Unlock()
		return errors.New("already start")
	}
	s.Status = 1
	s.lock.Unlock()
	var err error
	s.console, err = NewConsole(s.Path, s.Command)
	if err != nil {
		return err
	}
	if err := s.console.Start(); err != nil {
		return err
	}
	defer s.console.Kill()
	o := s.console.Output()
	for o.Scan() {
		l, r := o.Text(), Log{}
		m := config.ConsoleOutput.FindStringSubmatch(l)
		r.Time = time.Now().Format("2006/01/02 15:04:05")
		if len(m) == 0 {
			r.Output = l
		} else {
			r.Level = m[config.ConsoleOutput.SubexpIndex("level")]
			r.Output = m[config.ConsoleOutput.SubexpIndex("output")]
		}
		switch r.Output {
		case "Server started.":
			s.lock.Lock()
			s.Status = 2
			s.lock.Unlock()
		case "Stopping server...":
			s.lock.Lock()
			s.Status = 3
			s.lock.Unlock()
		}
		s.broadcast.Say(r)
	}
	return s.console.Wait()
}
