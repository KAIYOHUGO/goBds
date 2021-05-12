package utils

import (
	"reflect"
)

type promise func(v ...interface{}) (o interface{}, err error)
type await <-chan interface{}

// trun func to a await chan
func Async(fn promise, v ...interface{}) await {
	r := make(chan interface{})
	go func() {
		o, err := fn(v)
		if err != nil {
			r <- err
			return
		}
		r <- o
	}()
	return r
}

// await func done
func Await(v await) interface{} {
	r := <-v
	return r
}

// Arace race data & return first data
func Arace(v ...await) await {
	r := make(chan interface{})
	c := make([]reflect.SelectCase, len(v))
	for i, el := range v {
		c[i].Dir = reflect.SelectRecv
		c[i].Chan = reflect.ValueOf(el)
	}
	go func() {
		for {
			_, o, ok := reflect.Select(c)
			if ok {
				r <- o.Interface()
				return
			}
		}
	}()
	return r
}
