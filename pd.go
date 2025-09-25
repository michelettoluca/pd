package pd

import (
	"errors"
)

func New(message string) error {
	return newBuilder().New(message)
}

func Errorf(format string, args ...any) error {
	return newBuilder().Errorf(format, args...)
}

func Wrap(err error) error {
	return newBuilder().Wrap(err)
}

func Wrapf(err error, format string, args ...any) error {
	return newBuilder().Wrapf(err, format, args...)
}

func Status(status int) Builder {
	return newBuilder().Status(status)
}

func Code(code string) Builder {
	return newBuilder().Code(code)
}

func findDeepest(err Error) Error {
	if err.err == nil {
		return err
	}

	var pdErr Error
	if ok := errors.As(err.err, &pdErr); ok {
		return findDeepest(pdErr)
	}

	return err
}

func findClosestMatch(err Error, predicate func(err Error) bool) (*Error, bool) {
	if predicate(err) {
		return &err, true
	}

	if err.err == nil {
		return nil, false
	}

	var pdErr Error
	if ok := errors.As(err.err, &pdErr); ok {
		return findClosestMatch(pdErr, predicate)
	}

	return nil, false
}
