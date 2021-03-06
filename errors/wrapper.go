// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package errors

import (
	"strings"
)

// ErrWrapErrors for wrap multiple errors
type ErrWrapErrors struct {
	errors []error
}

func NewWrapErrorsErr() *ErrWrapErrors {
	err := &ErrWrapErrors{
	}
	return err
}

// Wrap error
func (w *ErrWrapErrors) Wrap(err ...error) *ErrWrapErrors {
	w.errors = append(w.errors, err...)
	return w
}

// Get the wrapped errors
func (w *ErrWrapErrors) WrappedErrors() []error {
	return w.errors
}

// Error interface method
func (w *ErrWrapErrors) Error() string {
	var str strings.Builder

	str.WriteString("errors: [")

	for i, cnt := 0, len(w.errors); i < cnt; i++ {
		if i > 0 {
			str.WriteString(", ")
		}
		str.WriteString("{")
		str.WriteString(w.errors[i].Error())
		str.WriteString("}")
	}

	str.WriteString("]")

	return str.String()
}

// To error or nil (if no wrapped error)
func (w *ErrWrapErrors) ToError() error {
	if w.errors != nil {
		return w
	}
	return nil
}
