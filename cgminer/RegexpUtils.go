package cgminer

import "regexp"

var RegexpPoolUrl = regexp.MustCompile(`^((http|stratum\+tcp)?:\/\/)[^\s]+`)
var RegexpUser = regexp.MustCompile(`[A-Za-z0-9_\-]+`)
