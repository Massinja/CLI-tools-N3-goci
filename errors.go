package main

import (
	"errors"
	"fmt"
)

var (
	ErrValidation = errors.New("validation failed")

	// to handle errors outside of CI step.
	ErrSignal = errors.New("received interrupt signal")
)

type stepError struct {
	step  string
	msg   string
	cause error
}

func (s *stepError) Error() string {
	if s.cause == nil {
		return "STEP: " + s.step + ": " + s.msg
	}

	return fmt.Sprintf("STEP: %s: %s%v", s.step, s.msg, s.cause)
}

func (s *stepError) Is(target error) bool {
	t, ok := target.(*stepError)
	if !ok {
		return false
	}

	return t.step == s.step
}

func (s *stepError) Unwrap() error {
	return s.cause
}
