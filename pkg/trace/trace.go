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
	pcSize := 10
	runtimeCallerSkip := 3
	pc := make([]uintptr, pcSize)
	runtime.Callers(runtimeCallerSkip, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])

	return Trace{
		File:     file,
		Line:     line,
		Function: f.Name(),
	}
}
