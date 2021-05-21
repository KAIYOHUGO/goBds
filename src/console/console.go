package console

import (
	"bufio"
	"errors"
	"io"
	"os/exec"
)

type Console struct {
	proc   *exec.Cmd
	stdin  io.Writer
	stdout io.Reader
}

func NewConsole(v string) (*Console, error) {
	r := &Console{
		proc: exec.Command(v),
	}
	var err error
	r.stdin, err = r.proc.StdinPipe()
	if err != nil {
		return nil, err
	}
	r.stdout, err = r.proc.StdoutPipe()
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (s *Console) Start() error {
	return s.proc.Start()
}
func (s *Console) Kill() error {
	if s != nil && s.proc != nil {
		return s.proc.Process.Kill()
	}
	return errors.New("nil pointer")
}
func (s *Console) Input(v string) error {
	_, err := s.stdin.Write([]byte(v + "\n"))
	return err
}
func (s *Console) Output() *bufio.Scanner {
	return bufio.NewScanner(s.stdout)
}
func (s *Console) Wait() error {
	return s.proc.Wait()
}
