// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package restful

import (
	"fmt"
)

// ApiErrorEntity
type ApiErrorEntity struct {
	StatusCode   uint     `json:"statusCode"`
	Code         int      `json:"code"`
	Message      string   `json:"message"`
	Errors       []error  `json:"errors"`
}

// Create ApiErrorEntity with status code and error
func NewApiErrorEntity(statusCode uint, err error) *ApiErrorEntity {
	code := int(statusCode)
	return NewApiErrorEntityS(code, err.Error(), err).SetStatusCode(statusCode)
}

// Create ApiErrorEntity with code and error
func NewApiErrorEntityA(code int, err error) *ApiErrorEntity {
	return NewApiErrorEntityS(code, err.Error(), err)
}

// Create ApiErrorEntity with code, message, and errors
func NewApiErrorEntityS(code int, msg string, errs ...error) *ApiErrorEntity {
	apiErr := &ApiErrorEntity{
		StatusCode:  500,
		Code:        code,
		Message:     msg,
	}

	apiErr.Errors = make([]error, 0, 1)

	apiErr.AppendError(errs...)

	return apiErr
}

// Error interface method
func (apiErr *ApiErrorEntity) Error() string {
	return fmt.Sprintf("{statusCode: %d, code: %d, message: %s}", apiErr.StatusCode, apiErr.Code, apiErr.Message)
}

// Append error
func (apiErr *ApiErrorEntity) AppendError(errs ...error) *ApiErrorEntity {
	apiErr.Errors = append(apiErr.Errors, errs...)
	return apiErr
}

// Set status code
func (apiErr *ApiErrorEntity) SetStatusCode(statusCode uint) *ApiErrorEntity {
	apiErr.StatusCode = statusCode
	return apiErr
}

// Get status code
func (apiErr *ApiErrorEntity) GetStatusCode() int {
	return int(apiErr.StatusCode)
}
