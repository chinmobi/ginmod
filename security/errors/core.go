// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package errors

// Common caused error types base.
type CausedErrorBase struct {
	Name    string  `json:"name"`
	Message string  `json:"message"`
	Cause   error   `json:"-"`
}

func (err *CausedErrorBase) Error() string {
	str := err.Name + ": \"" + err.Message + "\""
	if err.Cause != nil {
		str += " Caused by: {" + err.Cause.Error() + "}"
	}
	return str
}


// ErrAuthentication
type ErrAuthentication struct {
	CausedErrorBase
}

func NewAuthenticationErr(cause error) *ErrAuthentication {
	err := &ErrAuthentication{
		CausedErrorBase: CausedErrorBase{
			Name: "ErrAuthentication",
			Message: cause.Error(),
			Cause: cause,
		},
	}
	return err
}


// ErrAccessDenied
type ErrAccessDenied struct {
	CausedErrorBase
}

func NewAccessDeniedErr(cause error) *ErrAccessDenied {
	err := &ErrAccessDenied{
		CausedErrorBase: CausedErrorBase{
			Name: "ErrAccessDenied",
			Message: cause.Error(),
			Cause: cause,
		},
	}
	return err
}
