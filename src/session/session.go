package session

import (
	"gobds/src/config"
	"time"
)

var (
	Session = &session{}
)

func CheckTime() {
	for t := range time.NewTimer(time.Duration(config.MaxSessionLiveTime)).C {
		for i, el := range Session.list {
			if el.time+config.MaxSessionLiveTime < t.Unix() {
				Session.Del(i)
			}
		}
	}
}
