package trace

import (
	"runtime"
)

type Trace struct {
	File     string
	Line     int
	Function string
}

func Parse() Trace {
	pc := make([]uintptr, 10)
	runtime.Callers(3, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	return Trace{
		File:     file,
		Line:     line,
		Function: f.Name(),
	}
}
