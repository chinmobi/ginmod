// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package restful

import (
	"net/http"

	"github.com/chinmobi/ginmod/errors"

	"github.com/gin-gonic/gin"
)

// Response bad request error
func RespBadRequest(c *gin.Context, apiVersion string, err error) {
	err = errors.NewBadRequestErrOf(err)

	apiErr := NewApiErrorEntity(http.StatusBadRequest, err)

	c.JSON(apiErr.GetStatusCode(), CreateApiErrorBody(apiVersion, apiErr))
}

// Response error caused by service
func RespServiceError(c *gin.Context, apiVersion string, err error) {
	var statusCode uint = http.StatusInternalServerError

	switch err.(type) {
	case *errors.ErrNotFound:
		statusCode = http.StatusNotFound

	case *errors.ErrAlreadyExists, *errors.ErrBadRequest:
		statusCode = http.StatusBadRequest

	case *errors.ErrLackOfParameter, *errors.ErrInvalidParameter:
		statusCode = http.StatusBadRequest

	case *errors.ErrMethodNotAllowed:
		statusCode = http.StatusMethodNotAllowed

	case *errors.ErrNotImplemented:
		statusCode = http.StatusNotImplemented

	case *errors.ErrInternalError:
		// Nothing to do.

	default:
		err = errors.NewInternalCausedErrOf(err)
	}

	apiErr := NewApiErrorEntity(statusCode, err)

	c.JSON(apiErr.GetStatusCode(), CreateApiErrorBody(apiVersion, apiErr))
}
