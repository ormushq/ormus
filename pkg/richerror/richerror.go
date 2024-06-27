package richerror

import (
	"errors"
)

type Kind int

const (
	KindInvalid Kind = iota + 1
	KindForbidden
	KindNotFound
	KindUnexpected
)

type RichError struct {
	operation string
	wrapError error
	message   string
	kind      Kind
	meta      map[string]interface{}
}

func New(operation string) RichError {
	return RichError{operation: operation}
}

func (r RichError) WithMessage(message string) RichError {
	r.message = message

	return r
}

func (r RichError) WithWrappedError(err error) RichError {
	r.wrapError = err

	return r
}

func (r RichError) WithKind(kind Kind) RichError {
	r.kind = kind

	return r
}

func (r RichError) WhitKind(kind Kind) RichError {
	r.kind = kind

	return r
}

func (r RichError) WhitMeta(meta map[string]interface{}) RichError {
	r.meta = meta

	return r
}

func (r RichError) Kind() Kind {
	if r.kind != 0 {
		return r.kind
	}

	var re RichError

	ok := errors.As(r.wrapError, &re)
	if !ok {
		return 0
	}

	return re.Kind()
}

func (r RichError) Message() string {
	if r.message != "" {
		return r.message
	}

	var re RichError

	ok := errors.As(r.wrapError, &re)
	if !ok {
		return r.wrapError.Error()
	}

	return re.Message()
}

func (r RichError) Error() string {
	return r.message
}
