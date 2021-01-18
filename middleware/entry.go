// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package middleware

// Entry for adapters grouped by category.
type Entry struct {
	category  string
	adapters  []Adapter
	next      *Entry
}

// Create entry
func NewEntry(category string) *Entry {
	e := &Entry{
		category: category,
	}
	return e
}

// Create next entry
func (e *Entry) NewNext(category string) *Entry {
	n := &Entry{
		category: category,
	}
	e.next = n
	return n
}

// Get the entry's category
func (e *Entry) Category() string {
	return e.category
}

// Get the next entry
func (e *Entry) Next() *Entry {
	return e.next
}

// Add adapter to the entry
func (e *Entry) addAdapter(adapter ...Adapter) {
	e.adapters = append(e.adapters, adapter...)
}

// Get the entry's adapters
func (e *Entry) Adapters() []Adapter {
	return e.adapters
}
