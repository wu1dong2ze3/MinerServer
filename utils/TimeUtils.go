package utils

import "time"

const oneDay = 24 * 3600 //一天多少秒
const oneHour = 3600
const oneMin = 60

func AddBaseOnNow(second int) *time.Time {
	res := time.Now().Add(time.Duration(second) * time.Second)
	return &res
}

func DayBySecondRound(sec int) int {
	return sec / oneDay
}

func HourBySecondMod(sec int) int {
	return (sec % oneDay) / oneHour
}

func MinBySecondMod(sec int) int {
	return ((sec % oneDay) % oneHour) / oneMin
}

func SecondBySeccondMod(sec int) int {
	return ((sec % oneDay) % oneHour) % oneMin
}
