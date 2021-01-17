// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package ctx

import (
	"github.com/chinmobi/ginmod/security"
)

const CTX_SECURITY_HOLDER = "CTX_SECURITY_HOLDER"

type SecurityContext = security.SecurityContext

// Context holder used for per request, NOT goroutine safe!
type SecurityContextHolder interface {
	GetSecurityContex() SecurityContext
}

// Implement of the SecurityContextHolder
type ContextHolder struct {
	context   Context
	security  security.Context
}

func NewContextHolder() *ContextHolder {
	h := &ContextHolder{
	}
	return h
}

// Get the SecurityContext
func (h *ContextHolder) GetSecurityContex() SecurityContext {
	return &h.security
}

// Get the security raw context
func (h *ContextHolder) GetSecurity() *security.Context {
	return &h.security
}

// Got as the SecurityContextHolder (self)
func (h *ContextHolder) GetSecurityContextHolder() SecurityContextHolder {
	return h
}

// Get the Context
func (h *ContextHolder) GetContext() *Context {
	return &h.context
}
