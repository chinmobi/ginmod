// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/chinmobi/ginmod/security/auth"
)

// AuthenticationProcessorBase
type AuthenticationProcessorBase struct {
	manager      auth.AuthManager
	authHandler  AuthHandler
}

type AuthProcessorBase = AuthenticationProcessorBase

// Get the AuthManager
func (ap *AuthProcessorBase) AuthManager() auth.AuthManager {
	return ap.manager
}

// Get the AuthHandler
func (ap *AuthProcessorBase) AuthHandler() AuthHandler {
	return ap.authHandler
}

// ProcessorConfigurer
type ProcessorConfigurer struct {
	manager        auth.ProviderManager
	handlerSet     AuthHandlerSet
	processorBase  *AuthProcessorBase
}

func NewProcessorConfigurer() *ProcessorConfigurer {
	configurer := &ProcessorConfigurer{
	}
	return configurer
}

// Get the AuthProcessorBase
func (pc *ProcessorConfigurer) ProcessorBase() *AuthProcessorBase {
	if pc.processorBase != nil {
		return pc.processorBase
	}

	base := &AuthProcessorBase{
		manager: &pc.manager,
		authHandler: &pc.handlerSet,
	}

	pc.processorBase = base

	return base
}

// Add AuthProvider to ProviderManager
func (pc *ProcessorConfigurer) AddProvider(provider ...auth.AuthProvider) {
	pc.manager.AddProvider(provider...)
}

// Add OnAuthSuccessFunc to AuthHandlerSet
func (pc *ProcessorConfigurer) AddAuthSuccessFunc(onSuccess ...OnAuthSuccessFunc) {
	pc.handlerSet.AddAuthSuccessFunc(onSuccess...)
}

// Add OnAuthFailureFunc to AuthHandlerSet
func (pc *ProcessorConfigurer) AddAuthFailureFunc(onFailure ...OnAuthFailureFunc) {
	pc.handlerSet.AddAuthFailureFunc(onFailure...)
}

// Get the ProviderManager
func (pc *ProcessorConfigurer) ProviderManager() *auth.ProviderManager {
	return &pc.manager
}

// Get the AuthHandlerSet
func (pc *ProcessorConfigurer) AuthHandlerSet() *AuthHandlerSet {
	return &pc.handlerSet
}

// Clear configured ProviderManager and AuthHandlerSet
func (pc *ProcessorConfigurer) Reset() {
	pc.manager.Clear()
	pc.handlerSet.Clear()
}
