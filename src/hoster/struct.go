package hoster

import (
	"bufio"
	"fmt"
	"gobds/src/usefull"
	"io"
	"os/exec"
	"strconv"
)

// List ...
// server info
type List struct {
	Path      string
	Proc      *exec.Cmd
	Status    int8
	out       io.ReadCloser
	in        io.WriteCloser
	Broadcast *usefull.Broadcast
}

// Setup ...
// setup server; need setup before run any mothed
func (s *List) Setup() {
	usefull.Log("run setup")
	s.Proc = exec.Command(s.Path)
	s.out, _ = s.Proc.StdoutPipe()
	s.in, _ = s.Proc.StdinPipe()
	s.Status = 0
	go func() {
		o := bufio.NewScanner(s.out)
		for o.Scan() {
			s.Broadcast.Say(o.Text())
		}
	}()
}

// Start ...
// start server & wait it off
func (s *List) Start() {
	if s.Status > 0 {
		return
	}
	usefull.Log(strconv.Itoa(int(s.Status)))
	s.Status = 1
	fmt.Print("start server \n")
	if err := s.Proc.Start(); err != nil {
		usefull.Log(strconv.Itoa(int(s.Status)))
		usefull.Err("cant not start", err)
		s.Status = -1
		return
	}
	usefull.Log(strconv.Itoa(int(s.Status)))
	s.Proc.Wait()
	usefull.Wan("server stop")
	s.Setup()
	return
}

// Stop ...
// stop server== .cmd("stop")
func (s *List) Stop() {
	s.Cmd("stop")
}

// Kill ...
// kill process
func (s *List) Kill() {
	if s.Status == 0 {
		usefull.Wan("kill status==0")
		return
	}
	if err := s.Proc.Process.Kill(); err != nil {
		usefull.Err("kill fail", err)
		s.Status = -1
	}
}

// Cmd ...
// run cmd in terminal
func (s *List) Cmd(c string) {
	if s.Status <= 0 {
		usefull.Wan("cmd status<=0" + strconv.Itoa(int(s.Status)))
		return
	}
	if _, err := s.in.Write([]byte(c + "\n")); err != nil {
		usefull.Err("cmd error", err)
		s.Status = -2
		return
	}
}
