package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

func main() {
	a := gob.NewEncoder(os.Stdout)
	o := &struct {
		test string
	}{
		test: "hi",
	}
	fmt.Println(o)
	a.Encode(&o)
	fmt.Println(o)
}
