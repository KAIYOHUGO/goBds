package console

import (
	"gobds/src/config"
	"testing"
)

func TestConsole(t *testing.T) {
	s, err := NewConsole(config.TestServerFile)
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
	w := NewWrapper(config.TestServerFile)
	p, l := w.Broadcast.New()
	defer w.Broadcast.Close(l)
	go func() {
		for v := range p {
			t.Logf("%+v\n", v.(Log))
			t.Log(w.Status)
		}
	}()
	t.Log(w.Status)
	w.InputQueue("$start")
	w.InputQueue("stop")
	t.Log(<-w.Err())
	w.InputQueue("$start")
	w.InputQueue("stop")
	t.Log(<-w.Err())
}

func TestUnFindFileWrapper(t *testing.T) {
	w := NewWrapper("unfind file")
	p, l := w.Broadcast.New()
	defer w.Broadcast.Close(l)
	go func() {
		for v := range p {
			t.Logf("%+v\n", v.(Log))
			t.Log(w.Status)
		}
	}()
	t.Log(w.Status)
	w.InputQueue("$start")
	t.Log(<-w.Err())
}
