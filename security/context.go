// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package security

// SecurityContext to manage Authentication for per request.
type SecurityContext interface {
	GetAuthentication() Authentication
	SetAuthentication(auth Authentication) Authentication
	CleanAuthentication() Authentication
}

// Implement of SecurityContext
type Context struct {
	auth Authentication
}

func NewContext() *Context {
	ctx := &Context{
	}
	return ctx;
}

// --- SecurityContext methods ---

func (ctx *Context) GetAuthentication() Authentication {
	return ctx.auth
}

func (ctx *Context) SetAuthentication(auth Authentication) Authentication {
	old := ctx.auth
	ctx.auth = auth
	return old
}

func (ctx *Context) CleanAuthentication() Authentication {
	old := ctx.auth
	ctx.auth = nil
	return old
}
