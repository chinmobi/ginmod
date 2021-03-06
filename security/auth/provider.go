// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

// AuthenticationProvider
type AuthenticationProvider interface {
	Authenticate(auth Authentication) (Authentication, error)
	Supports(auth Authentication) bool
}

type AuthProvider = AuthenticationProvider

// ProviderManager
type ProviderManager struct {
	providers []AuthProvider
}

// Get providers
func (pm *ProviderManager) Providers() []AuthProvider {
	return pm.providers
}

// Add provider
func (pm *ProviderManager) AddProvider(provider ...AuthProvider) {
	pm.providers = append(pm.providers, provider...)
}

// Clear providers
func (pm *ProviderManager) Clear() {
	if pm.providers != nil {
		pm.providers = pm.providers[0:0]
	}
}

// Authenticate the authentication
func (pm *ProviderManager) Authenticate(auth Authentication) (Authentication, error) {
	for i, cnt := 0, len(pm.providers); i < cnt; i++ {
		provider := pm.providers[i]

		if !provider.Supports(auth) {
			continue
		}

		result, err := provider.Authenticate(auth)
		if err != nil {
			return auth, err
		}

		if result == nil {
			result = auth
		}

		if result.IsAuthenticated() {
			return result, nil
		}

		if auth != result {
			auth = result
			i = -1 // New authentication, restart the loop.
		}
	}

	return auth, nil
}
