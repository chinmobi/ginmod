// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package errors

// Common original error types base.
type OriginalErrorBase struct {
	Name    string  `json:"name"`
	Message string  `json:"message"`
}

func (err *OriginalErrorBase) Error() string {
	return err.Name + ": \"" + err.Message + "\""
}

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
