package database

import (
	"gobds/src/config"
	"math/rand"
	"testing"
	"time"
)

func TestEndecode(t *testing.T) {
	b, err := Encode(config.Account{Name: "Paula"})
	if err != nil {
		panic(err)
	}
	s := config.Account{}
	err = Decode(b, &s)
	if err != nil {
		panic(err)
	}
	t.Log(s.Name)
	GC()
}

func TestSession(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	s, err := NewSession(config.Account{Name: "Paula"})
	if err != nil {
		panic(err)
	}
	t.Log(s)
	k, err := GetSession(s)
	if err != nil {
		panic(err)
	}
	t.Log(k.Name)
	DelSession(s)
	k, err = GetSession(s)
	// should return error
	if err == nil {
		panic(err)
	}
	t.Log("not exist")
	GC()
}

func TestReadWrite(t *testing.T) {
	Write(DB["account"], "Paula", config.Account{Name: "Paula", Password: "0623"})
	var s config.Account
	err := Read(DB["account"], "Paula", &s)
	if err != nil {
		panic(err)
	}
	t.Log(s.Name, "=>", s.Password)
	GC()
}
