package hoster

import (
	"bufio"
	"fmt"
	"gobds/src/msg"
	"io"
	"os/exec"
	"strconv"
	"sync"
)

// List ...
// server info
type List struct {
	Path   string
	Proc   *exec.Cmd
	Status int8
	out    io.ReadCloser
	in     io.WriteCloser
	Broadcast
}

// Setup ...
// setup server; need setup before run any mothed
func (s *List) Setup() {
	msg.Log("run setup")
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
	msg.Log(strconv.Itoa(int(s.Status)))
	s.Status = 1
	fmt.Print("start server \n")
	if err := s.Proc.Start(); err != nil {
		msg.Log(strconv.Itoa(int(s.Status)))
		msg.Err("cant not start", err)
		s.Status = -1
		return
	}
	msg.Log(strconv.Itoa(int(s.Status)))
	s.Proc.Wait()
	msg.Wan("server stop")
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
		msg.Wan("kill status==0")
		return
	}
	if err := s.Proc.Process.Kill(); err != nil {
		msg.Err("kill fail", err)
		s.Status = -1
	}
}

// Cmd ...
// run cmd in terminal
func (s *List) Cmd(c string) {
	if s.Status <= 0 {
		msg.Wan("cmd status<=0" + strconv.Itoa(int(s.Status)))
		return
	}
	if _, err := s.in.Write([]byte(c + "\n")); err != nil {
		msg.Err("cmd error", err)
		s.Status = -2
		return
	}
}

type Broadcast struct {
	list []chan string
	sync.Mutex
}

func (s *Broadcast) Add(v chan string) {
	s.Mutex.Lock()
	s.list = append(s.list, v)
	s.Mutex.Unlock()
}
func (s *Broadcast) Del(v chan string) {
	s.Mutex.Lock()
	for i, el := range s.list {
		if el == v {
			s.list[i], s.list = s.list[len(s.list)-1], s.list[:len(s.list)-1]
		}
	}
	s.Mutex.Unlock()
}
func (s *Broadcast) Say(v string) {
	s.Mutex.Lock()
	for _, el := range s.list {
		el <- v
	}
	s.Mutex.Unlock()
}
