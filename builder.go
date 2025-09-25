package pd

import (
	"errors"
	"fmt"
)

type Builder Error

func newBuilder() Builder {
	return Builder{
		message:    messageUndefined,
		status:     statusUndefined,
		code:       codeUndefined,
		stackTrace: newStackTrace(),
	}
}

func (b Builder) New(message string) error {
	b.message = message

	return Error(b)
}

func (b Builder) Errorf(format string, args ...any) error {
	b.message = fmt.Sprintf(format, args...)

	return Error(b)
}

func (b Builder) Wrap(err error) error {
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

func (b Builder) Wrapf(err error, format string, args ...any) error {
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

func (b Builder) Status(status int) Builder {
	b.status = status

	return b
}

func (b Builder) Code(code string) Builder {
	b.code = code

	return b
}
