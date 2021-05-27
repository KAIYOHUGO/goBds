package console

import (
	"gobds/src/config"
	"gobds/src/utils"
	"time"
)

type Wrapper struct {
	Broadcast *utils.Broadcast `json:"-"`
	Console   *Console         `json:"-"`
	File      string           `json:"file,omitempty"`
	// 3 stopping
	// 2 started
	// 1 starting
	// 0 init
	Status uint8 `json:"status,omitempty"`
	queue  chan string
	err    chan error
}

type Log struct {
	Time, Level, Output string
}

func NewWrapper(v string) *Wrapper {
	r := &Wrapper{
		Broadcast: utils.NewBroadcast(),
		File:      v,
		queue:     make(chan string, config.ChannelSize),
		err:       make(chan error),
	}
	go r.worker()
	return r
}
func (s *Wrapper) InputQueue(v ...string) bool {
	for _, i := range v {
		select {
		case s.queue <- i:
		default:
			return false
		}
	}
	return true
}
func (s *Wrapper) Err() error {
	return <-s.err
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
			if s.Status == 0 {
				go s.start()
			}
		case "$kill":
			if s.Status > 0 {
				if err := s.Console.Kill(); err != nil {
					s.error(err)
					break
				}
				s.Status = 0
			}
		default:
			if s.Status > 0 {
				s.Console.Input(i)
			}
		}
	}
}

func (s *Wrapper) start() error {
	var err error
	s.Console, err = NewConsole(s.File)
	if err != nil {
		return err
	}
	if err := s.Console.Start(); err != nil {
		return err
	}
	s.Status = 1
	defer s.Console.Kill()
	o := s.Console.Output()
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
			s.Status = 2
		case "Stopping server...":
			s.Status = 3
		}
		s.Broadcast.Say(r)
	}
	return s.Console.Wait()
}
