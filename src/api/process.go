package api

import (
	"fmt"
	"gobds/src/msg"
	"io"
	"os"
	"os/exec"
	"strconv"

	"github.com/gorilla/websocket"
)

// Plist ...
// server info
type Plist struct {
	Path   string
	Proc   *exec.Cmd
	Out    io.ReadCloser
	In     io.WriteCloser
	Status int8
}

// Setup ...
// setup server; need setup before run any mothed
func (s *Plist) Setup() {
	msg.Log("run setup")
	s.Proc = exec.Command(s.Path)
	s.Out, _ = s.Proc.StdoutPipe()
	s.In, _ = s.Proc.StdinPipe()
	s.Status = 0
	io.Copy(os.Stdout, s.Out)
}

// Start ...
// start server & wait it off
func (s *Plist) Start() {
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
func (s *Plist) Stop() {
	s.Cmd("stop")
}

// Kill ...
// kill process
func (s *Plist) Kill() {
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
func (s *Plist) Cmd(c string) {
	if s.Status <= 0 {
		msg.Wan("cmd status<=0" + strconv.Itoa(int(s.Status)))
		return
	}
	if _, err := s.In.Write([]byte(c + "\n")); err != nil {
		msg.Err("cmd error", err)
		s.Status = -2
		return
	}
}

type Pwss struct {
	Ws   *websocket.Conn
	Send chan interface{}
	Get  chan interface{}
}

// Prun ...
// idk
type Prun struct {
	Name string
	Cmd  string
}
