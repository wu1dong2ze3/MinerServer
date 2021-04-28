package utils

import "strings"

type S struct {
	S string
}

func (s S) NoAny(any string) S {
	s.S = strings.ReplaceAll(s.S, any, "")
	return s
}
func (s S) NoSpace() S {
	s.S = strings.ReplaceAll(s.S, " ", "")
	return s
}
func (s S) NoBr() S {
	s.S = strings.ReplaceAll(s.S, "\n", "")
	return s
}
func (s S) NoQuot() S {
	s.S = strings.ReplaceAll(s.S, "\"", "")
	return s
}
func (s S) Replace(old string, new string) S {
	s.S = strings.ReplaceAll(s.S, old, new)
	return s
}
func (s S) Flip(sep string) S {
	arr := strings.Split(s.S, sep)
	if len(arr) != 2 {
		return s
	} else {
		s.S = arr[1] + sep + arr[0]
		return s
	}
}
func (s S) NoSpaceBr() S {
	return s.NoBr().NoSpace()
}
