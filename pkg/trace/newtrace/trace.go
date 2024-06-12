package newtrace

import (
	"runtime"
)

type Trace struct {
	File     string
	Line     int
	Function string
}

func Parse(runtimeCallerSkip int) Trace {
	pcSize := 10
	if runtimeCallerSkip == 0 {
		runtimeCallerSkip = 5
	}
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
