// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package restful

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RespDataEntity(c *gin.Context, apiVersion string, data interface{}) {
	c.JSON(http.StatusOK, CreateApiDataBody(apiVersion, data))
}

func RespCreatedDataEntity(c *gin.Context, apiVersion string, data interface{}) {
	c.JSON(http.StatusCreated, CreateApiDataBody(apiVersion, data))
}

func RespNoContent(c *gin.Context) {
	//c.Writer.WriteHeader(http.StatusNoContent)
	c.Status(http.StatusNoContent)
}
