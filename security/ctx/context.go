// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package ctx

// Context used for per request, NOT goroutine safe!
type Context struct {
	keys  map[string]interface{}
}

// Get value by key
func (c *Context) Get(key string) interface{} {
	if c.keys == nil {
		return nil
	}
	return c.keys[key]
}

// Set context key/value, return old value by the key
func (c *Context) Set(key string, value interface{}) interface{} {
	if c.keys == nil {
		c.keys = make(map[string]interface{})
		c.keys[key] = value
		return nil
	}

	old := c.keys[key]
	c.keys[key] = value

	return old
}
