package hoster

import (
	"bufio"
	"container/list"
	"errors"
	"gobds/src/usefull"
	"io"
	"os/exec"
	"strconv"
)

// List ...
// server info
type List struct {
	Path      string
	proc      *exec.Cmd
	Status    int8
	out       io.ReadCloser
	in        io.WriteCloser
	Broadcast *usefull.Broadcast
	EventChan chan string
	CmdChan   chan string
}

func (s *List) init() {
	s.EventChan = make(chan string, 5)
	s.CmdChan = make(chan string, 5)
	s.Broadcast = &usefull.Broadcast{List: list.New()}
}

// setup ...
// setup server; need setup before run any mothed
func (s *List) setup() {
	usefull.Log("run setup")
	s.proc = exec.Command(s.Path)
	s.out, _ = s.proc.StdoutPipe()
	s.in, _ = s.proc.StdinPipe()
	s.Status = 0
	gc := func() {
		s.out.Close()
		s.in.Close()
		if s.Status != 0 {
			s.kill()
		}
	}
	defer gc()
	for {
		if t := <-s.EventChan; t == "start" {
			break
		}
	}
	s.Status = 1
	// start proc
	if err := s.proc.Start(); err != nil {
		usefull.Log(strconv.Itoa(int(s.Status)))
		usefull.Err("cant start", err)
		s.Status = -1
	}
	wg := make(chan struct{})
	usefull.Log("start!")
	go func() {
		o := bufio.NewScanner(s.out)
		for o.Scan() {
			s.Broadcast.Say(o.Text())
			println(o.Text())
		}
		if o.Err() != nil {
			usefull.Log("close")
		}
	}()
	go func() {
		for {
			select {
			case t := <-s.EventChan:
				switch t {
				case "stop":
					if s.cmd("stop") == nil {
						return
					}
				case "kill":
					if s.kill() == nil {
						return
					}
				case "restart":
					go func() {
						s.EventChan <- "stop"
						s.EventChan <- "kill"
						s.EventChan <- "start"
					}()
				}
			case t := <-s.CmdChan:
				if s.cmd(t) != nil {
					return
				}
			case <-wg:
				return
			}
		}
	}()
	s.proc.Wait()
	wg <- struct{}{}
	usefull.Wan("server stop")
	gc()
	s.setup()
}

// kill ...
// kill process
func (s *List) kill() error {
	if s.Status == 0 {
		usefull.Wan("kill status==0")
		return errors.New("server is close")
	}
	if err := s.proc.Process.Kill(); err != nil {
		usefull.Err("kill fail", err)
		s.Status = -1
		return errors.New("cant kill")
	}
	return nil
}

// cmd ...
// run cmd in terminal
func (s *List) cmd(c string) error {
	if s.Status <= 0 {
		usefull.Wan("cmd status<=0" + strconv.Itoa(int(s.Status)))
		return errors.New("server is close")
	}
	if _, err := s.in.Write([]byte(c + "\n")); err != nil {
		usefull.Err("cmd error", err)
		s.Status = -2
		return errors.New("unknow cmd error")
	}
	return nil
}
