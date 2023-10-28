package commander

import "errors"

var ErrBadArgs = errors.New("bad args")
var ErrCmdNotFound = errors.New("command not found")
