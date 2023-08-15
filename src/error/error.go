package pError

import (
	"fmt"
	"runtime"
)

type PlaygroundError struct {
	Msg      string
	StackMsg []string
}

func (e *PlaygroundError) Error() string {
	return e.Msg
}

func NewPlayGroundError(msg string) *PlaygroundError {
	const depth = 10
	var stackMsg []string
	callers := make([]uintptr, depth)
	n := runtime.Callers(0, callers)
	callers = callers[:n]

	for _, stack := range callers {
		fn := runtime.FuncForPC(stack)
		file, line := fn.FileLine(stack)
		stackMsg = append(stackMsg, fmt.Sprintf("%s:%d %s\n", file, line, fn.Name()))
	}

	return &PlaygroundError{
		StackMsg: stackMsg,
		Msg:      msg,
	}
}
