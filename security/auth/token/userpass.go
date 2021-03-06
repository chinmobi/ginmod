// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package token

// UsernamePasswordAuthToken
type UsernamePasswordAuthToken struct {
	AbstractAuthToken
	username, password string
}

func NewUsernamePasswordAuthToken(username, password string) *UsernamePasswordAuthToken {
	token := &UsernamePasswordAuthToken{
//		AbstractAuthToken: {
//			authorities: []security.GrantedAuthority{},
//		},
		username: username,
		password: password,
	}
	return token
}

// --- Authentication methods ---

func (u *UsernamePasswordAuthToken) GetCredentials() interface{} {
	return u.password
}

func (u *UsernamePasswordAuthToken) GetPrincipal() interface{} {
	return u.username
}

// --- Additional methods ---

func (u *UsernamePasswordAuthToken) EraseCredentials() {
	u.password = ""
}

func (u *UsernamePasswordAuthToken) GetPassword() string {
	return u.password
}

func (u *UsernamePasswordAuthToken) GetUsername() string {
	return u.username
}
