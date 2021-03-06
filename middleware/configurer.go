// Copyright 2020 Zhaoping Yu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package middleware

// Configurer to configure middleware entries.
type Configurer struct {
	entries *Entry
}

// Builder for helping to configure entries.
type Builder struct {
	configurer *Configurer
	entry *Entry
}

// Create Configurer
func NewConfigurer(entries *Entry) *Configurer {
	c := &Configurer{
		entries: entries,
	}
	return c
}

func (c *Configurer) getOrNewEntry(category string) *Entry {
	entry := c.entries
	for entry != nil {
		if entry.Category() == category {
			return entry
		}

		if entry.next == nil {
			entry = entry.NewNext(category)
			return entry
		} else {
			entry = entry.next
		}
	}
	entry = NewEntry(category)
	c.entries = entry
	return entry
}

// Create builder of category to configure entries
func (c *Configurer) Build(category string) *Builder {
	entry := c.getOrNewEntry(category)

	b := &Builder{
		configurer: c,
		entry: entry,
	}
	return b
}

// Add adapter to the builder
func (b *Builder) AddMwAdapter(adapter ...Adapter) *Builder {
	b.entry.addAdapter(adapter...)
	return b
}

// To another builder
func (b *Builder) Build(category string) *Builder {
	return b.configurer.Build(category)
}
