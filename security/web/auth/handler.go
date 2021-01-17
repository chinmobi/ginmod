// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/chinmobi/ginmod/security"
	"github.com/chinmobi/ginmod/security/errors"

	"github.com/gin-gonic/gin"
)

type Authentication = security.Authentication

type ErrAuthentication = errors.ErrAuthentication

type OnAuthSuccessFunc func(c *gin.Context, auth Authentication) (bool, error)
type OnAuthFailureFunc func(c *gin.Context, authErr *ErrAuthentication) (bool, error)

// AuthSuccessHandler
type AuthSuccessHandler interface {
	OnAuthSuccess(c *gin.Context, auth Authentication) (bool, error)
}

// AuthFailureHandler
type AuthFailureHandler interface {
	OnAuthFailure(c *gin.Context, authErr *ErrAuthentication) (bool, error)
}

// AuthHandler
type AuthHandler interface {
	AuthSuccessHandler
	AuthFailureHandler
}

// AuthHandlerSetter
type AuthHandlerSetter interface {
	AddAuthSuccessFunc(onSuccess ...OnAuthSuccessFunc)
	AddAuthFailureFunc(onFailure ...OnAuthFailureFunc)
}

// Implement of AuthHandlerSetter
type AuthHandlerSet struct {
	successFuncChain  []OnAuthSuccessFunc
	failureFuncChain  []OnAuthFailureFunc
}

func NewAuthHandlerSet() *AuthHandlerSet {
	set := &AuthHandlerSet{
	}
	return set
}

// --- AuthHandlerSetter methods ---

// Add OnAuthSuccessFunc
func (set *AuthHandlerSet) AddAuthSuccessFunc(onSuccess ...OnAuthSuccessFunc) {
	set.successFuncChain = append(set.successFuncChain, onSuccess...)
}

// Add OnAuthFailureFunc
func (set *AuthHandlerSet) AddAuthFailureFunc(onFailure ...OnAuthFailureFunc) {
	set.failureFuncChain = append(set.failureFuncChain, onFailure...)
}

// Clear the handler set
func (set *AuthHandlerSet) Clear() {
	if set.successFuncChain != nil {
		set.successFuncChain = set.successFuncChain[0:0]
	}
	if set.failureFuncChain != nil {
		set.failureFuncChain = set.failureFuncChain[0:0]
	}
}

// --- AuthHandler methods ---

// Handle OnAuthSuccess
func (set *AuthHandlerSet) OnAuthSuccess(c *gin.Context, auth Authentication) (bool, error) {
	for i := len(set.successFuncChain)-1; i >= 0; i-- {
		onSuccess := set.successFuncChain[i]

		done, err := onSuccess(c, auth)
		if err != nil {
			return done, err
		}
		if done {
			return true, nil
		}
	}
	return false, nil
}

// Handle OnAuthFailure
func (set *AuthHandlerSet) OnAuthFailure(c *gin.Context, authErr *ErrAuthentication) (bool, error) {
	for i := len(set.failureFuncChain)-1; i >= 0; i-- {
		onFailure := set.failureFuncChain[i]

		done, err := onFailure(c, authErr)
		if err != nil {
			return done, err
		}
		if done {
			return true, nil
		}
	}
	return false, nil
}

// Wrap the errors.NewAuthenticationErr
func NewAuthenticationErr(cause error) *ErrAuthentication {
	return errors.NewAuthenticationErr(cause)
}
