// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package errors

import (
	"errors"
)

// ErrUsernameNotFound
func NewUsernameNotFoundErr() error {
	return errors.New("UsernameNotFound")
}

// ErrBadCredentials (ErrIncorrectPassword)
func NewBadCredentialsErr() error {
	return errors.New("BadCredentials")
}

// ErrAccountStatus:
// (ErrAccountExpired, ErrAccountLocked, ErrCredentialExpired, ErrAccountDisabled)

// ErrAccountExpired
func NewAccountExpiredErr() error {
	return errors.New("AccountExpired")
}

// ErrAccountLocked
func NewAccountLockedErr() error {
	return errors.New("AccountLocked")
}

// ErrCredentialExpired
func NewCredentialExpiredErr() error {
	return errors.New("CredentialExpired")
}

// ErrAccountDisabled
func NewAccountDisabledErr() error {
	return errors.New("AccountDisabled")
}
