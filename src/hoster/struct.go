package hoster

import (
	"bufio"
	"container/list"
	"errors"
	"gobds/src/usefull"
	"io"
	"os/exec"
)

// List ...
// server info
type List struct {
	Path string
	proc *exec.Cmd
	// Status level
	// 0 close or start fail
	// 1 runing
	// 2 cmd error
	// 3 GC error
	// 4 kill error
	Status    uint8
	out       io.ReadCloser
	in        io.WriteCloser
	Broadcast *usefull.Broadcast
	CmdChan   chan string
}

func (s *List) init() {
	s.CmdChan = make(chan string, 10)
	s.Broadcast = &usefull.Broadcast{List: list.New()}
	go func() {
		for {
			s.setup()
		}
	}()
}

// setup ...
// setup server; need setup before run any mothed
func (s *List) setup() {
	usefull.Log("run setup")
	s.proc = exec.Command(s.Path)
	s.out, _ = s.proc.StdoutPipe()
	s.in, _ = s.proc.StdinPipe()
	s.Status = 0
	wg := make(chan struct{})
	defer func() {
		usefull.Wan("server stop")
		s.out.Close()
		s.in.Close()
		close(wg)
		if s.Status > 0 {
			s.kill()
		}
	}()
	go func() {
		o := bufio.NewScanner(s.out)
		for o.Scan() {
			s.Broadcast.Say(o.Text())
		}
		if o.Err() != nil {
			usefull.Log("close")
		}
	}()
	for v := range s.CmdChan {
		select {
		case <-wg:
			return
		default:
		}
		usefull.Log("CMD:" + v)
		if v[0] == '$' {
			switch v[1:] {
			case "start":
				if s.Status == 0 {
					if err := s.proc.Start(); err == nil {
						s.Status = 1
						go func() {
							if s.proc.Wait() != nil {
								s.Status = 3
							}
							s.Status = 0
							wg <- struct{}{}
						}()
					}
				}
			case "restart":
				if s.Status > 0 {
					func() {
						for {
							select {
							case <-s.CmdChan:
							default:
								return
							}
						}
					}()
					s.CmdChan <- "$stop"
					s.CmdChan <- "$kill"
					s.CmdChan <- "$start"
				}

			case "kill":
				if s.Status > 0 {
					s.kill()
				}
			}
		} else {
			s.cmd(v)
		}
	}
}

// kill ...
// kill process
func (s *List) kill() error {
	if err := s.proc.Process.Kill(); err != nil {
		usefull.Err("kill fail", err)
		s.Status = 4
		return errors.New("cant kill")
	}
	return nil
}

// cmd ...
// run cmd in terminal
func (s *List) cmd(c string) error {
	if _, err := s.in.Write([]byte(c + "\n")); err != nil {
		usefull.Err("cmd error", err)
		s.Status = 2
		return errors.New("unknow cmd error")
	}
	usefull.Log("run cmd:" + c)
	return nil
}
