package pd

import (
	"errors"
	"fmt"
)

type builder Error

type builderError interface {
	New(message string) error
	Errorf(format string, args ...any) error
	Wrap(err error) error
	Wrapf(err error, format string, args ...any) error
}

type builderCode interface {
	Code(code string) builderError
}

func newBuilder() builder {
	return builder{
		message:    messageUndefined,
		status:     statusUndefined,
		code:       codeUndefined,
		stackTrace: newStackTrace(),
	}
}

func (b builder) New(message string) error {
	b.message = message

	return Error(b)
}

func (b builder) Errorf(format string, args ...any) error {
	b.message = fmt.Sprintf(format, args...)

	return Error(b)
}

func (b builder) Wrap(err error) error {
	if err == nil {
		return nil
	}
	b.err = err

	var e Error
	if errors.As(err, &e) {
		b.status = e.status
		b.code = e.code
	}

	return Error(b)
}

func (b builder) Wrapf(err error, format string, args ...any) error {
	if err == nil {
		return nil
	}

	b.message = fmt.Sprintf(format, args...)
	b.err = err

	var e Error
	if errors.As(err, &e) {
		b.status = e.status
		b.code = e.code
	}

	return Error(b)
}

func (b builder) Status(status int) builderCode {
	b.status = status

	return b
}

func (b builder) Code(code string) builderError {
	b.code = code

	return b
}
