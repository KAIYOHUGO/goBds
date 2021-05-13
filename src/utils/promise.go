package utils

import (
	"reflect"
)

type Promise <-chan interface{}

// trun func to a await chan
func Async(fn func(v ...interface{}) (o interface{}, err error), v ...interface{}) Promise {
	r := make(chan interface{})
	go func() {
		defer close(r)
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
func Await(v Promise) interface{} {
	r := <-v
	return r
}

// Arace race data & return first data
func Arace(v ...Promise) Promise {
	r := make(chan interface{})
	go func() {
		defer close(r)
		c := make([]reflect.SelectCase, len(v))
		for i, el := range v {
			c[i].Dir = reflect.SelectRecv
			c[i].Chan = reflect.ValueOf(el)
		}
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
