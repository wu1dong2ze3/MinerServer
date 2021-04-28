package utils

import "time"

func AddBaseOnNow(second int) *time.Time {
	res := time.Now().Add(time.Duration(second) * time.Second)
	return &res
}
