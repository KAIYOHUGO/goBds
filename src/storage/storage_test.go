package storage

import "testing"

type Demo struct {
	Word string
}

func Test(t *testing.T) {
	db, err := Open("./test.json")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	s := Demo{Word: "hello world"}
	db.Write(s)
	db.Read(&s)
	t.Log(s)
}
