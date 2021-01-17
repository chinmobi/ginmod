// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package access

import (
	"github.com/chinmobi/ginmod/security/access"
)

// InterceptorBuilder
type InterceptorBuilder struct {
	permissions   *access.PermissionsGroup
	deniedHandler AccessDeniedHandler
	evaluator     access.PrivilegeEvaluator
}

// Create InterceptorBuilder with the OnAccessDeniedFunc
func NewBuilder(onAccessDenied OnAccessDeniedFunc) *InterceptorBuilder {
	builder := &InterceptorBuilder{
		deniedHandler: WrapAccessDeniedFunc(onAccessDenied),
		evaluator: access.SimplePrivilegeEvaluator{},
	}

	return builder
}

// Set the PermissionsGroup for build
func (b *InterceptorBuilder) SetPermissions(permissions *access.PermissionsGroup) {
	b.permissions = permissions
}

// Set the AccessDeniedHandler for build
func (b *InterceptorBuilder) SetDeniedHandler(handler AccessDeniedHandler) {
	b.deniedHandler = handler
}

// Set the PrivilegeEvaluator for build
func (b *InterceptorBuilder) SetEvaluator(evaluator access.PrivilegeEvaluator) {
	b.evaluator = evaluator
}

// Build SecurityInterceptor
func (b *InterceptorBuilder) Build() *SecurityInterceptor {
	si := &SecurityInterceptor{
		permissions:    b.permissions,
		deniedHandler:  b.deniedHandler,
		evaluator:      b.evaluator,
	}
	return si
}

// Build SecurityInterceptor for the PermissionsEntry
func (b *InterceptorBuilder) BuildFor(entry *PermissionsEntry) *SecurityInterceptor {
	si := &SecurityInterceptor{
		permissions:    &entry.permissions,
		deniedHandler:  b.deniedHandler,
		evaluator:      b.evaluator,
	}

	entry.interceptor = si

	return si
}
