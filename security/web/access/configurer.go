// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package access

import (
	"github.com/chinmobi/ginmod/security/access"

	"github.com/gin-gonic/gin"
)

type SimpleRolePermission = access.SimpleRolePermission

// PermissionsEntry
type PermissionsEntry struct {
	permissions  access.PermissionsGroup
	handlerFunc  gin.HandlerFunc
	interceptor  *SecurityInterceptor
}

// Reset the PermissionsEntry
func (pe *PermissionsEntry) Reset() {
	pe.permissions.Clear()
	pe.handlerFunc = nil
	pe.interceptor = nil
}

// Get the PermissionsGroup
func (pe *PermissionsEntry) PermissionsGroup() *access.PermissionsGroup {
	return &pe.permissions
}

// Add RolePermission that required for access
func (pe *PermissionsEntry) AddPermission(permission ...access.RolePermission) {
	pe.permissions.AddPermission(permission...)
}

func (pe *PermissionsEntry) configure(c *gin.Context) {
	c.Set(CTX_ACCESS_PERMISSIONS, &pe.permissions)

	c.Next()
}

// Get the PermissionsEntry's gin.HandlerFunc for setting up required permissions
func (pe *PermissionsEntry) ConfigureHandlerFunc() gin.HandlerFunc {
	if pe.handlerFunc == nil {
		pe.handlerFunc = pe.createHandlerFunc()
	}
	return pe.handlerFunc
}

func (pe *PermissionsEntry) createHandlerFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		pe.configure(c)
	}
}

// Get the AccessDecisionAgent
func (pe *PermissionsEntry) AccessDecisionAgent() *AccessDecisionAgent {
	if pe.interceptor == nil {
		pe.interceptor = buildDefaultInterceptor(pe)
	}

	return pe.interceptor.DecisionAgent()
}

// PermissionsConfigurer, configure required permissions for access
type PermissionsConfigurer struct {
	entries []*PermissionsEntry
}

// Create PermissionsConfigurer
func NewPermissionsConfigurer() *PermissionsConfigurer {
	pc := &PermissionsConfigurer{
	}
	return pc
}

// Reset the PermissionsConfigurer
func (pc *PermissionsConfigurer) Reset() {
	if pc.entries != nil {
		for i, cnt := 0, len(pc.entries); i < cnt; i++ {
			pc.entries[i].Reset()
		}

		pc.entries = pc.entries[0:0]
	}
}

// Create ConfigureEntry with group to setup required permissions
func (pc *PermissionsConfigurer) ConfigureEntry(group string) *PermissionsEntry {
	for i, cnt := 0, len(pc.entries); i < cnt; i++ {
		entry := pc.entries[i]
		if entry.permissions.Name() == group {
			return entry
		}
	}

	entry := &PermissionsEntry{
	}
	entry.permissions.Init(group)

	pc.entries = append(pc.entries, entry)

	return entry
}
