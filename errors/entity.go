// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package errors


// NewErrNotFound
type ErrNotFound struct {
	OriginalErrorBase
}

func NewNotFoundErrBy(entity, cond, value string) *ErrNotFound {
	msg := entity + " by " + cond + ": |" + value + "|"
	return NewNotFoundErr(msg)
}

func NewNotFoundErr(msg string) *ErrNotFound {
	err := &ErrNotFound{
		OriginalErrorBase: OriginalErrorBase{
			Name: "ErrNotFound",
			Message: msg,
		},
	}
	return err
}


// ErrAlreadyExists
type ErrAlreadyExists struct {
	OriginalErrorBase
}

func NewAlreadyExistsErrFor(entity, field, value string) *ErrAlreadyExists {
	msg := entity + " for " + field + ": |" + value + "|"
	return NewAlreadyExistsErr(msg)
}

func NewAlreadyExistsErr(msg string) *ErrAlreadyExists {
	err := &ErrAlreadyExists{
		OriginalErrorBase: OriginalErrorBase{
			Name: "ErrAlreadyExists",
			Message: msg,
		},
	}
	return err
}
