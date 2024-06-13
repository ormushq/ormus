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
	pc := make([]uintptr, pcSize)
	n := runtime.Callers(runtimeCallerSkip, pc)
	pc = pc[:n]
	frames := runtime.CallersFrames(pc)
	frame, _ := frames.Next()
	return Trace{
		File:     frame.File,
		Line:     frame.Line,
		Function: frame.Function,
	}
}
