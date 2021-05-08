package errs

import "fmt"

type CodeError struct {
	error string
	code  int
}

var errs map[int]*CodeError = make(map[int]*CodeError)

func (c CodeError) Error() string {
	if c.code != 200 && c.code != 0 {
		return fmt.Sprintf(" %s [-- ErrorCode=%d --]", c.error, c.code)
	} else {
		return ""
	}
}

func (c CodeError) Append(e error) error {
	c.error = e.Error() + " :" + c.error
	return c
}

func (c CodeError) Add(e error) *CodeError {
	c.error = c.error + " :" + e.Error()
	return &c
}
func (c CodeError) AddByString(s string) *CodeError {
	c.error = c.error + " :" + s
	return &c
}

func (c CodeError) Code() int {
	return c.code
}

func New(s string, code int) *CodeError {
	return &CodeError{s, code}
}

func Definition(code int, message string) *CodeError {
	err := CodeError{message, code}
	errs[code] = &err
	return &err
}
