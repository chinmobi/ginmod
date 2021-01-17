// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package middleware

import (
	"github.com/gin-gonic/gin"
)

// Adapter for gin.HandlerFunc to manage middleware.
type Adapter interface {
	HandlerFunc() gin.HandlerFunc
	TearDown() error
}
