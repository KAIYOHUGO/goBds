package utils

import (
	"errors"
	"math/rand"
	"testing"
	"time"
)

func TestAsyncAwait(t *testing.T) {
	rand.Seed(time.Now().Unix())
	a := Async(func(v ...interface{}) (o interface{}, err error) {
		o, err = rand.Intn(1000), nil
		time.Sleep(time.Duration(o.(int)))
		return
	})
	t.Log(Await(a))
}

func TestArace(t *testing.T) {
	rand.Seed(time.Now().Unix())
	a, b, c := Async(func(v ...interface{}) (o interface{}, err error) {
		o, err = rand.Intn(1000), nil
		t.Log(o)
		time.Sleep(time.Duration(o.(int)))
		return
	}), Async(func(v ...interface{}) (o interface{}, err error) {
		o, err = rand.Intn(1000), nil
		t.Log(o)
		time.Sleep(time.Duration(o.(int)))
		return
	}), Async(func(v ...interface{}) (o interface{}, err error) {
		o, err = rand.Intn(1000), nil
		t.Log(o)
		time.Sleep(time.Duration(o.(int)))
		return
	})
	t.Log(Await(Arace(a, b, c)))
}

func TestBroadcast(t *testing.T) {
	b := NewBroadcast()
	c, l := b.New()
	defer b.Close(l)
	go func() {
		for o := range c {
			t.Log(o)
		}
	}()
	b.Say("hello")
	b.Say("world")
	b.Say("why")
	b.Say("u")
	b.Say("do")
	b.Say("this")
	b.Say("to")
	b.Say("me")
	time.Sleep(time.Second)
}

func TestMsg(t *testing.T) {
	Log("hello world")
	Wan("hello world")
	Err("why", errors.New("u leave me"))
}

func TestSHA(t *testing.T) {
	if s := Sha1([]byte("paula")); s != "62e52d2ac616f25dfddd0968a947fa7e84e5c086" {
		t.Fatal(s)
	}
}
