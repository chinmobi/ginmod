// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package restful

// ApiDataBody
type ApiDataBody struct {
	ApiVersion  string       `json:"apiVersion"`
	Data        interface{}  `json:"data"`
}

// Create ApiDataBody
func CreateApiDataBody(apiVersion string, data interface{}) *ApiDataBody {
	dataBody := &ApiDataBody{
		ApiVersion: apiVersion,
		Data: data,
	}
	return dataBody
}

// ApiErrorBody
type ApiErrorBody struct {
	ApiVersion  string           `json:"apiVersion"`
	Error       *ApiErrorEntity  `json:"error"`
}

// Create ApiErrorBody
func CreateApiErrorBody(apiVersion string, err *ApiErrorEntity) *ApiErrorBody {
	errBody := &ApiErrorBody{
		ApiVersion: apiVersion,
		Error: err,
	}
	return errBody
}
