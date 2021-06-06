package console

import (
	"gobds/src/config"
	"testing"
	"time"
)

func TestConsole(t *testing.T) {
	s, err := NewConsole(config.TestServerPath, config.TestServerPath+config.TestServerCommand)
	if err != nil {
		panic(err)
	}
	err = s.Start()
	if err != nil {
		panic(err)
	}
	defer s.Kill()
	go func() {
		o := s.Output()
		for o.Scan() {
			t.Log(o.Text())
		}
	}()
	s.Input("stop")
	s.Wait()
}

func TestWrapper(t *testing.T) {
	w := NewWrapper(config.TestServerPath, config.TestServerPath+config.TestServerCommand)
	p, l := w.Join()
	defer w.Leave(l)

	go func() {
		for v := range p {
			t.Logf("%+v\n", v.(Log))
			t.Log(w.Status)
		}
	}()
	go func() {
		for v := w.Err(); v != nil; v = w.Err() {
			t.Log("error", v)
		}
	}()
	t.Log(w.Status)
	w.InputQueue("$start")
	time.Sleep(time.Second)
	w.InputQueue("stop")
	time.Sleep(time.Second * 15)
}

func TestUnFindFileWrapper(t *testing.T) {
	w := NewWrapper("unfind file", "")
	p, l := w.Join()
	defer w.Leave(l)
	go func() {
		for v := range p {
			t.Logf("%+v\n", v.(Log))
			t.Log(w.Status)
		}
	}()
	go func() {
		for v := w.Err(); v != nil; v = w.Err() {
			t.Log("error", v)
		}
	}()
	t.Log(w.Status)
	w.InputQueue("$start")
	w.Command = config.TestServerPath
	time.Sleep(time.Second)
	w.InputQueue("$start")
	time.Sleep(time.Second * 20)
}
