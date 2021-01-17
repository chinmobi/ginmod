// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/chinmobi/ginmod/security/auth"
	"github.com/chinmobi/ginmod/security/ctx"

	"github.com/gin-gonic/gin"
)

// AuthenticationHelper
type AuthenticationHelper interface {
	AttemptAuthentication(c *gin.Context) (Authentication, error)
	TearDown() error
}

type AuthHelper = AuthenticationHelper

// AuthenticationProcessor
type AuthenticationProcessor struct {
	base        *AuthProcessorBase
	helper       AuthHelper
	handlerFunc  gin.HandlerFunc
}

type AuthProcessor = AuthenticationProcessor

func NewAuthProcessor(base *AuthProcessorBase, helper AuthHelper) *AuthProcessor {
	processor := &AuthProcessor{
		base: base,
		helper: helper,
	}
	return processor
}

// Get the AuthManager
func (ap *AuthProcessor) AuthManager() auth.AuthManager {
	return ap.base.manager
}

// Get the AuthHandler
func (ap *AuthProcessor) AuthHandler() AuthHandler {
	return ap.base.authHandler
}

func (ap *AuthProcessor) process(c *gin.Context) {
	if ap.processAuth(c) {
		c.Next()
	}
}

func (ap *AuthProcessor) processAuth(c *gin.Context) bool {
	securityContext := ctx.SetSecurityHolder(c).GetSecurityContex()

	err := ap.doProcess(c, securityContext)
	if err != nil {
		securityContext.CleanAuthentication()

		authErr, ok := err.(*ErrAuthentication)
		if !ok {
			authErr = NewAuthenticationErr(err)
		}

		done, _ := ap.AuthHandler().OnAuthFailure(c, authErr)
		return !done
	}

	return true
}

func (ap *AuthProcessor) doProcess(c *gin.Context, s ctx.SecurityContext) error {
	auth := s.GetAuthentication()
	if auth != nil && auth.IsAuthenticated() {
		return nil
	}

	auth, err := ap.helper.AttemptAuthentication(c)
	if err != nil {
		return err
	}
	if auth == nil {
		return nil
	}

	s.SetAuthentication(auth)

	if auth.IsAuthenticated() {
		return nil
	}

	result, err := ap.AuthManager().Authenticate(auth)
	if err != nil {
		return err
	}

	if result != nil {
		s.SetAuthentication(result)
	}

	return nil
}

// Get the AuthProcessor's gin.HandlerFunc
func (ap *AuthProcessor) HandlerFunc() gin.HandlerFunc {
	if ap.handlerFunc == nil {
		ap.handlerFunc = ap.createHandlerFunc()
	}
	return ap.handlerFunc
}

func (ap *AuthProcessor) createHandlerFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		ap.process(c)
	}
}

// TearDown the AuthProcessor
func (ap *AuthProcessor) TearDown() error {
	return ap.helper.TearDown()
}
