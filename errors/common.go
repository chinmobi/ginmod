// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package errors

// ErrLackOfParameter
type ErrLackOfParameter struct {
	OriginalErrorBase
}

func newLackOfParameterErr(params []string) *ErrLackOfParameter {
	var msg string
	for i, cnt := 0, len(params); i < cnt; i++ {
		if i > 0 {
			msg += ", "
		}
		msg += params[i]
	}

	err := &ErrLackOfParameter{
		OriginalErrorBase: OriginalErrorBase{
			Name: "ErrLackOfParameter",
			Message: msg,
		},
	}
	return err
}

func NewLackOfParameterErr(params ...string) *ErrLackOfParameter {
	return newLackOfParameterErr(params)
}


// ErrInvalidParameter
type ErrInvalidParameter struct {
	OriginalErrorBase
	ParamName   string  `json:"param_name"`
	ParamValue  string  `json:"param_value"`
}

func NewInvalidParameterErrOf(name, value string) *ErrInvalidParameter {
	msg := name + ": [" + value + "]"
	return NewInvalidParameterErr(name, value, msg)
}

func NewInvalidParameterErr(name, value, msg string) *ErrInvalidParameter {
	err := &ErrInvalidParameter{
		OriginalErrorBase: OriginalErrorBase{
			Name: "ErrInvalidParameter",
			Message: msg,
		},
		ParamName: name,
		ParamValue: value,
	}
	return err
}


// ErrInternalError
type ErrInternalError struct{
	CausedErrorBase
}

func NewInternalCausedErrOf(cause error) *ErrInternalError {
	return NewInternalCausedErr(cause.Error(), cause)
}

func NewInternalCausedErr(msg string, cause error) *ErrInternalError {
	err := &ErrInternalError{
		CausedErrorBase: CausedErrorBase{
			Name: "ErrInternalError",
			Message: msg,
			Cause: cause,
		},
	}
	return err
}


// ErrBadRequest
type ErrBadRequest struct {
	CausedErrorBase
}

func NewBadRequestErrOf(cause error) *ErrBadRequest {
	return NewBadRequestErr(cause.Error(), cause)
}

func NewBadRequestErr(msg string, cause error) *ErrBadRequest {
	err := &ErrBadRequest{
		CausedErrorBase: CausedErrorBase{
			Name: "ErrBadRequest",
			Message: msg,
			Cause: cause,
		},
	}
	return err
}


// ErrMethodNotAllowed
type ErrMethodNotAllowed struct {
	CausedErrorBase
}

func NewMethodNotAllowedErrOf(cause error) *ErrMethodNotAllowed {
	return NewMethodNotAllowedErr(cause.Error(), cause)
}

func NewMethodNotAllowedErr(msg string, cause error) *ErrMethodNotAllowed {
	err := &ErrMethodNotAllowed{
		CausedErrorBase: CausedErrorBase{
			Name: "ErrMethodNotAllowed",
			Message: msg,
			Cause: cause,
		},
	}
	return err
}


// ErrNotImplemented
type ErrNotImplemented struct {
	OriginalErrorBase
}

func NewNotImplementedErr(msg string) *ErrNotImplemented {
	err := &ErrNotImplemented{
		OriginalErrorBase: OriginalErrorBase{
			Name: "ErrNotImplemented",
			Message: msg,
		},
	}
	return err
}
