package utils

import (
	"fmt"
	"strings"
)

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

/**
src:"2001 01 02 03"
Selet(" ","_",0,1,2)
res:"2001_01_02"
*/
func (s S) Select(splite string, insert string, sel ...int) S {
	arr := strings.Split(s.S, splite)
	//test
	fmt.Println("arr=", arr)
	for _, v := range arr {
		fmt.Println("--", v, "--")
	}
	fmt.Println("~~~~~~~~")

	//end
	var str = ""
	for i, v := range sel {
		if v >= len(arr) || v < 0 {
			return s
		}
		if i != len(sel)-1 {
			str += arr[v] + insert
		} else {
			str += arr[v]
		}
	}
	s.S = str
	return s
}

func (s S) RmTail(splite string) S {
	s.S = strings.TrimRight(s.S, splite)
	return s
}
