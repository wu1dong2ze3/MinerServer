package utils

import "regexp"

var IsLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
var IsLetterAndNumber = regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString

func Len5_12(s string) bool {
	if len(s) >= 5 && len(s) <= 12 {
		return true
	}
	return false
}
