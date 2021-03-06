// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package token

// AnonymousAuthToken
type AnonymousAuthToken struct {
	AbstractAuthToken
	principal interface{}
}

func NewAnonymousAuthToken(principal interface{}) *AnonymousAuthToken {
	token := &AnonymousAuthToken{
		principal: principal,
	}
	return token
}

// --- Authentication methods ---

// Always returns an empty string
func (a *AnonymousAuthToken) GetCredentials() interface{} {
	return ""
}

// Get the principal for anonymous
func (a *AnonymousAuthToken) GetPrincipal() interface{} {
	return a.principal
}
