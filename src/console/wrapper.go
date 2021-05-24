package console

import (
	"gobds/src/config"
	"gobds/src/utils"
	"time"
)

type Wrapper struct {
	Broadcast *utils.Broadcast
	Console   *Console
	File      string
	// 3 stopping
	// 2 started
	// 1 starting
	// 0 init
	Status uint8
	queue  chan string
	err    chan error
}

type Log struct {
	Time, Level, Output string
}

func NewWrapper(v string) *Wrapper {
	r := Wrapper{
		Broadcast: utils.NewBroadcast(),
		File:      v,
		queue:     make(chan string, config.ChannelSize),
		err:       make(chan error),
	}
	go func() {
		defer close(r.err)
		for {
			select {
			case r.err <- r.worker():
				r.Status = 0
			default:
				<-r.err
			}
		}
	}()
	return &r
}

func (s *Wrapper) worker() error {
	var (
		err error
	)
	quit := make(chan struct{})
	defer func() {
		close(quit)
		if err := recover(); err != nil {
			select {
			case s.err <- err.(error):
				s.Status = 0
			default:
				<-s.err
			}
		}
	}()
	s.Console, err = NewConsole(s.File)
	if err != nil {
		return err
	}
	defer s.Console.Kill()
	o := s.Console.Output()
	go func() {
		for {
			select {
			case i := <-s.queue:
				if i[0] == '$' {
					switch i[1:] {
					case "start":
						if s.Console.Start() == nil {
							s.Status = 1
						}
					case "kill":
						s.Console.Kill()
					}
					continue
				}
				s.Console.Input(i)
			case <-quit:
				return
			}
		}
	}()
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
func (s *Wrapper) Err() <-chan error {
	return s.err
}
