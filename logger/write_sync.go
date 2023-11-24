package logger

import (
	"io"
)

// A writeSyncer is an io.Writer that can also flush any buffered data. Note
// that *os.File (and thus, os.Stderr and os.Stdout) implement WriteSyncer.
type writeSyncer interface {
	io.Writer
	Sync() error
}

// AddSync converts an io.Writer to a WriteSyncer. It attempts to be
// intelligent: if the concrete type of the io.Writer implements WriteSyncer,
// we'll use the existing Sync method. If it doesn't, we'll add a no-op Sync.
func addSync(w io.Writer) writeSyncer {
	switch w := w.(type) {
	case writeSyncer:
		return w
	default:
		return writerWrapper{w}
	}
}

type writerWrapper struct {
	io.Writer
}

func (w writerWrapper) Sync() error {
	return nil
}
