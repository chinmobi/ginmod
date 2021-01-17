// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/chinmobi/ginmod/security"
	"github.com/chinmobi/ginmod/security/errors"
)

type Authentication = security.Authentication

type ErrAuthentication = errors.ErrAuthentication

// AuthenticationManager
type AuthenticationManager interface {
	Authenticate(auth Authentication) (Authentication, error)
}

type AuthManager = AuthenticationManager

// Wrap the errors.NewAuthenticationErr

func NewAuthenticationErr(cause error) *ErrAuthentication {
	return errors.NewAuthenticationErr(cause)
}
